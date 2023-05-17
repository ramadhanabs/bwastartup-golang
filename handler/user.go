package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
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
		errors := helper.ErrorResponse(err)

		response := helper.APIResponse("Registration Failed", http.StatusBadRequest, "error", nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registration Failed", http.StatusBadRequest, "error", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedUser := user.FormatUser(newUser, "tokenasal")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formattedUser, nil)

	c.JSON(http.StatusOK, response)
}
