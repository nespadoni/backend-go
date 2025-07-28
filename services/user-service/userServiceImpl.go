package services

import (
	"backend-go/handler"
	"backend-go/models"
	"backend-go/repositories"
	"backend-go/request"
	"backend-go/response"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type UserServiceImpl struct {
	repo     repositories.UserRepository
	validate *validator.Validate
}

// NewUserServiceImpl cria uma nova instância do serviço
func NewUserServiceImpl(repo repositories.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		repo:     repo,
		validate: validate,
	}
}

// GetAll retorna todos os usuários
func (us *UserServiceImpl) GetAll() ([]response.UserResponse, error) {
	users := us.repo.FindAll()

	var userResponses []response.UserResponse
	if err := copier.Copy(&userResponses, users); err != nil {
		return nil, handler.Converter("User -> UserResponse", err)
	}

	return userResponses, nil
}

// GetById retorna um usuário pelo ID
func (us *UserServiceImpl) GetById(userId int) (response.UserResponse, error) {
	user := us.repo.FindById(userId)

	// Se usuário não existir, tratar como NotFound
	if userId == 0 {
		return response.UserResponse{}, handler.NotFound("Usuário", nil)
	}

	var userResponse response.UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		return response.UserResponse{}, handler.Converter("User -> UserResponse", err)
	}

	return userResponse, nil
}

// CreateUser cria um novo usuário
func (us *UserServiceImpl) CreateUser(req request.CreateUserRequest) (response.UserResponse, error) {
	// Validação dos dados de entrada
	if err := us.validate.Struct(req); err != nil {
		return response.UserResponse{}, handler.Validation(err)
	}

	// Conversão Request -> Model
	var newUser models.User
	if err := copier.Copy(&newUser, &req); err != nil {
		return response.UserResponse{}, handler.Converter("CreateUserRequest -> User", err)
	}

	// Persiste no repositório
	createdUser := us.repo.SaveUser(newUser)

	// Conversão Model -> Response
	var userResponse response.UserResponse
	if err := copier.Copy(&userResponse, createdUser); err != nil {
		return response.UserResponse{}, handler.Converter("User -> UserResponse", err)
	}

	return userResponse, nil
}
