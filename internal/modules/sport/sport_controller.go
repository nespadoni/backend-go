package sport

import (
	"backend-go/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	sportService *Service
}

func NewSportController(sportService *Service) *Controller {
	return &Controller{sportService: sportService}
}

// FindAll godoc
// @Summary Lista todos os esportes
// @Description Retorna uma lista de todos os esportes cadastrados
// @Tags sports
// @Accept json
// @Produce json
// @Success 200 {array} ListResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/sports [get]
func (c *Controller) FindAll(ctx *gin.Context) {
	sports, err := c.sportService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}

	ctx.JSON(http.StatusOK, sports)
}

// FindPopular godoc
// @Summary Lista esportes populares
// @Description Retorna uma lista dos esportes marcados como populares e ativos
// @Tags sports
// @Accept json
// @Produce json
// @Success 200 {array} ListResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/sports/popular [get]
func (c *Controller) FindPopular(ctx *gin.Context) {
	sports, err := c.sportService.FindPopular()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}

	ctx.JSON(http.StatusOK, sports)
}

// FindById godoc
// @Summary Busca esporte por ID
// @Description Retorna um esporte específico pelo seu ID
// @Tags sports
// @Accept json
// @Produce json
// @Param id path string true "Sport ID"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/sports/{id} [get]
func (c *Controller) FindById(ctx *gin.Context) {
	sportIDStr := ctx.Param("id")
	sportID, err := strconv.ParseUint(sportIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_sport_id",
			Message: "ID do esporte deve ser um número válido",
		})
		return
	}

	response, err := c.sportService.FindById(uint(sportID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "sport_not_found",
			Message: "Esporte não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Cria um novo esporte
// @Description Cria um novo esporte com os dados fornecidos
// @Tags sports
// @Accept json
// @Produce json
// @Param sport body CreateRequest true "Sport data"
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/sports [post]
func (c *Controller) Create(ctx *gin.Context) {
	var newSport CreateRequest
	if err := ctx.ShouldBindJSON(&newSport); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados do esporte inválidos",
		})
		return
	}

	response, err := c.sportService.Create(newSport)
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
// @Summary Atualiza um esporte
// @Description Atualiza os dados de um esporte existente
// @Tags sports
// @Accept json
// @Produce json
// @Param id path string true "Sport ID"
// @Param sport body UpdateRequest true "Sport data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/sports/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	sportIDStr := ctx.Param("id")
	sportID, err := strconv.ParseUint(sportIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_sport_id",
			Message: "ID do esporte deve ser um número válido",
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

	sport, err := c.sportService.Update(uint(sportID), request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, sport)
}

// UpdateStatus godoc
// @Summary Atualiza status do esporte
// @Description Atualiza o status ativo/popular de um esporte
// @Tags sports
// @Accept json
// @Produce json
// @Param id path string true "Sport ID"
// @Param status body UpdateStatusRequest true "Status data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/sports/{id}/status [patch]
func (c *Controller) UpdateStatus(ctx *gin.Context) {
	sportIDStr := ctx.Param("id")
	sportID, err := strconv.ParseUint(sportIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_sport_id",
			Message: "ID do esporte deve ser um número válido",
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

	sport, err := c.sportService.UpdateStatus(uint(sportID), request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_status_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, sport)
}

// Delete godoc
// @Summary Deleta um esporte
// @Description Remove um esporte do sistema
// @Tags sports
// @Accept json
// @Produce json
// @Param id path string true "Sport ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/sports/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	sportIDStr := ctx.Param("id")
	sportID, err := strconv.ParseUint(sportIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_sport_id",
			Message: "ID do esporte deve ser um número válido",
		})
		return
	}

	if err := c.sportService.Delete(uint(sportID)); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "sport_not_found",
			Message: "Esporte não encontrado ou erro ao deletar",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
