package sport

import (
	"backend-go/internal/models"
	"backend-go/internal/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	repository.BaseRepository
}

func NewSportRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.Sport, error) {
	var sports []models.Sport

	if err := r.DB.Find(&sports).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar esportes: %w", err)
	}

	return sports, nil
}

func (r *Repository) FindById(id uint) (models.Sport, error) {
	var sport models.Sport

	if err := r.DB.Preload("Positions").First(&sport, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Sport{}, fmt.Errorf("esporte com ID %d não encontrado", id)
		}
		return models.Sport{}, fmt.Errorf("erro ao buscar esporte: %w", err)
	}

	return sport, nil
}

func (r *Repository) FindByName(name string) (models.Sport, error) {
	var sport models.Sport

	if err := r.DB.Where("name = ?", name).First(&sport).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Sport{}, fmt.Errorf("esporte com nome %s não encontrado", name)
		}
		return models.Sport{}, fmt.Errorf("erro ao buscar esporte por nome: %w", err)
	}

	return sport, nil
}

func (r *Repository) Create(sport *models.Sport) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		// Verificar se já existe um esporte com o mesmo nome
		var existingSport models.Sport
		if err := tx.Where("name = ?", sport.Name).First(&existingSport).Error; err == nil {
			return fmt.Errorf("já existe um esporte com o nome '%s'", sport.Name)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("erro ao verificar esporte existente: %w", err)
		}

		if err := tx.Create(sport).Error; err != nil {
			return fmt.Errorf("erro ao criar esporte: %w", err)
		}

		return tx.Preload("Positions").First(sport, sport.ID).Error
	})
}

func (r *Repository) Update(id uint, sport *models.Sport) (*models.Sport, error) {
	// Verificar se existe
	if err := r.DB.First(&models.Sport{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("esporte com ID %d não encontrado", id)
		}
		return nil, fmt.Errorf("erro ao verificar esporte: %w", err)
	}

	// Atualizar
	if err := r.DB.Model(&models.Sport{}).Where("id = ?", id).Updates(sport).Error; err != nil {
		return nil, fmt.Errorf("erro ao atualizar esporte: %w", err)
	}

	// Buscar registro atualizado
	var updatedSport models.Sport
	if err := r.DB.Preload("Positions").First(&updatedSport, id).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar esporte atualizado: %w", err)
	}

	return &updatedSport, nil
}

func (r *Repository) Delete(id uint) error {
	result := r.DB.Delete(&models.Sport{}, id)

	if result.Error != nil {
		return fmt.Errorf("erro ao deletar esporte: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("esporte com ID %d não encontrado", id)
	}

	return nil
}

func (r *Repository) FindPopular() ([]models.Sport, error) {
	var sports []models.Sport

	if err := r.DB.Where("is_popular = ? AND is_active = ?", true, true).Find(&sports).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar esportes populares: %w", err)
	}

	return sports, nil
}
