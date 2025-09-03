package interfaces

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ModuleFactory interface {
	CreateController(db *gorm.DB, validate *validator.Validate) interface{}
}

type RepositoryFactory interface {
	CreateRepository(db *gorm.DB) interface{}
}

type ServiceFactory interface {
	CreateService(repo interface{}, validate *validator.Validate) interface{}
}
