package services

import (
	"backend-go/request"
	"backend-go/response"
)

type UserService interface {
	GetAll() []response.UserResponse
	GetById(userId int) response.UserResponse
	CreateUser(user request.CreateUserRequest) response.UserResponse
	UpdateUser(userId int, user request.UpdateUserRequest) response.UserResponse
	DeleteUser(userId int)
}
