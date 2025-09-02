package athletic

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

func NewAthleticRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.Athletic, error) {
	var athletic []models.Athletic

	result := r.DB.Find(&athletic)
	if result.Error != nil {
		return nil, fmt.Errorf("erro ao buscar atletica no banco de dados: %w", result.Error)
	}
	return athletic, nil
}

func (r *Repository) FindById(athleticId int) (models.Athletic, error) {
	var athletic models.Athletic

	if err := r.DB.Preload("University").First(&athletic, athleticId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Athletic{}, fmt.Errorf("atletica com ID %d não encontrado", athleticId)
		}
		return models.Athletic{}, fmt.Errorf("erro ao buscar atletica: %w", err)
	}
	return athletic, nil
}

func (r *Repository) Create(athletic *models.Athletic) error {
	if err := r.DB.Create(athletic).Error; err != nil {
		return fmt.Errorf("erro ao criar atletica: %w", err)
	}

	return r.DB.Preload("University").First(athletic, athletic.ID).Error
}

func (r *Repository) Update(id int, athletic models.Athletic) error {
	if err := r.DB.First(&models.Athletic{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("atlética com ID %d não encontrada", id)
		}
		return fmt.Errorf("erro ao verificar atlética: %w", err)
	}

	if err := r.DB.Model(&models.Athletic{}).Where("id = ?", id).Updates(athletic).Error; err != nil {
		return fmt.Errorf("erro ao atualizar os dados da atlética: %w", err)
	}

	return nil
}

func (r *Repository) Delete(id int) error {
	result := r.DB.Delete(&models.Athletic{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao deletar atlética: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("atlética com ID %d não encontrada", id)
	}

	return nil
}
