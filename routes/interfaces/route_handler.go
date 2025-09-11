package interfaces

import "github.com/gin-gonic/gin"

type RouteHandler interface {
	RegisterRoutes(rg *gin.RouterGroup)
}
