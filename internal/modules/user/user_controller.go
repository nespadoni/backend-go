package user

import (
	"backend-go/pkg/utils" // Adicionar este import
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *Service
}

func NewUserController(userService *Service) *Controller {
	return &Controller{userService: userService}
}

// Remover esta definição local - usar utils.ErrorResponse

// FindAll godoc
// @Summary Lista todos os usuários
// @Description Retorna uma lista de todos os usuários
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} UserResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/user [get]
func (c *Controller) FindAll(ctx *gin.Context) {
	users, err := c.userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "internal_server_error",
			Message: "Erro interno do servidor",
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// FindById godoc
// @Summary Busca usuário por ID
// @Description Retorna um usuário específico pelo seu ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/user/{id} [get]
func (c *Controller) FindById(ctx *gin.Context) {
	userIDStr := ctx.Param("id")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "ID do usuário deve ser um número válido",
		})
		return
	}

	response, err := c.userService.GetById(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "user_not_found",
			Message: "Usuário não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// PostUser godoc
// @Summary Cria um novo usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data"
// @Success 201 {object} UserResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/user [post]
func (c *Controller) PostUser(ctx *gin.Context) {
	var newUser CreateUserRequest
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados do usuário inválidos",
		})
		return
	}

	response, err := c.userService.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "creation_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// UpdateUser godoc
// @Summary Atualiza um usuário
// @Description Atualiza os dados de um usuário existente
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "User data"
// @Success 200 {object} UserResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/user/{id} [put]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	idUserStr := ctx.Param("id")

	userId, err := strconv.ParseUint(idUserStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "ID do usuário deve ser um número válido",
		})
		return
	}

	var updateRequest UpdateUserRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_request_body",
			Message: "Dados de atualização inválidos",
		})
		return
	}

	user, err := c.userService.UpdateUser(uint(userId), updateRequest)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Error:   "user_not_found",
				Message: "Usuário não encontrado",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Deleta um usuário
// @Description Remove um usuário do sistema
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/user/{id} [delete]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	userIdStr := ctx.Param("id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "ID do usuário deve ser um número válido",
		})
		return
	}

	if err := c.userService.DeleteUser(uint(userId)); err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Error:   "user_not_found",
			Message: "Usuário não encontrado",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
