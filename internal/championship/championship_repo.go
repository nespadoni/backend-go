package championship

import (
	"backend-go/internal/models"
	"backend-go/internal/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ChampionshipRepository struct {
	repository.BaseRepository
}

func NewChampionshipRepository(db *gorm.DB) *ChampionshipRepository {
	return &ChampionshipRepository{BaseRepository: repository.BaseRepository{DB: db}}
}

func (r *ChampionshipRepository) FindAll() ([]models.Championship, error) {
	var championships []models.Championship

	result := r.DB.Find(&championships)
	if result.Error != nil {
		// Retorna o erro quando houver problema na consulta
		return nil, result.Error
	}

	return championships, nil
}

func (r *ChampionshipRepository) FindById(championship *models.Championship) (*models.Championship, error) {
	var result models.Championship

	if err := r.DB.First(&result, championship.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("campeonato não encontrado")
		}
		return nil, fmt.Errorf("erro no banco de dados: %w", err)
	}

	return &result, nil
}

func (r *ChampionshipRepository) Create(championship *models.Championship) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		var athletic models.Athletic
		if err := tx.First(&athletic, championship.AthleticID).Error; err != nil {
			return fmt.Errorf("atlética não encontrada: %w", err)
		}

		return tx.Create(championship).Error
	})
}

func (r *ChampionshipRepository) Update(championship *models.Championship) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		var existing models.Championship

		if err := tx.First(&existing, championship.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("registro não encontado")
			}

			return fmt.Errorf("erro de banco: %w", err)
		}

		return tx.Save(championship).Error
	})
}

func (r *ChampionshipRepository) Delete(championship *models.Championship) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		var existing models.Championship

		if err := tx.First(&existing, championship.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("registro não encontrado")
			}

			return fmt.Errorf("erro de banco: %w", err)
		}

		return tx.Delete(championship).Error
	})
}
