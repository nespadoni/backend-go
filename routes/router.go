package routes

import (
	"backend-go/internal/championship"
	"backend-go/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	database := db.InitDB()

	championRepo := championship.NewChampionshipRepository(database)
	championService := championship.NewChampionshipService(championRepo)
	championController := championship.NewChampionshipController(championService)

	r := gin.Default()

	r.GET("/championship", func(ctx *gin.Context) {
		championship, err := championController.GetChampionship()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, championship)
		}
		ctx.JSON(http.StatusOK, championship)
	})

	r.Run()
}
