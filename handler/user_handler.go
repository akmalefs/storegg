package handler

import (
	"net/http"
	"storegg-backend/helper"
	"storegg-backend/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	input := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}

		newResponse := helper.APIResponse("Registered account failed", http.StatusUnprocessableEntity, "erros", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, newResponse)

		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		newResponse := helper.APIResponse("Registered account failed", http.StatusUnprocessableEntity, "erros", nil)

		c.JSON(http.StatusUnprocessableEntity, newResponse)

		return
	}

	userFormat := user.NewUserResponse(newUser, "qwertylgfdsasfg")

	newResponse := helper.APIResponse("Registered account success", http.StatusOK, "success", userFormat)
	c.JSON(http.StatusOK, newResponse)
}
