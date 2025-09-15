package routes

import (
	"backend-go/config"
	"backend-go/docs"
	"backend-go/internal/auth"
	"backend-go/internal/modules/athletic"
	"backend-go/internal/modules/championship"
	"backend-go/internal/modules/sport"
	"backend-go/internal/modules/university"
	"backend-go/internal/modules/user"
	"backend-go/pkg/middleware"
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200", "http://127.0.0.1:4200", "https://rivaly.up.railway.app",
			"https://front-web.railway.internal"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    false,
		MaxAge:           12 * time.Hour,
	}))

	// Middleware adicional para tratar preflight requests
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "https://rivaly.up.railway.app")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Authorization")
			c.Header("Access-Control-Max-Age", "86400")
			c.Status(200)
			return
		}
		c.Next()
	})

	// Controladores
	validate := validator.New()

	userRepo := user.NewUserRepository(database)
	authService := auth.NewAuthService(userRepo, validate, cfg.JWTSecret)
	authController := auth.NewAuthController(authService)

	championController := startChampionship(database)
	userController := startUser(database, userRepo, validate)
	universityController := startUniversity(database)
	athleticController := startAthletic(database)
	sportController := startSport(database)
	//permissionService := auth.NewPermissionService(database)
	//authMiddleware := auth.NewMiddleware(permissionService)

	// Configuração do Swagger para Railway
	docs.SwaggerInfo.BasePath = "/api"
	if cfg.Port == "80" || cfg.Port == "443" {
		docs.SwaggerInfo.Host = "backend-go-production-c4f4.up.railway.app"
	} else {
		docs.SwaggerInfo.Host = "localhost:" + cfg.Port
	}

	// Rotas da API
	api := r.Group("/api")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
		}

		// Grupo de rotas protegidas pelo middleware JWT
		authorized := api.Group("/")
		authorized.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			championshipRoutes := authorized.Group("/championships")
			{
				championshipRoutes.POST("/", championController.Create)
				championshipRoutes.PUT("/:id", championController.Update)
				championshipRoutes.PATCH("/:id/status", championController.UpdateStatus)
				championshipRoutes.DELETE("/:id", championController.Delete)
			}

			userRoutes := authorized.Group("/users")
			{
				userRoutes.GET("/:id", userController.FindById)
				userRoutes.DELETE("/:id", userController.DeleteUser)
				userRoutes.PUT("/:id", userController.UpdateUser)
				userRoutes.POST("/profile-photo", userController.UploadProfilePhoto) // Nova rota
			}

			universityRoutes := authorized.Group("/universities")
			{
				universityRoutes.POST("/", universityController.Create)
				universityRoutes.PUT("/:id", universityController.Update)
				universityRoutes.DELETE("/:id", universityController.Delete)
			}

			athleticRoutes := authorized.Group("/athletics")
			{
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

		// Rotas públicas (que não precisam de login)
		api.GET("/championships", championController.FindAll)
		api.GET("/championships/:id", championController.FindById)
		api.GET("/users", userController.FindAll)
		api.GET("/universities", universityController.FindAll)
		api.GET("/universities/:id", universityController.FindById)
		api.GET("/athletics", athleticController.FindAll)
		api.GET("/athletics/:id", athleticController.FindById)
	}

	// Rota do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Static("/uploads", "./uploads")

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

func startUser(database *gorm.DB, userRepo *user.Repository, validate *validator.Validate) user.Controller {
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
