package championship

type ChampionshipController struct {
	service *ChampionshipService
}

func NewChampionshipController(service *ChampionshipService) *ChampionshipController {
	return &ChampionshipController{service: service}
}

func (controller *ChampionshipController) GetChampionship() ([]Championship, error) {

	championship, err := controller.service.FindChampionship()
	if err != nil {

		return championship, err
	}

	return championship, err
}
