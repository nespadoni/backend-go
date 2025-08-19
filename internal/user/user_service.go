package user

import (
	"backend-go/internal/models"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type UserService struct {
	repo     *UserRepository
	validate *validator.Validate
}

func NewUserService(repo *UserRepository, validate *validator.Validate) *UserService {
	return &UserService{
		repo:     repo,
		validate: validate,
	}
}

func (s *UserService) GetAll() ([]UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro no serviço ao buscar usuários: %w", err)
	}

	var userResponses []UserResponse
	if err := copier.Copy(&userResponses, users); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return userResponses, nil
}

func (s *UserService) GetById(userID int) (UserResponse, error) {
	user, err := s.repo.GetById(userID)
	if err != nil {
		return UserResponse{}, fmt.Errorf("erro no serviço ao buscar usuário: %w", err)
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		return UserResponse{}, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return userResponse, nil
}

func (s *UserService) CreateUser(req CreateUserRequest) (UserResponse, error) {
	// Validar request
	if err := s.validate.Struct(req); err != nil {
		return UserResponse{}, fmt.Errorf("dados inválidos: %w", err)
	}

	// Verificar se email já existe
	_, err := s.repo.GetByEmail(req.Email)
	if err == nil {
		return UserResponse{}, fmt.Errorf("email %s já está em uso", req.Email)
	}

	// Criar novo usuário
	var newUser models.User
	if err := copier.Copy(&newUser, &req); err != nil {
		return UserResponse{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Create(&newUser); err != nil {
		return UserResponse{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, &newUser); err != nil {
		return UserResponse{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return userResponse, nil
}

func (s *UserService) UpdateUser(id string, req UpdateUserRequest) (UserResponse, error) {
	// Validar request
	if err := s.validate.Struct(&req); err != nil {
		return UserResponse{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var user models.User
	if err := copier.Copy(&user, &req); err != nil {
		return UserResponse{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Update(id, &user); err != nil {
		return UserResponse{}, fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, &user); err != nil {
		return UserResponse{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return userResponse, nil
}

func (s *UserService) DeleteUser(userID int) error {
	return s.repo.Delete(userID)
}
