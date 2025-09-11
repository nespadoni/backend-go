package routes

import (
	"backend-go/config"
	"backend-go/docs"
	"backend-go/internal/modules/athletic"
	"backend-go/internal/modules/championship"
	"backend-go/internal/modules/sport"
	"backend-go/internal/modules/university"
	"backend-go/internal/modules/user"
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
	universityController := startUniversity(database)
	athleticController := startAthletic(database)
	sportController := startSport(database)
	//permissionService := auth.NewPermissionService(database)
	//authMiddleware := auth.NewMiddleware(permissionService)

	// Configuração do Swagger
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = "localhost" + cfg.Port

	// Rotas da API
	api := r.Group("/api")
	{

		championshipRoutes := api.Group("/championships")
		{
			championshipRoutes.GET("/", championController.FindAll)
			championshipRoutes.GET("/:id", championController.FindById)
			championshipRoutes.POST("/", championController.Create)
			championshipRoutes.PUT("/:id", championController.Update)
			championshipRoutes.PATCH("/:id/status", championController.UpdateStatus)
			championshipRoutes.DELETE("/:id", championController.Delete)
		}

		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/:id", userController.FindById)
			userRoutes.GET("/", userController.FindAll)
			userRoutes.DELETE("/:id", userController.DeleteUser)
			userRoutes.POST("/", userController.PostUser)
			userRoutes.PUT("/:id", userController.UpdateUser)
		}

		universityRoutes := api.Group("/universities")
		{
			universityRoutes.GET("/", universityController.FindAll)
			universityRoutes.GET("/:id", universityController.FindById)
			universityRoutes.POST("/", universityController.Create)
			universityRoutes.PUT("/:id", universityController.Update)
			universityRoutes.DELETE("/:id", universityController.Delete)
		}

		athleticRoutes := api.Group("/athletics")
		{
			athleticRoutes.GET("/", athleticController.FindAll)
			athleticRoutes.GET("/:id", athleticController.FindById)
			athleticRoutes.POST("/", athleticController.Create)
			athleticRoutes.PUT("/:id", athleticController.Update)
			athleticRoutes.PATCH("/:id/status", athleticController.UpdateStatus)
			athleticRoutes.DELETE("/:id", athleticController.Delete)
		}

		sportRoutes := api.Group("/sports")
		{
			sportRoutes.GET("/", sportController.FindAll)
			sportRoutes.GET("/popular", sportController.FindPopular)
			sportRoutes.GET("/:id", sportController.FindById)
			sportRoutes.POST("/", sportController.Create)
			sportRoutes.PUT("/:id", sportController.Update)
			sportRoutes.PATCH("/:id/status", sportController.UpdateStatus)
			sportRoutes.DELETE("/:id", sportController.Delete)
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

func startChampionship(database *gorm.DB) championship.Controller {
	validate := validator.New()
	championRepo := championship.NewChampionshipRepository(database)
	championService := championship.NewChampionshipService(championRepo, validate)
	championController := championship.NewChampionshipController(championService)

	return *championController
}

func startUser(database *gorm.DB) user.Controller {
	validate := validator.New()
	userRepo := user.NewUserRepository(database)
	userService := user.NewUserService(userRepo, validate)
	userController := user.NewUserController(userService)

	return *userController
}

func startUniversity(database *gorm.DB) university.Controller {
	validate := validator.New()
	repo := university.NewUniversityRepository(database)
	service := university.NewUniversityService(repo, validate)
	controller := university.NewUniversityController(service)

	return *controller
}

func startAthletic(database *gorm.DB) athletic.Controller {
	validate := validator.New()
	repo := athletic.NewAthleticRepository(database)
	service := athletic.NewAthleticService(repo, validate)
	controller := athletic.NewAthleticController(service)

	return *controller
}

func startSport(database *gorm.DB) sport.Controller {
	validate := validator.New()
	repo := sport.NewSportRepository(database)
	service := sport.NewSportService(repo, validate)
	controller := sport.NewSportController(service)

	return *controller
}
