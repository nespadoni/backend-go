package auth

import (
	"backend-go/internal/models"
	"backend-go/internal/modules/user"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *user.Repository
	validate  *validator.Validate
	jwtSecret string
}

func NewAuthService(userRepo *user.Repository, validate *validator.Validate, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		validate:  validate,
		jwtSecret: jwtSecret,
	}
}

// Register cria um novo usuário com foto e retorna token
func (s *AuthService) Register(req RegisterRequest, file multipart.File, header *multipart.FileHeader) (map[string]interface{}, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("dados inválidos: %w", err)
	}

	// Verificar se o e-mail já está em uso
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("e-mail já está em uso")
	}

	// Criptografar a senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao criptografar senha: %w", err)
	}

	var newUser models.User
	err = copier.Copy(&newUser, &req)
	if err != nil {
		return nil, err
	}
	newUser.Password = string(hashedPassword)

	if req.UniversityID != nil && *req.UniversityID != "none" && *req.UniversityID != "" {
		if universityID, err := strconv.ParseUint(*req.UniversityID, 10, 32); err == nil {
			universityIDUint := uint(universityID)
			newUser.UniversityID = &universityIDUint
		}
	}

	// Criar usuário no banco para obter o ID
	if err := s.userRepo.Create(&newUser); err != nil {
		return nil, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	// Salvar a foto de perfil se ela foi enviada
	if file != nil && header != nil {
		// Validações
		if !isValidImageType(header) {
			log.Printf("Tentativa de upload de tipo de arquivo inválido para o usuário ID: %d", newUser.Id)
		} else if header.Size > 5*1024*1024 { // 5MB
			log.Printf("Arquivo de foto de perfil muito grande para o usuário ID: %d", newUser.Id)
		} else {
			// Salvar a foto e obter a URL
			photoURL, err := s.saveProfilePhoto(file, header, newUser.Id)
			if err != nil {
				// Loga o erro mas não impede o registro de ser concluído
				log.Printf("Não foi possível salvar a foto de perfil para o usuário ID %d: %v", newUser.Id, err)
			} else {
				// Atualiza o usuário com a URL da foto
				newUser.ProfilePhotoURL = &photoURL
				if err := s.userRepo.Update(newUser.Id, &newUser); err != nil {
					log.Printf("Falha ao atualizar o usuário ID %d com a URL da foto: %v", newUser.Id, err)
				}
			}
		}
	}

	// Gerar token JWT
	token, err := s.generateJWT(newUser)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token: %w", err)
	}

	// Converter para response
	var userResponse user.Response
	err = copier.Copy(&userResponse, &newUser)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
		"user":  userResponse,
	}, nil
}

// Login autentica um usuário e retorna um JWT junto com os dados do usuário
func (s *AuthService) Login(req LoginRequest) (map[string]interface{}, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("dados inválidos: %w", err)
	}

	foundUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("credenciais inválidas") // Erro genérico por segurança
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("credenciais inválidas") // Erro genérico
	}

	// Gerar o token JWT
	token, err := s.generateJWT(foundUser)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token: %w", err)
	}

	// Preparar a resposta do usuário para não expor dados sensíveis
	var userResponse user.Response
	if err := copier.Copy(&userResponse, &foundUser); err != nil {
		return nil, fmt.Errorf("erro ao mapear dados do usuário: %w", err)
	}

	// Criar a resposta final
	result := map[string]interface{}{
		"token": token,
		"user":  userResponse,
	}

	return result, nil
}

// isValidImageType verifica se o tipo de conteúdo do arquivo é uma imagem válida.
func isValidImageType(header *multipart.FileHeader) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/jpg":  true,
	}
	contentType := header.Header.Get("Content-Type")
	return validTypes[contentType]
}

// saveProfilePhoto salva o arquivo da foto de perfil e retorna sua URL de acesso.
func (s *AuthService) saveProfilePhoto(file multipart.File, header *multipart.FileHeader, userID uint) (string, error) {
	uploadDir := "uploads/profile_photos"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	// Usar UnixNano para garantir um nome de arquivo único
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("user_%d_%d%s", userID, time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			log.Printf("Erro ao fechar arquivo: %v", err)
		}
	}(dst)

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	// Retorna a URL relativa para ser usada no frontend
	return "/" + filePath, nil
}

func (s *AuthService) generateJWT(user models.User) (string, error) {
	// Define as "claims" (informações) do token
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expira em 72 horas
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com a chave secreta
	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
