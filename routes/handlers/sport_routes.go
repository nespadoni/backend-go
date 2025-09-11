package handlers

import (
	"backend-go/internal/modules/sport"

	"github.com/gin-gonic/gin"
)

type SportRoutes struct {
	controller *sport.Controller
}

func NewSportRoutes(controller *sport.Controller) *SportRoutes {
	return &SportRoutes{controller: controller}
}

func (sr *SportRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	sportGroup := rg.Group("/sports")
	{
		sportGroup.GET("/", sr.controller.FindAll)
		sportGroup.GET("/popular", sr.controller.FindPopular)
		sportGroup.GET("/:id", sr.controller.FindById)
		sportGroup.POST("/", sr.controller.Create)
		sportGroup.PUT("/:id", sr.controller.Update)
		sportGroup.PATCH("/:id/status", sr.controller.UpdateStatus)
		sportGroup.DELETE("/:id", sr.controller.Delete)
	}
}
