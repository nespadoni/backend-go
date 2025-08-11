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

	var userResponses []UserResponse
	if erro := copier.Copy(&userResponses, users); erro != nil {
		return nil, err
	}

	return userResponses, nil
}

func (s *UserService) GetById(userId int) (UserResponse, error) {

	user, err := s.repo.GetById(userId)
	if err != nil {
		return UserResponse{}, err
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		return UserResponse{}, err
	}

	return userResponse, nil
}

func (s *UserService) CreateUser(req CreateUserRequest) (UserResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return UserResponse{}, err
	}

	var newUser models.User
	if err := copier.Copy(&newUser, &req); err != nil {
		return UserResponse{}, err
	}

	err := s.repo.SaveUser(&newUser)
	if err != nil {
		return UserResponse{}, err
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, &newUser); err != nil {
		return UserResponse{}, err
	}

	return userResponse, nil
}

func (s *UserService) UpdateUser(id string, req UpdateUserRequest) (UserResponse, error) {
	if err := s.validate.Struct(&req); err != nil {
		return UserResponse{}, err
	}

	var user models.User
	if err := copier.Copy(&user, &req); err != nil {
		return UserResponse{}, err
	}
	fmt.Println("User: ", user)

	if err := s.repo.AtualizarUsuario(id, &user); err != nil {
		return UserResponse{}, err
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, &user); err != nil {
		return UserResponse{}, err
	}

	return userResponse, nil
}

func (s *UserService) DeleteUser(userId int) error {

	if err := s.repo.DeleteUser(userId); err != nil {
		return err
	}

	return nil
}
