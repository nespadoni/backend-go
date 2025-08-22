package user

import (
	"backend-go/internal/models"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]models.User, error) {
	var users []models.User

	if err := r.db.Preload("University").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
	}

	return users, nil
}

func (r *Repository) GetById(id int) (models.User, error) {
	var user models.User

	err := r.db.Preload("University").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("usuário com ID %d não encontrado", id)
		}
		return models.User{}, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

func (r *Repository) GetByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("usuário com email %s não encontrado", email)
		}
		return models.User{}, fmt.Errorf("erro ao buscar usuário por email: %w", err)
	}

	return user, nil
}

func (r *Repository) Create(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return r.db.Preload("University").First(user, user.ID).Error
}

func (r *Repository) Update(id string, user *models.User) error {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	// Verificar se existe
	if err := r.db.First(&models.User{}, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("usuário com ID %s não encontrado", id)
		}
		return fmt.Errorf("erro ao verificar usuário: %w", err)
	}

	// Atualizar
	if err := r.db.Model(&models.User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	// Recarregar com relacionamentos
	return r.db.Preload("University").First(user, userID).Error
}

func (r *Repository) Delete(id int) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("usuário com ID %d não encontrado", id)
	}

	return nil
}
