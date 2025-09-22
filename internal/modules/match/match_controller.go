package match

import (
	"backend-go/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewMatchController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) FindAll(ctx *gin.Context) {
	matches, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, matches)
}

func (c *Controller) FindById(ctx *gin.Context) {
	matchesIDStr := ctx.Param("id")
	matchesID, err := strconv.ParseUint(matchesIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "invalid_match_id",
			Message: "ID da partida deve ser um número válido",
		})
		return
	}

	response, err := c.service.FindByID(uint(matchesID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "match_not_found",
			Message: "Partida não encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) Create(ctx *gin.Context) {
	var newMatch CreateRequest
	if err := ctx.ShouldBindJSON(&newMatch); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados da partida inválidos",
		})
		return
	}

	response, err := c.service.Create(newMatch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "creation_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) Update(ctx *gin.Context) {
	matchIDStr := ctx.Param("id")
	matchID, err := strconv.ParseUint(matchIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "invalid_match_id",
			Message: "ID da partida deve ser um número válido",
		})
		return
	}

	var request UpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados de atualização inválidos",
		})
		return
	}

	match, err := c.service.Update(uint(matchID), request)
	if err != nil {
		if err.Error() == "partida não encontrada" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "match_not_found",
				Message: "Partida não encontrada",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, match)
}

func (c *Controller) Delete(ctx *gin.Context) {
	matchIDStr := ctx.Param("id")
	matchID, err := strconv.ParseUint(matchIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "invalid_match_id",
			Message: "ID da universidade deve ser um número válido",
		})
		return
	}

	if err := c.service.Delete(uint(matchID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "delete_match_failed",
			Message: "Falha ao deletar partida",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
