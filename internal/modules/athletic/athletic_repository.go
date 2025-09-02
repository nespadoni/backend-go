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
	var athletics []models.Athletic

	if err := r.DB.Preload("University").Preload("Creator").Find(&athletics).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar atléticas no banco de dados: %w", err)
	}
	return athletics, nil
}

func (r *Repository) FindById(athleticId uint) (models.Athletic, error) {
	var athletic models.Athletic

	if err := r.DB.Preload("University").Preload("Creator").First(&athletic, athleticId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Athletic{}, fmt.Errorf("atletica com ID %d não encontrado", athleticId)
		}
		return models.Athletic{}, fmt.Errorf("erro ao buscar atletica: %w", err)
	}
	return athletic, nil
}

func (r *Repository) Create(athletic *models.Athletic) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		// Valida se a universidade existe
		var university models.University
		if err := tx.First(&university, athletic.UniversityID).Error; err != nil {
			return fmt.Errorf("universidade não encontrada: %w", err)
		}

		// Valida se o criador existe
		var creator models.User
		if err := tx.First(&creator, athletic.CreatorID).Error; err != nil {
			return fmt.Errorf("usuário criador não encontrado: %w", err)
		}

		if err := tx.Create(athletic).Error; err != nil {
			return fmt.Errorf("erro ao criar atlética: %w", err)
		}

		// Carregar a atlética criada com relacionamentos
		return tx.Preload("University").Preload("Creator").First(athletic, athletic.ID).Error
	})
}

func (r *Repository) Update(id uint, athletic *models.Athletic) (*models.Athletic, error) {
	// Verificar se a atlética existe no DB
	if err := r.DB.First(&models.Athletic{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("atlética com ID %d não encontrada", id)
		}
		return nil, fmt.Errorf("erro ao verificar atlética: %w", err)
	}

	// Atualizar os dados da atlética
	if err := r.DB.Model(&models.Athletic{}).Where("id = ?", id).Updates(athletic).Error; err != nil {
		return nil, fmt.Errorf("erro ao atualizar os dados da atlética: %w", err)
	}

	// Buscar o registro atualizado com relacionamentos
	var updatedAthletic models.Athletic
	if err := r.DB.Preload("University").Preload("Creator").First(&updatedAthletic, id).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar atlética atualizada: %w", err)
	}

	return &updatedAthletic, nil
}

func (r *Repository) Delete(id uint) error {
	result := r.DB.Delete(&models.Athletic{}, id)
	if result.Error != nil {
		return fmt.Errorf("erro ao deletar atlética: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("atlética com ID %d não encontrada", id)
	}

	return nil
}
