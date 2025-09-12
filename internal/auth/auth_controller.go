package auth

import (
	"backend-go/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{service: service}
}

// Register godoc
// @Summary Registra um novo usuário com foto de perfil
// @Description Cria uma nova conta de usuário, incluindo a foto de perfil, a partir de um formulário multipart.
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Nome completo do usuário"
// @Param email formData string true "Email do usuário"
// @Param password formData string true "Senha do usuário"
// @Param telephone formData string true "Telefone do usuário"
// @Param universityId formData string false "ID da universidade (opcional, 'none' se não aplicável)"
// @Param profilePhoto formData file false "Foto de perfil do usuário (opcional)"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	// Popula o request a partir do formulário multipart
	universityIDValue := ctx.PostForm("universityId")
	req := RegisterRequest{
		Name:      ctx.PostForm("name"),
		Email:     ctx.PostForm("email"),
		Password:  ctx.PostForm("password"),
		Telephone: ctx.PostForm("telephone"),
	}
	if universityIDValue != "" {
		req.UniversityID = &universityIDValue
	}

	// Obtém o arquivo da foto
	file, header, err := ctx.Request.FormFile("profilePhoto")
	if err != nil && err != http.ErrMissingFile {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid_file", Message: "Erro ao processar arquivo."})
		return
	}

	// Garante que o arquivo seja fechado se existir
	if file != nil {
		defer file.Close()
	}

	// Chama o serviço para registrar o usuário, passando o arquivo
	result, err := c.service.Register(req, file, header)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "registration_failed", Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// Login godoc
// @Summary Realiza o login do usuário
// @Description Autentica o usuário e retorna um token JWT com os dados do usuário
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Credenciais de Login"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid_request", Message: err.Error()})
		return
	}

	// A função Login do serviço agora retorna um map
	result, err := c.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{Error: "authentication_failed", Message: err.Error()})
		return
	}

	// Retorna o resultado completo (token + user)
	ctx.JSON(http.StatusOK, result)
}
