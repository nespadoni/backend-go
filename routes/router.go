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

	// Middlewares básicos
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// CORS configuração corrigida para Railway
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4200",
			"http://127.0.0.1:4200",
			"https://rivaly.up.railway.app",
			"https://front-web.railway.internal", // Adicionar domínio interno
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Accept",
			"Accept-Encoding",
			"Cache-Control",
			"Connection",
			"DNT",
			"Host",
			"Pragma",
			"Referer",
			"User-Agent",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
		},
		AllowCredentials: true,
		AllowWildcard:    false, // Importante para Railway
		MaxAge:           12 * time.Hour,
	}))

	// Middleware adicional para tratar preflight requests
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "https://localhost:5432")
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

	// Configuração do Swagger para Railway
	docs.SwaggerInfo.BasePath = "/api"
	if cfg.Port == "80" || cfg.Port == "443" {
		docs.SwaggerInfo.Host = "backend-go-production-c4f4.up.railway.app"
	} else {
		docs.SwaggerInfo.Host = "localhost:" + cfg.Port
	}

	// Health check endpoint (primeira rota para teste)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API está funcionando!",
			"cors":    "enabled",
		})
	})

	// Rotas da API
	api := r.Group("/api")
	{
		// Rotas de autenticação
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
		}

		// ROTAS PÚBLICAS (SEM AUTENTICAÇÃO) - MOVIDAS PARA CIMA
		api.GET("/championships", championController.FindAll)
		api.GET("/championships/:id", championController.FindById)
		api.GET("/users", userController.FindAll)

		// ROTA UNIVERSITIES PÚBLICA (ESTA É A QUE ESTÁ FALHANDO)
		api.GET("/universities", universityController.FindAll)
		api.GET("/universities/:id", universityController.FindById)

		api.GET("/athletics", athleticController.FindAll)
		api.GET("/athletics/:id", athleticController.FindById)

		// Rotas de esportes públicas
		api.GET("/sports", sportController.FindAll)
		api.GET("/sports/popular", sportController.FindPopular)
		api.GET("/sports/:id", sportController.FindById)

		usersRoutes := api.Group("/users")
		{
			usersRoutes.GET("/:id", userController.FindById)
			usersRoutes.PUT("/:id", userController.UpdateUser)
			usersRoutes.POST("/profile-photo", userController.UploadProfilePhoto)
			usersRoutes.DELETE("/:id", userController.DeleteUser)
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

			sportRoutes := authorized.Group("/sports")
			{
				sportRoutes.POST("/", sportController.Create)
				sportRoutes.PUT("/:id", sportController.Update)
				sportRoutes.PATCH("/:id/status", sportController.UpdateStatus)
				sportRoutes.DELETE("/:id", sportController.Delete)
			}
		}
	}

	// Rota do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Servir arquivos estáticos
	r.Static("/uploads", "./uploads")

	// Log das configurações
	log.Printf("CORS configurado para: http://localhost:8080")
	log.Printf("Servidor rodando na porta: %s", cfg.Port)

	if cfg.Port == "80" || cfg.Port == "443" {
		log.Printf("Swagger disponível em https://backend-go-production-c4f4.up.railway.app/swagger/index.html")
	} else {
		log.Printf("Swagger disponível em http://localhost:%s/swagger/index.html", cfg.Port)
	}

	// Inicia o servidor
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
