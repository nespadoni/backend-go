package routes

import (
	"backend-go/internal/championship"
	"backend-go/internal/user"
	"backend-go/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitRouter() {
	database := db.InitDB()

	championController := startChampionship(database)
	userController := startUser(database)

	r := gin.Default()

	r.GET("/championship", championController.GetChampionship)
	r.GET("/user/:id", userController.FindById)
	r.GET("/user", userController.FindAll)
	r.DELETE("/user/:id", userController.DeleteUser)
	r.POST("/user", userController.PostUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.Run()
}

func startChampionship(database *gorm.DB) championship.ChampionshipController {
	championRepo := championship.NewChampionshipRepository(database)
	championService := championship.NewChampionshipService(championRepo)
	championController := championship.NewChampionshipController(championService)

	return *championController
}

func startUser(database *gorm.DB) user.UserController {
	validate := validator.New()
	userRepo := user.NewUserRepository(database)
	userService := user.NewUserService(userRepo, validate)
	userController := user.NewUserController(userService)

	return *userController
}
