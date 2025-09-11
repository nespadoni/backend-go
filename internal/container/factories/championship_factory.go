package factories

import (
	"backend-go/internal/modules/championship"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ChampionshipFactory struct{}

func NewChampionshipFactory() *ChampionshipFactory {
	return &ChampionshipFactory{}
}

func (f *ChampionshipFactory) CreateController(db *gorm.DB, validate *validator.Validate) interface{} {
	repo := f.CreateRepository(db)
	service := f.CreateService(repo, validate)
	return championship.NewChampionshipController(service.(*championship.Service))
}

func (f *ChampionshipFactory) CreateRepository(db *gorm.DB) interface{} {
	return championship.NewChampionshipRepository(db)
}

func (f *ChampionshipFactory) CreateService(repo interface{}, validate *validator.Validate) interface{} {
	return championship.NewChampionshipService(repo.(*championship.Repository), validate)
}
