package user

import (
	"backend-go/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/nespadoni/goerror"
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
		return nil, goerror.NovoErroInterno("COPIER_ERROR", erro.Error())
	}

	return userResponses, err
}

func (s *UserService) GetById(userId int) (UserResponse, error) {

	user, err := s.repo.GetById(userId)
	if err != nil {
		return UserResponse{}, err
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		return UserResponse{}, goerror.NovoErroBancoDados("COPIER_ERROR", err.Error())
	}

	return userResponse, nil
}

func (s *UserService) CreateUser(req CreateUserRequest) (UserResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return UserResponse{}, goerror.NovoErroValidacao("VALIDATE_ERROR", err.Error())
	}

	var newUser models.User
	if err := copier.Copy(&newUser, &req); err != nil {
		return UserResponse{}, goerror.NovoErroInterno("COPIER_ERROR", err.Error())
	}

	createdUser, err := s.repo.SaveUser(newUser)
	if err != nil {
		return UserResponse{}, err
	}

	var userResponse UserResponse
	if err := copier.Copy(&userResponse, createdUser); err != nil {
		return UserResponse{}, goerror.NovoErroInterno("COPIER_ERROR", err.Error())
	}

	return userResponse, nil
}

// func (s *UserService) DeleteUser(userId int) error {
// 	if userId == 0 {
// 		return goerror.NovoErroValidacao("3", "teste")
// 	}

// 	if err := s.repo.DeleteUser(userId); err != nil {

// 	}

// 	return nil
// }
