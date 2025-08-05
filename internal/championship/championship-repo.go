package championship

import "gorm.io/gorm"

type ChampionshipRepository struct {
	DB *gorm.DB
}

func NewChampionshipRepository(db *gorm.DB) *ChampionshipRepository {
	return &ChampionshipRepository{DB: db}
}

func (repo *ChampionshipRepository) FindAll() ([]Championship, error) {
	var championship []Championship

	result := repo.DB.Find(&championship)
	if result.Error != nil {
		return nil, result.Error
	}

	return championship, nil
}
