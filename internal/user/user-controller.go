package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService: userService}
}

func (controller UserController) CreateUser(ctx *gin.Context) {

}

func (controller UserController) FindById(ctx *gin.Context) {
	userIDstr := ctx.Param("id")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	reponse, erro := controller.userService.GetById(userID)
	if erro != nil {
		ctx.JSON(http.StatusBadRequest, erro.Error())
	}

	ctx.JSON(http.StatusOK, reponse)
}
