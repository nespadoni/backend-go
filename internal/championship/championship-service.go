package championship

import "backend-go/internal/models"

type ChampionshipService struct {
	repo ChampionshipRepository
}

func NewChampionshipService(repo *ChampionshipRepository) *ChampionshipService {
	return &ChampionshipService{repo: *repo}
}

func (service ChampionshipService) FindChampionship() ([]models.Championship, error) {

	championship, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return championship, nil
}
