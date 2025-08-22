package championship

import "backend-go/internal/models"

type Service struct {
	repo Repository
}

func NewChampionshipService(repo *Repository) *Service {
	return &Service{repo: *repo}
}

func (service Service) FindChampionship() ([]models.Championship, error) {

	championship, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return championship, nil
}
