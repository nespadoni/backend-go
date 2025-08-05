package championship

type ChampionshipService struct {
	repo ChampionshipRepository
}

func NewChampionshipService(repo *ChampionshipRepository) *ChampionshipService {
	return &ChampionshipService{repo: *repo}
}

func (service ChampionshipService) FindChampionship() ([]Championship, error) {

	championship, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return championship, nil
}
