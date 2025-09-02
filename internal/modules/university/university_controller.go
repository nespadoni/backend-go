package university

import (
	"backend-go/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	universityService *Service
}

func NewUniversityController(university *Service) *Controller {
	return &Controller{universityService: university}
}

func (c *Controller) FindAll(ctx *gin.Context) {
	universities, err := c.universityService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}

	ctx.JSON(http.StatusOK, universities)
}

func (c *Controller) FindById(ctx *gin.Context) {
	universityIDStr := ctx.Param("id")
	universityID, err := strconv.ParseUint(universityIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_university_id",
			Message: "ID da universidade deve ser um número válido",
		})
		return
	}

	response, err := c.universityService.FindById(uint(universityID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "university_not_found",
			Message: "Universidade não encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
