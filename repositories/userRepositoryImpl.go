package repositories

import (
	"backend-go/handler"
	"backend-go/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (us *UserRepositoryImpl) FindAll() []models.User {
	var users []models.User
	result := us.DB.Find(&users)
	handler.NewReadError("Error finding all users: ", result.Error)
	return users
}

func (us *UserRepositoryImpl) FindById(id int) models.User {
	var user models.User
	result := us.DB.Where("ID = ?", id).Find(&user)
	handler.NewReadError("Error finding by Id", result.Error)
	return user
}

func (us *UserRepositoryImpl) FindByName(name string) models.User {
	var user models.User
	result := us.DB.Where("Name = ?", name).Find(&user)
	handler.NewReadError("Error finding by Name", result.Error)
	return user
}

func (us *UserRepositoryImpl) FindByEmail(email string) models.User {
	var user models.User
	result := us.DB.Where("Email = ?", email).Find(&user)
	handler.NewReadError("Error finding by Email", result.Error)
	return user
}

func (us *UserRepositoryImpl) SaveUser(user models.User) models.User {
	result := us.DB.Create(&user)
	handler.NewCreateError("Error saving user: ", result.Error)
	return user
}

func (us *UserRepositoryImpl) UpdateUser(userId int, newUser models.User) {
	result := us.DB.Where("id = ?", userId).Updates(newUser)
	handler.NewUpdateError("Error updating user: ", result.Error)
}

func (us *UserRepositoryImpl) DeleteUser(userId int) {
	user := models.User{}
	result := us.DB.Where("id = ?", userId).Delete(user)
	handler.NewDeleteError("Error deleting user: ", result.Error)
}
