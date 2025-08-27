package university

import (
	"backend-go/internal/models"
	"backend-go/internal/repository"

	"gorm.io/gorm"
)

type Repository struct {
	repository.BaseRepository
}

func NewUniversityRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.University, error) {
	var champions []models.University

}
