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

// FindAll godoc
// @Summary Lista todos os campeonatos
// @Description Retorna lista de campeonatos cadastrados
// @Tags championships
// @Accept json
// @Produce json
// @Success 200 {array} ListResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/championships [get]
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

// FindById godoc
// @Summary Busca campeonato por Id
// @Description Retorna campeonato específico pelo Id
// @Tags championships
// @Accept json
// @Produce json
// @Param id path string true "Championship Id"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/championships/{id} [get]
func (c *Controller) FindById(ctx *gin.Context) {
	championshipIDStr := ctx.Param("id")
	championshipID, err := strconv.ParseUint(championshipIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_championship_id",
			Message: "Id do campeonato deve ser um número válido",
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

// Create godoc
// @Summary Cria novo campeonato
// @Description Cria campeonato com dados fornecidos
// @Tags championships
// @Accept json
// @Produce json
// @Param championship body CreateRequest true "Championship data"
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/championships [post]
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

// Update godoc
// @Summary Atualiza campeonato
// @Description Atualiza dados do campeonato
// @Tags championships
// @Accept json
// @Produce json
// @Param id path string true "Championship Id"
// @Param championship body UpdateRequest true "Championship data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/championships/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	championshipIDStr := ctx.Param("id")

	championshipId, err := strconv.ParseUint(championshipIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_championship_id",
			Message: "Id do campeonato deve ser um número válido",
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

// UpdateStatus godoc
// @Summary Atualiza status do campeonato
// @Description Atualiza status ativo/inativo
// @Tags championships
// @Accept json
// @Produce json
// @Param id path string true "Championship Id"
// @Param status body UpdateStatusRequest true "Status data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/championships/{id}/status [patch]
func (c *Controller) UpdateStatus(ctx *gin.Context) {
	championshipIDStr := ctx.Param("id")
	championshipID, err := strconv.ParseUint(championshipIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_championship_id",
			Message: "Id do campeonato deve ser um número válido",
		})
		return
	}

	var request UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados de atualização de status inválidos",
		})
		return
	}

	championship, err := c.championshipService.UpdateStatus(uint(championshipID), request)
	if err != nil {
		if err.Error() == "campeonato não encontrado" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "championship_not_found",
				Message: "Campeonato não encontrado",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_status_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, championship)
}

// Delete godoc
// @Summary Deleta campeonato
// @Description Remove campeonato do sistema
// @Tags championships
// @Accept json
// @Produce json
// @Param id path string true "Championship Id"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/championships/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	championshipIDStr := ctx.Param("id")
	championshipID, err := strconv.ParseUint(championshipIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_championship_id",
			Message: "Id do campeonato deve ser um número válido",
		})
		return
	}

	if err := c.championshipService.Delete(uint(championshipID)); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "championship_not_found",
			Message: "Campeonato não encontrado ou erro ao deletar",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
