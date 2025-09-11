package factories

import (
	"backend-go/internal/modules/sport"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SportFactory struct{}

func NewSportFactory() *SportFactory {
	return &SportFactory{}
}

func (f *SportFactory) CreateController(db *gorm.DB, validate *validator.Validate) interface{} {
	repo := f.CreateRepository(db)
	service := f.CreateService(repo, validate)
	return sport.NewSportController(service.(*sport.Service))
}

func (f *SportFactory) CreateRepository(db *gorm.DB) interface{} {
	return sport.NewSportRepository(db)
}

func (f *SportFactory) CreateService(repo interface{}, validate *validator.Validate) interface{} {
	return sport.NewSportService(repo.(*sport.Repository), validate)
}
