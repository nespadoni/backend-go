package championship

import (
	"backend-go/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	championshipService *Service
}

func NewChampionshipController(championshipService *Service) *Controller {
	return &Controller{championshipService: championshipService}
}

func (c *Controller) FindAll(ctx *gin.Context) {
	championships, err := c.championshipService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}
	ctx.JSON(http.StatusOK, championships)
}

func (c *Controller) FindById(ctx *gin.Context) {
	championshipIDStr := ctx.Param("id")
	championshipID, err := strconv.ParseUint(championshipIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "ID do campeonato deve ser um número válido",
		})
		return
	}
	response, err := c.championshipService.FindById(uint(championshipID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "championship_not_found",
			Message: "Campeonato não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) Create(ctx *gin.Context) {
	var newChampionship CreateRequest
	if err := ctx.ShouldBindJSON(&newChampionship); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados do campeonato inválidos",
		})
		return
	}

	response, err := c.championshipService.Create(newChampionship)
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
	championshipIDStr := ctx.Param("id")

	championshipId, err := strconv.ParseUint(championshipIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "ID do campeonato deve ser um número válido",
		})
		return
	}

	var request UpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados de atualização inválidos",
		})
		return
	}

	championship, err := c.championshipService.Update(uint(championshipId), request)
	if err != nil {
		if err.Error() == "campeonato não encontrado" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "championship_not_found",
				Message: "Campeonato não encontrado",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, championship)
}
