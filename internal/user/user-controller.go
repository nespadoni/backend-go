package user

import (
	"fmt"
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

func (c UserController) FindAll(ctx *gin.Context) {
	users, err := c.userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, users)
}

func (c UserController) PostUser(ctx *gin.Context) {
	var newUser CreateUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := c.userService.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, response)
}

func (c UserController) UpdateUser(ctx *gin.Context) {
	idUser := ctx.Param("id")
	var updateRequest UpdateUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("User request:", updateRequest)
	user, err := c.userService.UpdateUser(idUser, updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, user)
	}

	ctx.JSON(http.StatusOK, user)
}

func (c UserController) DeleteUser(ctx *gin.Context) {

	userIdStr := ctx.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := c.userService.DeleteUser(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusNoContent, nil)
}
