package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct{
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.ErrorValidationFormat(err)

		errorMess := gin.H{"error":errors}

		res := helper.APIResponse("Register Account Failed",http.StatusUnprocessableEntity,"Error",errorMess)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		res := helper.APIResponse("Register Account Failed",http.StatusBadRequest,"Error",nil)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// token



	formatter := user.FormatUser(newUser, "token cuyyy")

	res := helper.APIResponse("Your Account Has Been Created", http.StatusOK,"success", formatter)

	c.JSON(http.StatusOK, res)

}

func (h *userHandler) Login(c *gin.Context)  {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormat(err)

		errorMess := gin.H{"error":errors}

		res := helper.APIResponse("Login Failed",http.StatusUnprocessableEntity,"Error",errorMess)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	loggedInUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMess := gin.H{"error":err.Error()}

		res := helper.APIResponse("Login Failed",http.StatusUnprocessableEntity,"Error",errorMess)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	formatter := user.FormatUser(loggedInUser, "token cuyyy")


	res := helper.APIResponse("Login Success", http.StatusOK,"success", formatter)

	c.JSON(http.StatusOK, res)
}