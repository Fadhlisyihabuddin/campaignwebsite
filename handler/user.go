package handler

import (
	"campaignwebsite/helper"
	"campaignwebsite/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
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

	formatter := user.FormatUser(newuser, "tokentokentoken")

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)

}
