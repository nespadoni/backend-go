package locator

import (
	"fmt"
	"sync"
)

type ServiceLocator struct {
	services map[string]interface{}
	mutex    sync.RWMutex
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		services: make(map[string]interface{}),
	}
}

func (sl *ServiceLocator) Register(name string, service interface{}) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	sl.services[name] = service
}

func (sl *ServiceLocator) Get(name string) (interface{}, error) {
	sl.mutex.RLock()
	defer sl.mutex.RUnlock()

	service, exists := sl.services[name]
	if !exists {
		return nil, fmt.Errorf("service %s não encontrado", name)
	}
	return service, nil
}

func (sl *ServiceLocator) Has(name string) bool {
	sl.mutex.RLock()
	defer sl.mutex.RUnlock()
	_, exists := sl.services[name]
	return exists
}
