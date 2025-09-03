package factories

import (
	"backend-go/internal/modules/university"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UniversityFactory struct{}

func NewUniversityFactory() *UniversityFactory {
	return &UniversityFactory{}
}

func (f *UniversityFactory) CreateController(db *gorm.DB, validate *validator.Validate) interface{} {
	repo := f.CreateRepository(db)
	service := f.CreateService(repo, validate)
	return university.NewUniversityController(service.(*university.Service))
}

func (f *UniversityFactory) CreateRepository(db *gorm.DB) interface{} {
	return university.NewUniversityRepository(db)
}

func (f *UniversityFactory) CreateService(repo interface{}, validate *validator.Validate) interface{} {
	return university.NewUniversityService(repo.(*university.Repository), validate)
}
