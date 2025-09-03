package registry

import (
	"backend-go/internal/container/factories"
	"backend-go/internal/container/interfaces"
	"fmt"
)

type FactoryRegistry struct {
	factories map[string]interfaces.ModuleFactory
}

func NewFactoryRegistry() *FactoryRegistry {
	return &FactoryRegistry{
		factories: make(map[string]interfaces.ModuleFactory),
	}
}

func (r *FactoryRegistry) Register(name string, factory interfaces.ModuleFactory) {
	r.factories[name] = factory
}

func (r *FactoryRegistry) Get(name string) (interfaces.ModuleFactory, error) {
	factory, exists := r.factories[name]
	if !exists {
		return nil, fmt.Errorf("factory %s não encontrada", name)
	}
	return factory, nil
}

func (r *FactoryRegistry) RegisterDefaultFactories() {
	r.Register("user", factories.NewUserFactory())
	r.Register("university", factories.NewUniversityFactory())
	r.Register("athletic", factories.NewAthleticFactory())
	r.Register("championship", factories.NewChampionshipFactory())
	r.Register("sport", factories.NewSportFactory())
}
