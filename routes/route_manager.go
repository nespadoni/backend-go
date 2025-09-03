package routes

import (
	"backend-go/internal/auth"
	"backend-go/internal/container/interfaces"
	"backend-go/internal/modules/athletic"
	"backend-go/internal/modules/championship"
	"backend-go/internal/modules/sport"
	"backend-go/internal/modules/university"
	"backend-go/internal/modules/user"
	routeHandlers "backend-go/routes/handlers"
	routeInterfaces "backend-go/routes/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteManager struct {
	container      interfaces.Container
	authMiddleware *auth.Middleware
	handlers       []routeInterfaces.RouteHandler
}

func NewRouteManager(container interfaces.Container, db *gorm.DB) *RouteManager {
	permissionService := auth.NewPermissionService(db) // Ajustar conforme necessário
	authMiddleware := auth.NewMiddleware(permissionService)

	rm := &RouteManager{
		container:      container,
		authMiddleware: authMiddleware,
		handlers:       make([]routeInterfaces.RouteHandler, 0),
	}

	rm.initializeHandlers()
	return rm
}

func (rm *RouteManager) initializeHandlers() {
	rm.handlers = []routeInterfaces.RouteHandler{
		routeHandlers.NewUserRoutes(
			rm.container.GetUserController().(*user.Controller),
			rm.authMiddleware,
		),
		routeHandlers.NewUniversityRoutes(
			rm.container.GetUniversityController().(*university.Controller),
		),
		routeHandlers.NewAthleticRoutes(
			rm.container.GetAthleticController().(*athletic.Controller),
		),
		routeHandlers.NewChampionshipRoutes(
			rm.container.GetChampionshipController().(*championship.Controller),
		),
		routeHandlers.NewSportRoutes(
			rm.container.GetSportController().(*sport.Controller),
		),
	}
}

func (rm *RouteManager) RegisterRoutes(api *gin.RouterGroup) {
	for _, handler := range rm.handlers {
		handler.RegisterRoutes(api)
	}
}
