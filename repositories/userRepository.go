package repositories

import (
	"backend-go/models"
)

type UserRepository interface {
	FindAll() []models.User
	FindById(id int) models.User
	FindByName(name string) models.User
	FindByEmail(email string) models.User
	SaveUser(user models.User) models.User
	UpdateUser(userId int, user models.User)
	DeleteUser(userId int)
}
