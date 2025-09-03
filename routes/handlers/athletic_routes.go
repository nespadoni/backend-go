package handlers

import (
	"backend-go/internal/modules/athletic"

	"github.com/gin-gonic/gin"
)

type AthleticRoutes struct {
	controller *athletic.Controller
}

func NewAthleticRoutes(controller *athletic.Controller) *AthleticRoutes {
	return &AthleticRoutes{controller: controller}
}

func (ar *AthleticRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	athleticGroup := rg.Group("/athletics")
	{
		athleticGroup.GET("/", ar.controller.FindAll)
		athleticGroup.GET("/:id", ar.controller.FindById)
		athleticGroup.POST("/", ar.controller.Create)
		athleticGroup.PUT("/:id", ar.controller.Update)
		athleticGroup.PATCH("/:id/status", ar.controller.UpdateStatus)
		athleticGroup.DELETE("/:id", ar.controller.Delete)
	}
}
