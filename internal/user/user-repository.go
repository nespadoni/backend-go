package user

import (
	"backend-go/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var usuarios []models.User
	if resultado := r.DB.Find(&usuarios); resultado.Error != nil {
		return usuarios, resultado.Error
	}

	return usuarios, nil
}

func (r *UserRepository) GetById(id int) (models.User, error) {
	var usuario models.User
	resultado := r.DB.Where("ID = ?", id).Find(&usuario)

	if resultado.Error != nil {
		return usuario, resultado.Error
	}

	return usuario, nil
}

func (r *UserRepository) BuscarPorNome(nome string) (models.User, error) {
	var usuario models.User
	resultado := r.DB.Where("Nome = ?", nome).Find(&usuario)
	if resultado.Error != nil {
		return usuario, resultado.Error
	}
	return usuario, nil
}

func (r *UserRepository) BuscarPorEmail(email string) (models.User, error) {
	var usuario models.User
	resultado := r.DB.Where("Email = ?", email).Find(&usuario)

	if resultado.Error != nil {
		return usuario, resultado.Error
	}

	return usuario, nil
}

func (r *UserRepository) SaveUser(usuario *models.User) error {
	var userRole models.Role
	//ID 2 É o padrão de usuário normal
	r.DB.Where("ID = ?", 2).Find(&userRole)
	usuario.Role = userRole
	usuario.RoleID = userRole.ID

	resultado := r.DB.Create(&usuario)

	if resultado.Error != nil {
		return resultado.Error
	}

	err := r.DB.Preload("University").First(&usuario, usuario.ID)
	if err != nil {
		return err.Error
	}

	return nil
}

func (r *UserRepository) AtualizarUsuario(id string, usuario *models.User) error {
	fmt.Println("ID STRING :", id)
	resultado := r.DB.Where("id = ?", id).Updates(&usuario)
	if resultado.Error != nil {
		return resultado.Error
	}

	r.DB.Where("ID = ?", id).Preload("University").Preload("Role").First(&usuario)
	fmt.Println("DEPOIS DO FIRST", usuario)

	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	usuario := models.User{}
	resultado := r.DB.Where("id = ?", id).Delete(&usuario)

	if resultado.Error != nil {
		return resultado.Error
	}

	return nil
}
