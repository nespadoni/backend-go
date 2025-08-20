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

func (repo *ChampionshipRepository) FindAll() ([]models.Championship, error) {
	var championships []models.Championship

	result := repo.DB.Find(&championships)
	if result.Error != nil {
		// Retorna o erro quando houver problema na consulta
		return nil, result.Error
	}

	return championships, nil
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
