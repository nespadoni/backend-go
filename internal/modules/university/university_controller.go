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

// FindAll godoc
// @Summary Lista todas as universidades
// @Description Retorna uma lista de todas as universidades cadastradas
// @Tags universities
// @Accept json
// @Produce json
// @Success 200 {array} ListResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/universities [get]
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

// FindById godoc
// @Summary Busca universidade por Id
// @Description Retorna uma universidade específica pelo seu Id
// @Tags universities
// @Accept json
// @Produce json
// @Param id path string true "University Id"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/universities/{id} [get]
func (c *Controller) FindById(ctx *gin.Context) {
	universityIDStr := ctx.Param("id")
	universityID, err := strconv.ParseUint(universityIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_university_id",
			Message: "Id da universidade deve ser um número válido",
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

// Create godoc
// @Summary Cria uma nova universidade
// @Description Cria uma nova universidade com os dados fornecidos
// @Tags universities
// @Accept json
// @Produce json
// @Param university body CreateRequest true "University data"
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/universities [post]
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

// Update godoc
// @Summary Atualiza uma universidade
// @Description Atualiza os dados de uma universidade existente
// @Tags universities
// @Accept json
// @Produce json
// @Param id path string true "University Id"
// @Param university body UpdateRequest true "University data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/universities/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	universityIDStr := ctx.Param("id")
	universityID, err := strconv.ParseUint(universityIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_university_id",
			Message: "Id da universidade deve ser um número válido",
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

// Delete godoc
// @Summary Deleta uma universidade
// @Description Remove uma universidade do sistema
// @Tags universities
// @Accept json
// @Produce json
// @Param id path string true "University Id"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/universities/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	universityIDStr := ctx.Param("id")
	universityID, err := strconv.ParseUint(universityIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_university_id",
			Message: "Id da universidade deve ser um número válido",
		})
		return
	}

	if err := c.universityService.Delete(uint(universityID)); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "delete_university_failed",
			Message: "Falha ao deletar universidade",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
