package user

import (
	"backend-go/internal/models"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type Service struct {
	repo     *Repository
	validate *validator.Validate
}

func NewUserService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		validate: validate,
	}
}

func (s *Service) GetAll() ([]Response, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro no serviço ao buscar usuários: %w", err)
	}

	var userResponses []Response
	if err := copier.Copy(&userResponses, users); err != nil {
		return nil, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return userResponses, nil
}

func (s *Service) GetById(userID int) (Response, error) {
	user, err := s.repo.GetById(userID)
	if err != nil {
		return Response{}, fmt.Errorf("erro no serviço ao buscar usuário: %w", err)
	}

	var userResponse Response
	if err := copier.Copy(&userResponse, user); err != nil {
		return Response{}, fmt.Errorf("erro ao converter dados: %w", err)
	}

	return userResponse, nil
}

func (s *Service) CreateUser(req CreateUserRequest) (Response, error) {
	// Validar request
	if err := s.validate.Struct(req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	// Verificar se email já existe
	_, err := s.repo.GetByEmail(req.Email)
	if err == nil {
		return Response{}, fmt.Errorf("email %s já está em uso", req.Email)
	}

	// Criar novo usuário
	var newUser models.User
	if err := copier.Copy(&newUser, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Create(&newUser); err != nil {
		return Response{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	var userResponse Response
	if err := copier.Copy(&userResponse, &newUser); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return userResponse, nil
}

func (s *Service) UpdateUser(id string, req UpdateUserRequest) (Response, error) {
	// Validar request
	if err := s.validate.Struct(&req); err != nil {
		return Response{}, fmt.Errorf("dados inválidos: %w", err)
	}

	var user models.User
	if err := copier.Copy(&user, &req); err != nil {
		return Response{}, fmt.Errorf("erro ao processar dados: %w", err)
	}

	if err := s.repo.Update(id, &user); err != nil {
		return Response{}, fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	var userResponse Response
	if err := copier.Copy(&userResponse, &user); err != nil {
		return Response{}, fmt.Errorf("erro ao converter resposta: %w", err)
	}

	return userResponse, nil
}

func (s *Service) DeleteUser(userID int) error {
	return s.repo.Delete(userID)
}
