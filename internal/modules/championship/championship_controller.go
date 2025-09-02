package championship

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func NewChampionshipController(service *Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) FindAll(ctx *gin.Context) {

	championships, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "fetch_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, championships)
}

func (c *Controller) FindById(ctx *gin.Context) {
	championshipIdStr := ctx.Param("id")

	championshipId, err := strconv.ParseUint(championshipIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}

	response, err := c.service.FindById(uint(championshipId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "user_not_found",
			Message: "Usuário não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) Create(ctx *gin.Context) {
	var newChampionship CreateRequest
	if err := ctx.ShouldBindJSON(&newChampionship); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados do campeonato inválidos",
		})
		return
	}

	response, err := c.service.Create(newChampionship)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "creation_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) Update(ctx *gin.Context) {
	championshipIdStr := ctx.Param("id")
	championshipId, err := strconv.ParseUint(championshipIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_id",
			Message: "ID inválido",
		})
		return
	}

	var updateRequest UpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados de atualização inválidos",
		})
		return
	}

	championship, err := c.service.Update(uint(championshipId), updateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, championship)
}

func (c *Controller) Delete(ctx *gin.Context) {
	championshipIdStr := ctx.Param("id")
	championshipId, err := strconv.ParseUint(championshipIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_id",
			Message: "ID inválido",
		})
		return
	}

	if err := c.service.Delete(uint(championshipId)); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "championship_not_found",
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
