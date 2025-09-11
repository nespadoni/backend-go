package factories

import (
	"backend-go/internal/modules/athletic"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AthleticFactory struct{}

func NewAthleticFactory() *AthleticFactory {
	return &AthleticFactory{}
}

func (f *AthleticFactory) CreateController(db *gorm.DB, validate *validator.Validate) interface{} {
	repo := f.CreateRepository(db)
	service := f.CreateService(repo, validate)
	return athletic.NewAthleticController(service.(*athletic.Service))
}

func (f *AthleticFactory) CreateRepository(db *gorm.DB) interface{} {
	return athletic.NewAthleticRepository(db)
}

func (f *AthleticFactory) CreateService(repo interface{}, validate *validator.Validate) interface{} {
	return athletic.NewAthleticService(repo.(*athletic.Repository), validate)
}
