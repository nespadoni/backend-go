package championship

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChampionshipController struct {
	service *ChampionshipService
}

func NewChampionshipController(service *ChampionshipService) *ChampionshipController {
	return &ChampionshipController{service: service}
}

func (controller *ChampionshipController) GetChampionship(ctx *gin.Context) {

	championship, err := controller.service.FindChampionship()
	if err != nil {

		ctx.JSON(http.StatusBadRequest, championship)
	}

	ctx.JSON(http.StatusOK, championship)
}
