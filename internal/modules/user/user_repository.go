package user

import (
	"backend-go/internal/models"
	"backend-go/internal/repository" // Adicionar import
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	repository.BaseRepository // Usar BaseRepository como os outros
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.User, error) {
	var users []models.User

	if err := r.DB.Preload("University").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
	}

	return users, nil
}

func (r *Repository) GetById(id uint) (models.User, error) {
	var user models.User

	err := r.DB.Preload("University").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("usuário com Id %d não encontrado", id)
		}
		return models.User{}, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

func (r *Repository) GetByEmail(email string) (models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("usuário com email %s não encontrado", email)
		}
		return models.User{}, fmt.Errorf("erro ao buscar usuário por email: %w", err)
	}

	return user, nil
}

func (r *Repository) Create(user *models.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return r.DB.Preload("University").First(user, user.Id).Error
}

func (r *Repository) Update(id uint, user *models.User) error {
	if err := r.DB.First(&models.User{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("usuário com Id %d não encontrado", id)
		}
		return fmt.Errorf("erro ao verificar usuário: %w", err)
	}

	if err := r.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return r.DB.Preload("University").First(user, id).Error
}

func (r *Repository) Delete(id uint) error {
	result := r.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("usuário com Id %d não encontrado", id)
	}

	return nil
}
