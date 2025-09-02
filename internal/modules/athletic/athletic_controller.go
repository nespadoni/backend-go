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

func (c *Controller) FindById(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "ID da atlética deve ser um número válido",
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

func (c *Controller) Update(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "ID da atlética deve ser um número válido",
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

func (c *Controller) UpdateStatus(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "ID da atlética deve ser um número válido",
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

func (c *Controller) Delete(ctx *gin.Context) {
	athleticIDStr := ctx.Param("id")
	athleticID, err := strconv.ParseUint(athleticIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_athletic_id",
			Message: "ID da atlética deve ser um número válido",
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
