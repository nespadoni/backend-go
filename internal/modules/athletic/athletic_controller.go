package athletic

import (
	"backend-go/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	athleticService *Service
}

func NewAthleticController(athleticService *Service) *Controller {
	return &Controller{athleticService: athleticService}
}

// FindAll godoc
// @Summary Lista todas as atléticas
// @Description Retorna lista de atléticas cadastradas
// @Tags athletics
// @Accept json
// @Produce json
// @Success 200 {array} ListResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/athletics [get]
func (c *Controller) FindAll(ctx *gin.Context) {
	athletics, err := c.athleticService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}

	ctx.JSON(http.StatusOK, athletics)
}

// FindById godoc
// @Summary Busca atlética por Id
// @Description Retorna atlética específica pelo Id
// @Tags athletics
// @Accept json
// @Produce json
// @Param id path string true "Athletic Id"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/athletics/{id} [get]
func (c *Controller) FindById(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "Id da atlética deve ser um número válido",
		})
		return
	}

	response, err := c.athleticService.FindById(uint(athleticID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "athletic_not_found",
			Message: "Atlética não encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Cria nova atlética
// @Description Cria atlética com dados fornecidos
// @Tags athletics
// @Accept json
// @Produce json
// @Param athletic body CreateRequest true "Athletic data"
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/athletics [post]
func (c *Controller) Create(ctx *gin.Context) {
	var newAthletic CreateRequest
	if err := ctx.ShouldBindJSON(&newAthletic); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados da atlética inválidos",
		})
		return
	}

	response, err := c.athleticService.Create(newAthletic)
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
// @Summary Atualiza atlética
// @Description Atualiza dados da atlética
// @Tags athletics
// @Accept json
// @Produce json
// @Param id path string true "Athletic Id"
// @Param athletic body UpdateRequest true "Athletic data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/athletics/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "Id da atlética deve ser um número válido",
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

	athletic, err := c.athleticService.Update(uint(athleticID), request)
	if err != nil {
		if err.Error() == "atlética não encontrada" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "athletic_not_found",
				Message: "Atlética não encontrada",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, athletic)
}

// UpdateStatus godoc
// @Summary Atualiza status da atlética
// @Description Atualiza status ativo/inativo
// @Tags athletics
// @Accept json
// @Produce json
// @Param id path string true "Athletic Id"
// @Param status body UpdateStatusRequest true "Status data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/athletics/{id}/status [patch]
func (c *Controller) UpdateStatus(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "Id da atlética deve ser um número válido",
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

	athletic, err := c.athleticService.UpdateStatus(uint(athleticID), request)
	if err != nil {
		if err.Error() == "atlética não encontrada" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "athletic_not_found",
				Message: "Atlética não encontrada",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_status_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, athletic)
}

// Delete godoc
// @Summary Deleta atlética
// @Description Remove atlética do sistema
// @Tags athletics
// @Accept json
// @Produce json
// @Param id path string true "Athletic Id"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/athletics/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "Id da atlética deve ser um número válido",
		})
		return
	}

	if err := c.athleticService.Delete(uint(athleticID)); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "athletic_not_found",
			Message: "Atlética não encontrada ou erro ao deletar",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
