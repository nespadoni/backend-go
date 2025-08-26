package athletic

import (
	"backend-go/internal/models"
	"backend-go/internal/repository"
	"errors"
	"fmt"
	"strconv"

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

func (r *Repository) Update(id string, athletic models.Athletic) error {
	athleticID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	if err := r.DB.First(&models.Athletic{}, athleticID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("atletica com ID %s não encontrado", id)
		}
		return fmt.Errorf("erro ao verificar atletica: %w", err)
	}

	if err := r.DB.Model(&models.Athletic{}).Where("id = ?").Updates(athletic).Error; err != nil {
		return fmt.Errorf("erro ao atualizar os dados da atlética: %w", err)
	}

	return r.DB.Preload("University").First(athletic, athleticID).Error
}

func (r *Repository) Delete(id string) error {
	athleticId, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	result := r.DB.Delete(&models.Athletic{}, athleticId)
	if result.Error != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("usuario com ID %d não encontrado", athleticId)
	}

	return nil
}
