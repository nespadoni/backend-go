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

func (c *Controller) Create(ctx *gin.Context) {
	var newTournament CreateRequest
	if err := ctx.ShouldBindJSON(&newTournament); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados do torneio inválidos",
		})
		return
	}

	response, err := c.service.Create(newTournament)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "creation_failed",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) Update(ctx *gin.Context) {
	tournamentIdStr := ctx.Param("id")

	tournamentId, err := strconv.ParseUint(tournamentIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_tournament_id",
			Message: "Id do campeonato deve ser um número válido",
		})
		return
	}

	var request UpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados de atualização inválidos",
		})
		return
	}

	tournament, err := c.service.Update(uint(tournamentId), request)
	if err != nil {
		if err.Error() == "torneio não encontrado" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "torunament_not_found",
				Message: "Torneio não encontrado",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, tournament)
}

func (c *Controller) Delete(ctx *gin.Context) {
	tournamentIDStr := ctx.Param("id")
	tournamentID, err := strconv.ParseUint(tournamentIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "invalid_tournament_id",
			Message: "Id do torneio deve ser um número válido",
		})
		return
	}

	if err := c.service.Delete(uint(tournamentID)); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "tournament_delete_failed",
			Message: "Torneio não encontrado ou erro ao deletar",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)

}
