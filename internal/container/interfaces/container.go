package interfaces

type Container interface {
	Register(name string, factory func() interface{})
	Get(name string) interface{}
	GetUserController() interface{}
	GetUniversityController() interface{}
	GetAthleticController() interface{}
	GetChampionshipController() interface{}
	GetSportController() interface{}
}
