package container

import (
	"backend-go/internal/container/interfaces"
	"backend-go/internal/container/locator"
	"backend-go/internal/container/registry"
	"backend-go/internal/modules/athletic"
	"backend-go/internal/modules/championship"
	"backend-go/internal/modules/sport"
	"backend-go/internal/modules/university"
	"backend-go/internal/modules/user"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AppContainer struct {
	db              *gorm.DB
	validator       *validator.Validate
	serviceLocator  *locator.ServiceLocator
	factoryRegistry *registry.FactoryRegistry
}

func NewAppContainer(db *gorm.DB) interfaces.Container {
	container := &AppContainer{
		db:              db,
		validator:       validator.New(),
		serviceLocator:  locator.NewServiceLocator(),
		factoryRegistry: registry.NewFactoryRegistry(),
	}

	container.initializeFactories()
	container.initializeControllers()

	return container
}

func (c *AppContainer) initializeFactories() {
	c.factoryRegistry.RegisterDefaultFactories()
}

func (c *AppContainer) initializeControllers() {
	controllers := []string{"user", "university", "athletic", "championship", "sport"}

	for _, controllerName := range controllers {
		factory, err := c.factoryRegistry.Get(controllerName)
		if err != nil {
			panic(err)
		}

		controller := factory.CreateController(c.db, c.validator)
		c.serviceLocator.Register(controllerName+"Controller", controller)
	}
}

func (c *AppContainer) Register(name string, factory func() interface{}) {
	service := factory()
	c.serviceLocator.Register(name, service)
}

func (c *AppContainer) Get(name string) interface{} {
	service, err := c.serviceLocator.Get(name)
	if err != nil {
		panic(err)
	}
	return service
}

func (c *AppContainer) GetUserController() interface{} {
	return c.Get("userController").(*user.Controller)
}

func (c *AppContainer) GetUniversityController() interface{} {
	return c.Get("universityController").(*university.Controller)
}

func (c *AppContainer) GetAthleticController() interface{} {
	return c.Get("athleticController").(*athletic.Controller)
}

func (c *AppContainer) GetChampionshipController() interface{} {
	return c.Get("championshipController").(*championship.Controller)
}

func (c *AppContainer) GetSportController() interface{} {
	return c.Get("sportController").(*sport.Controller)
}
