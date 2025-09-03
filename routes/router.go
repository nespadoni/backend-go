package routes

import (
	"backend-go/config"
	"backend-go/internal/container"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(database *gorm.DB, cfg *config.Config) {
	// 1. Inicializar container
	appContainer := container.NewAppContainer(database)

	// 2. Configurar servidor
	serverConfig := NewServerConfig(cfg)
	r := serverConfig.SetupEngine()

	// 3. Configurar rotas
	routeManager := NewRouteManager(appContainer, database)
	api := r.Group("/api")
	routeManager.RegisterRoutes(api)

	// 4. Configurar recursos adicionais
	serverConfig.SetupSwagger(r)
	serverConfig.SetupHealthCheck(r)

	// 5. Iniciar servidor
	startServer(r, cfg)
}

func startServer(r *gin.Engine, cfg *config.Config) {
	log.Printf("Servidor rodando em http://localhost:%s", cfg.Port)
	log.Printf("Swagger disponível em http://localhost:%s/swagger/index.html", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
