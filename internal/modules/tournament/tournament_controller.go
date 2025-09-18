package tournament

import (
	"backend-go/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewTornamentController(service *Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) FindAll(ctx *gin.Context) {
	tournament, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}
	ctx.JSON(http.StatusOK, tournament)
}

func (c *Controller) FindById(ctx *gin.Context) {
	tournamentIDStr := ctx.Param("id")
	tournamentID, err := strconv.ParseUint(tournamentIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "invalid_tournament_id",
			Message: "ID do torneio deve ser um número válido",
		})
		return
	}

	response, err := c.service.FindByID(uint(tournamentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
