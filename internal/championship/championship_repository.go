package championship

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

func NewChampionshipRepository(db *gorm.DB) *Repository {
	return &Repository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *Repository) FindAll() ([]models.Championship, error) {
	var championships []models.Championship

	result := r.DB.Find(&championships)
	if result.Error != nil {
		// Retorna o erro quando houver problema na consulta
		return nil, result.Error
	}

	return championships, nil
}

func (r *Repository) FindById(id int) (models.Championship, error) {
	var championship models.Championship

	if err := r.DB.Preload("Athletic").First(&championship, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Championship{}, fmt.Errorf("campeonato com ID %d não encontrado", id)
		}
		return models.Championship{}, fmt.Errorf("erro ao buscar campeonato: %w", err)
	}

	return championship, nil

}

func (r *Repository) Create(championship *models.Championship) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		var athletic models.Athletic
		if err := tx.First(&athletic, championship.AthleticID).Error; err != nil {
			return fmt.Errorf("atlética não encontrada: %w", err)
		}

		return tx.Create(championship).Error
	})
}

func (r *Repository) Update(id int, championship *models.Championship) error {
	if err := r.DB.First(&models.Championship{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("campeonato com ID %d não encontrado", id)
		}
		return fmt.Errorf("erro ao verificar campeonato: %w", err)
	}

	if err := r.DB.Model(&models.Championship{}).Where("id = ?", id).Updates(championship).
		Error; err != nil {
		return fmt.Errorf("erro ao atualizar campeonato: %w", err)
	}

	return r.DB.Preload("Athletic").First(championship, id).Error

}

func (r *Repository) Delete(id int) error {
	result := r.DB.Delete(&models.Championship{}, id)

	if result.Error != nil {
		return fmt.Errorf("erro ao deletar campeonato: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("campeonato não encontrado")
	}

	return nil
}
