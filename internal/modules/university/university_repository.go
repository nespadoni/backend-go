package university

import (
	"backend-go/internal/models"
	"backend-go/internal/repository"
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
