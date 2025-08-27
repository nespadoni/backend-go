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

	championship, err := c.service.FindAll()
	if err != nil {

		ctx.JSON(http.StatusBadRequest, championship)
	}

	ctx.JSON(http.StatusOK, championship)
}

func (c *Controller) FindById(ctx *gin.Context) {
	championshipIdStr := ctx.Param("id")

	championshipId, err := strconv.Atoi(championshipIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}

	response, err := c.service.FindById(championshipId)
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
	championshipIdStr := ctx.Param("Id")
	championshipId, err := strconv.Atoi(championshipIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
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

	championship, err := c.service.Update(championshipId, updateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, championship)

}
