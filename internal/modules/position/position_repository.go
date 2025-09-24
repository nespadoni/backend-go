package position

import (
	"backend-go/internal/models"
	"backend-go/internal/repository"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	repository.BaseRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.Position, error) {
	var positions []models.Position
	if err := r.DB.Find(&positions).Error; err != nil {
		return nil, fmt.Errorf("failed to find all positions: %w", err)
	}
	return positions, nil
}

func (r *Repository) FindByID(id uint) (*models.Position, error) {
	var position models.Position
	if err := r.DB.First(&position, id).Error; err != nil {
		return nil, fmt.Errorf("failed to find position with id %v: %w", id, err)
	}
	return &position, nil
}

func (r *Repository) Create(position *models.Position) error {
	if err := r.DB.Create(position).Error; err != nil {
		return fmt.Errorf("failed to create position %v: %w", position, err)
	}
	return nil
}

func (r *Repository) Update(id uint, position *models.Position) (*models.Position, error) {
	position.Id = id

	// Tenta Atualizar
	result := r.DB.Updates(position)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to update position %v: %w", position, result.Error)
	}

	// Verificar se alguma linha foi afetada. Se não, o registro não foi encontrado
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("position with id %v not found", id)
	}

	return position, nil
}

func (r *Repository) Delete(id uint) error {
	result := r.DB.Delete(&models.Position{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete position %v: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("position with id %v not found", id)
	}

	return nil
}
