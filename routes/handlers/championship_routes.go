package handlers

import (
	"backend-go/internal/modules/championship"

	"github.com/gin-gonic/gin"
)

type ChampionshipRoutes struct {
	controller *championship.Controller
}

func NewChampionshipRoutes(controller *championship.Controller) *ChampionshipRoutes {
	return &ChampionshipRoutes{controller: controller}
}

func (cr *ChampionshipRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	championshipGroup := rg.Group("/championships")
	{
		championshipGroup.GET("/", cr.controller.FindAll)
		championshipGroup.GET("/:id", cr.controller.FindById)
		championshipGroup.POST("/", cr.controller.Create)
		championshipGroup.PUT("/:id", cr.controller.Update)
		championshipGroup.PATCH("/:id/status", cr.controller.UpdateStatus)
		championshipGroup.DELETE("/:id", cr.controller.Delete)
	}
}
