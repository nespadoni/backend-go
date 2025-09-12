package user

import (
	"backend-go/pkg/utils" // Adicionar este import
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *Service
}

func NewUserController(userService *Service) *Controller {
	return &Controller{userService: userService}
}

// UploadProfilePhoto godoc
// @Summary Upload de foto de perfil
// @Description Faz upload da foto de perfil do usuário
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param profilePhoto formData file true "Foto de perfil"
// @Success 200 {object} map[string]string
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/users/profile-photo [post]
func (c *Controller) UploadProfilePhoto(ctx *gin.Context) {
	// Extrair ID do usuário do contexto (middleware JWT)
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Error:   "unauthorized",
			Message: "Usuário não autenticado",
		})
		return
	}

	userIDFloat, ok := userIDInterface.(float64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Error:   "unauthorized",
			Message: "Formato de ID de usuário inválido no token.",
		})
		return
	}
	userID := uint(userIDFloat)

	// Obter o arquivo do formulário
	file, header, err := ctx.Request.FormFile("profilePhoto")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "file_not_found",
			Message: "Arquivo não encontrado",
		})
		return
	}
	defer file.Close()

	// Validar tipo de arquivo
	if !isValidImageType(header) {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_file_type",
			Message: "Tipo de arquivo inválido. Use JPG, PNG ou GIF.",
		})
		return
	}

	// Validar tamanho do arquivo (5MB max)
	if header.Size > 5*1024*1024 {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "file_too_large",
			Message: "Arquivo muito grande. Máximo: 5MB.",
		})
		return
	}

	// Salvar arquivo
	photoURL, err := c.saveProfilePhoto(file, header, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "upload_failed",
			Message: "Erro ao fazer upload da foto.",
		})
		return
	}

	// Atualizar usuário no banco
	err = c.userService.UpdateProfilePhoto(userID, photoURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Error:   "update_failed",
			Message: "Erro ao atualizar foto no perfil.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"profilePhotoUrl": photoURL,
		"message":         "Foto de perfil atualizada com sucesso",
	})
}

// Funções auxiliares
func isValidImageType(header *multipart.FileHeader) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
	}

	contentType := header.Header.Get("Content-Type")
	return validTypes[contentType]
}

func (c *Controller) saveProfilePhoto(file multipart.File, header *multipart.FileHeader, userID uint) (string, error) {
	// Criar diretório se não existir
	uploadDir := "uploads/profile_photos"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	// Gerar nome único para o arquivo
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("user_%d_%d%s", userID, time.Now().Unix(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	// Criar arquivo no servidor
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copiar conteúdo
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	// Formata a URL para o padrão web
	webPath := fmt.Sprintf("/uploads/profile_photos/%s", fileName)
	// Garante que a URL use sempre a barra correta ('/'), independentemente do sistema operacional
	urlPath := strings.ReplaceAll(webPath, "\\", "/")

	return urlPath, nil
}

// FindAll godoc
// @Summary Lista todos os usuários
// @Description Retorna uma lista de todos os usuários
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} Response
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/users [get]
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
// @Summary Busca usuário por Id
// @Description Retorna um usuário específico pelo seu Id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/users/{id} [get]
func (c *Controller) FindById(ctx *gin.Context) {
	userIDStr := ctx.Param("id")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "Id do usuário deve ser um número válido",
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
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/users [post]
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
// @Param id path string true "User Id"
// @Param user body UpdateUserRequest true "User data"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/users/{id} [put]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	idUserStr := ctx.Param("id")

	userId, err := strconv.ParseUint(idUserStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "Id do usuário deve ser um número válido",
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
// @Param id path string true "User Id"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/users/{id} [delete]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	userIdStr := ctx.Param("id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error:   "invalid_user_id",
			Message: "Id do usuário deve ser um número válido",
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
