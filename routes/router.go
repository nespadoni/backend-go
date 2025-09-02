package routes

import (
	"backend-go/config"
	"backend-go/docs"
	"backend-go/internal/auth"
	"backend-go/internal/modules/athletic"
	championship2 "backend-go/internal/modules/championship"
	"backend-go/internal/modules/university"
	user2 "backend-go/internal/modules/user"
	"backend-go/pkg/middleware"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// InitRouter
// @title Backend Rivaly API
// @version 1.0
// @description Esta é a API do Rivaly desenvolvida em Go + Gin
// @host localhost:8080
// @BasePath /
func InitRouter(database *gorm.DB, cfg *config.Config) {
	if cfg.Port == "80" || cfg.Port == "443" {
		gin.SetMode(gin.ReleaseMode) // Produção
	}

	r := gin.New()

	// Middlewares
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// Controladores
	championController := startChampionship(database)
	userController := startUser(database)
	athleticController := startAthletic(database)
	universityController := startUniversity(database)

	permissionService := auth.NewPermissionService(database)
	authMiddleware := auth.NewMiddleware(permissionService)

	// Configuração do Swagger
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = "localhost" + cfg.Port

	// Rotas da API
	api := r.Group("/api")
	{
		// Rotas de Universidades
		universityRoutes := api.Group("/universities")
		{
			universityRoutes.GET("/", universityController.FindAll)
			universityRoutes.GET("/:id", universityController.FindById)
			universityRoutes.POST("/", universityController.Create)
			universityRoutes.PUT("/:id", universityController.Update)
			universityRoutes.DELETE("/:id", universityController.Delete)
		}

		// Rotas de Usuários
		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/", userController.FindAll)
			userRoutes.GET("/:id", userController.FindById, authMiddleware.RequireLevel(80))
			userRoutes.POST("/", userController.PostUser)
			userRoutes.PUT("/:id", userController.UpdateUser)
			userRoutes.DELETE("/:id", userController.DeleteUser)
		}

		// Rotas de Atléticas
		athleticRoutes := api.Group("/athletics")
		{
			athleticRoutes.GET("/", athleticController.FindAll)
			athleticRoutes.GET("/:id", athleticController.FindById)
			athleticRoutes.POST("/", athleticController.Create)
			athleticRoutes.PUT("/:id", athleticController.Update)
			athleticRoutes.PATCH("/:id/status", athleticController.UpdateStatus)
			athleticRoutes.DELETE("/:id", athleticController.Delete)
		}

		// Rotas de Campeonatos
		championshipRoutes := api.Group("/championships")
		{
			championshipRoutes.GET("/", championController.FindAll)
			championshipRoutes.GET("/:id", championController.FindById)
			championshipRoutes.POST("/", championController.Create)
			championshipRoutes.PUT("/:id", championController.Update)
			championshipRoutes.PATCH("/:id/status", championController.UpdateStatus)
			championshipRoutes.DELETE("/:id", championController.Delete)
		}
	}

	// Rota do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API está funcionando!",
		})
	})

	// Inicia o servidor
	log.Printf("Servidor rodando em http://localhost:%s", cfg.Port)
	log.Printf("Swagger disponível em http://localhost:%s/swagger/index.html", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}

func startChampionship(database *gorm.DB) championship2.Controller {
	validate := validator.New()
	championRepo := championship2.NewChampionshipRepository(database)
	championService := championship2.NewChampionshipService(championRepo, validate)
	championController := championship2.NewChampionshipController(championService)

	return *championController
}

func startUser(database *gorm.DB) user2.Controller {
	validate := validator.New()
	userRepo := user2.NewUserRepository(database)
	userService := user2.NewUserService(userRepo, validate)
	userController := user2.NewUserController(userService)

	return *userController
}

func startAthletic(database *gorm.DB) athletic.Controller {
	validate := validator.New()
	athleticRepo := athletic.NewAthleticRepository(database)
	athleticService := athletic.NewAthleticService(athleticRepo, validate)
	athleticController := athletic.NewAthleticController(athleticService)

	return *athleticController
}

func startUniversity(database *gorm.DB) university.Controller {
	validate := validator.New()
	universityRepo := university.NewUniversityRepository(database)
	universityService := university.NewUniversityService(universityRepo, validate)
	universityController := university.NewUniversityController(universityService)

	return *universityController
}
