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

func (r *Repository) FindById(championship *models.Championship) (*models.Championship, error) {
	var result models.Championship

	if err := r.DB.First(&result, championship.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("campeonato não encontrado")
		}
		return nil, fmt.Errorf("erro no banco de dados: %w", err)
	}

	return &result, nil
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

func (r *Repository) Update(championship *models.Championship) error {
	result := r.DB.Save(championship)

	if result.Error != nil {
		return fmt.Errorf("erro ao deletar campeonato: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("campeonato não encontrado")
	}

	return nil
}

func (r *Repository) Delete(id uint) error {
	result := r.DB.Delete(&models.Championship{}, id)

	if result.Error != nil {
		return fmt.Errorf("erro ao deletar campeonato: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("campeonato não encontrado")
	}

	return nil
}
