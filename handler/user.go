package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct{
	userService user.Service
	authServicce auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	token, err := h.authServicce.GenerateToken(newUser.Id)
	if err != nil {
		res := helper.APIResponse("Register Account Failed",http.StatusBadRequest,"Error",nil)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatUser(newUser, token)

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

	// token
	token, err := h.authServicce.GenerateToken(loggedInUser.Id)
	if err != nil {
		res := helper.APIResponse("Login Failed",http.StatusBadRequest,"Error",nil)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)


	res := helper.APIResponse("Login Success", http.StatusOK,"success", formatter)

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) CheckEmailAvailibity(c *gin.Context)  {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormat(err)

		errorMess := gin.H{"error":errors}

		res := helper.APIResponse("Checking Email Failes",http.StatusUnprocessableEntity,"Error",errorMess)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvail(input)
	if err != nil {
		errorMess := gin.H{"error":"Server Error"}

		res := helper.APIResponse("Checking Email Failes",http.StatusUnprocessableEntity,"Error",errorMess)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	data := gin.H{
		"is_available" : isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	res := helper.APIResponse(metaMessage,http.StatusOK,"success",data)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) UploadAvatar(c *gin.Context)  {
	 file, err := c.FormFile("avatar")

	 if err != nil {
		 data := gin.H{"is_uploaded":false}
		 res := helper.APIResponse("Failed To Upload Avatar", http.StatusBadRequest, "error", data)
		 c.JSON(http.StatusBadRequest,res)
		 return
	 }

	 userId := 1

	//  path := "images/" + file.Filename

	 path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	 err =c.SaveUploadedFile(file, path)
	 if err != nil {
		data := gin.H{"is_uploaded":false}
		res := helper.APIResponse("Failed To Upload Avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest,res)
		return
	}

	

	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uploaded":false}
		res := helper.APIResponse("Failed To Upload Avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest,res)
		return
	}

	data := gin.H{"is_uploaded":true}
		res := helper.APIResponse("user Avatar Uploaded", http.StatusOK, "success", data)
		c.JSON(http.StatusOK,res)



}