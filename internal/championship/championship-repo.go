package championship

import (
	"backend-go/internal/models"

	"gorm.io/gorm"
)

type ChampionshipRepository struct {
	DB *gorm.DB
}

func NewChampionshipRepository(db *gorm.DB) *ChampionshipRepository {
	return &ChampionshipRepository{DB: db}
}

func (repo *ChampionshipRepository) FindAll() ([]models.Championship, error) {
	var championship []models.Championship

	result := repo.DB.Find(&championship)
	if result.Error != nil {
		return nil, result.Error
	}

	return championship, nil
}
