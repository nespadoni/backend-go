package routes

import (
	"backend-go/config"
	"backend-go/docs"
	"backend-go/pkg/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	config *config.Config
}

func NewServerConfig(cfg *config.Config) *ServerConfig {
	return &ServerConfig{config: cfg}
}

func (sc *ServerConfig) SetupEngine() *gin.Engine {
	// Configurar modo do Gin
	if sc.config.Port == "80" || sc.config.Port == "443" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Middlewares globais
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	return r
}

func (sc *ServerConfig) SetupSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = "localhost" + sc.config.Port

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (sc *ServerConfig) SetupHealthCheck(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API está funcionando!",
		})
	})
}
