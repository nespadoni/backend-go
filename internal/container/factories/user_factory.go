package factories

import (
	"backend-go/internal/modules/user"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (f *UserFactory) CreateController(db *gorm.DB, validate *validator.Validate) interface{} {
	repo := f.CreateRepository(db)
	service := f.CreateService(repo, validate)
	return user.NewUserController(service.(*user.Service))
}

func (f *UserFactory) CreateRepository(db *gorm.DB) interface{} {
	return user.NewUserRepository(db)
}

func (f *UserFactory) CreateService(repo interface{}, validate *validator.Validate) interface{} {
	return user.NewUserService(repo.(*user.Repository), validate)
}
