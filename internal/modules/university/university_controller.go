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

func (c *Controller) Create(ctx *gin.Context) {
	var newUniversity CreateRequest
	if err := ctx.ShouldBindJSON(&newUniversity); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados da universidade inválidos",
		})
		return
	}

	response, err := c.universityService.Create(newUniversity)
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
	universityIDStr := ctx.Param("id")
	universityID, err := strconv.ParseUint(universityIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_university_id",
			Message: "ID da universidade deve ser um número válido",
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

	university, err := c.universityService.Update(uint(universityID), request)
	if err != nil {
		if err.Error() == "universidade não encontrada" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "university_not_found",
				Message: "Universidade não encontrada",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, university)
}

func (c *Controller) Delete(ctx *gin.Context) {
	universityIDStr := ctx.Param("id")
	universityID, err := strconv.ParseUint(universityIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_university_id",
			Message: "ID da universidade deve ser um número válido",
		})
		return
	}

	if err := c.universityService.Delete(uint(universityID)); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "university_not_found",
			Message: "Universidade não encontrada ou erro ao deletar",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
