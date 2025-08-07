package user

import "backend-go/internal/models"

type CreateUserRequest struct {
	Name       string            `validate:"required" json:"name"`
	Email      string            `validate:"required" json:"email"`
	Password   string            `validate:"required" json:"password"`
	Phone      string            `validate:"required" json:"phone"`
	University models.University `validate:"required" json:"university"`
}

type UpdateUserRequest struct {
	Name       string            `validate:"required" json:"name"`
	Email      string            `validate:"required" json:"email"`
	Password   string            `validate:"required" json:"password"`
	Phone      string            `validate:"required" json:"phone"`
	University models.University `validate:"required" json:"university"`
}
