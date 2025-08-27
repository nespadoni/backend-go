package championship

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewChampionshipController(service *Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) FindAll(ctx *gin.Context) {

	championship, err := controller.service.FindAll()
	if err != nil {

		ctx.JSON(http.StatusBadRequest, championship)
	}

	ctx.JSON(http.StatusOK, championship)
}
