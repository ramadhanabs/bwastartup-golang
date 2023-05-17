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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorResponse(err)

		response := helper.APIResponse("Registration Failed", http.StatusBadRequest, "error", nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errors := helper.ErrorResponse(err)
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedUser := user.FormatUser(loggedInUser, "tokentoken")
	response := helper.APIResponse("Success Login", http.StatusOK, "success", formattedUser, nil)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorResponse(err)

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", nil, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errors := helper.ErrorResponse(err)

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", nil, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	metaMessage := "Email already registered"
	if IsEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data, nil)
	c.JSON(http.StatusOK, response)
}
