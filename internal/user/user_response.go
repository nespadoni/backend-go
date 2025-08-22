package user

import "backend-go/internal/models"

type Response struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Email      string            `json:"email"`
	Phone      string            `json:"phone"`
	Role       models.Role       `json:"role"`
	University models.University `json:"university"`
}
