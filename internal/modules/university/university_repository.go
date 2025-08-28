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
	var university []models.University

	if err := r.DB.Find(&university).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar universidades: %w", err)
	}

	return university, nil
}

func (r *Repository) FindById(id int) (models.University, error) {
	var university models.University

	if err := r.DB.First(&university, id).Error; err != nil {
		return models.University{}, fmt.Errorf("universidade com ID %d não encontrado", id)
	}
	return university, nil
}

func (r *Repository) Create(university *models.University) error {
	if err := r.DB.Create(university).Error; err != nil {
		return fmt.Errorf("erro ao registrar universidade no banco de dados: %w", err)
	}

	return r.DB.First(university, university.ID).Error
}

func (r *Repository) Update(id int, university *models.University) error {
	var existingUniversity models.University
	if err := r.DB.First(&existingUniversity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("universidade com ID %d não encontrado", id)
		}
		return fmt.Errorf("erro ao verificar universidade: %w", err)
	}

	if err := r.DB.Model(&models.University{}).Where("id = ?", id).Updates(university).Error; err != nil {
		return fmt.Errorf("erro ao atualizar universidade: %w", err)
	}

	return nil
}
