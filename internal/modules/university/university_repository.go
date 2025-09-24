package university

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

func NewUniversityRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.University, error) {
	var universities []models.University

	if err := r.DB.Find(&universities).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar universidades: %w", err)
	}

	return universities, nil
}

func (r *Repository) FindById(id uint) (models.University, error) {
	var university models.University

	if err := r.DB.First(&university, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.University{}, fmt.Errorf("universidade com Id %d não encontrada", id)
		}
		return models.University{}, fmt.Errorf("erro ao buscar universidade: %w", err)
	}
	return university, nil
}

func (r *Repository) Create(university *models.University) error {
	if err := r.DB.Create(university).Error; err != nil {
		return fmt.Errorf("erro ao registrar universidade no banco de dados: %w", err)
	}

	return r.DB.First(university, university.Id).Error
}

func (r *Repository) Update(id uint, university *models.University) (*models.University, error) {
	// Verificar se existe
	if err := r.DB.First(&models.University{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("universidade com Id %d não encontrada", id)
		}
		return nil, fmt.Errorf("erro ao verificar universidade: %w", err)
	}

	// Atualizar
	if err := r.DB.Model(&university).Where("id = ?", id).Updates(university).Error; err != nil {
		return nil, fmt.Errorf("erro ao atualizar universidade: %w", err)
	}

	// Buscar registro atualizado
	var updatedUniversity models.University
	if err := r.DB.First(&updatedUniversity, id).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar universidade atualizada: %w", err)
	}

	return &updatedUniversity, nil
}

func (r *Repository) Delete(id uint) error {
	result := r.DB.Delete(&models.University{}, id)

	if result.Error != nil {
		return fmt.Errorf("erro ao deletar universidade: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("universidade com Id %d não encontrada", id)
	}

	return nil
}
