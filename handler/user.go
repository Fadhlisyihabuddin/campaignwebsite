package handler

import (
	"campaignwebsite/auth"
	"campaignwebsite/helper"
	"campaignwebsite/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(" Register account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newuser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse(" Register account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newuser.ID)

	if err != nil {
		response := helper.ApiResponse(" Register account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newuser, token)

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(" Login account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggeduser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse(" Login account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggeduser.ID)

	if err != nil {
		response := helper.ApiResponse(" Login account Failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggeduser, token)

	response := helper.ApiResponse("Login Succes", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailable(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"error": "Server Error"}

		response := helper.ApiResponse("Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email Has Been Registered"

	if isEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.ApiResponse(metaMessage, http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("Upload Failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("Upload Failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("Upload Failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.ApiResponse("Upload Succes", http.StatusOK, "succes", data)
	c.JSON(http.StatusOK, response)
}
