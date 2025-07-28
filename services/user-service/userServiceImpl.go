package services

import (
	"backend-go/repositories"
	"backend-go/response"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repositories.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (us *UserServiceImpl) GetAll() []response.UserResponse {
	users := us.UserRepository.FindAll()

	var userResponse []response.UserResponse
	copier.Copy(&userResponse, users)

	return userResponse

}
