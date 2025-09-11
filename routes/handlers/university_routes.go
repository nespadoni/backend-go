package handlers

import (
	"backend-go/internal/modules/university"

	"github.com/gin-gonic/gin"
)

type UniversityRoutes struct {
	controller *university.Controller
}

func NewUniversityRoutes(controller *university.Controller) *UniversityRoutes {
	return &UniversityRoutes{controller: controller}
}

func (ur *UniversityRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	universityGroup := rg.Group("/universities")
	{
		universityGroup.GET("/", ur.controller.FindAll)
		universityGroup.GET("/:id", ur.controller.FindById)
		universityGroup.POST("/", ur.controller.Create)
		universityGroup.PUT("/:id", ur.controller.Update)
		universityGroup.DELETE("/:id", ur.controller.Delete)
	}
}
