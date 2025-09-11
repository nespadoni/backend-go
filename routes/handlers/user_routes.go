package handlers

import (
	"backend-go/internal/auth"
	"backend-go/internal/modules/user"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	controller     *user.Controller
	authMiddleware *auth.Middleware
}

func NewUserRoutes(controller *user.Controller, authMiddleware *auth.Middleware) *UserRoutes {
	return &UserRoutes{
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}

func (ur *UserRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	{
		userGroup.GET("/", ur.controller.FindAll)
		userGroup.GET("/:id", ur.controller.FindById, ur.authMiddleware.RequireLevel(80))
		userGroup.POST("/", ur.controller.PostUser)
		userGroup.PUT("/:id", ur.controller.UpdateUser)
		userGroup.DELETE("/:id", ur.controller.DeleteUser)
	}
}
