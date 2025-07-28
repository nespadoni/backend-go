package controllers

import (
	"backend-go/services/user-service"
)

type UserController struct {
	userService services.UserService
}
