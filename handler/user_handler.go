package handler

import (
	"net/http"
	"storegg-backend/auth"
	"storegg-backend/helper"
	"storegg-backend/user"

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

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		newResponse := helper.APIResponse("Regist account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, newResponse)
		return
	}

	userFormat := user.NewUserResponse(newUser, token)

	newResponse := helper.APIResponse("Registered account success", http.StatusOK, "success", userFormat)
	c.JSON(http.StatusOK, newResponse)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	input := user.LoginUserInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}

		newResponse := helper.APIResponse("Login is failed", http.StatusUnprocessableEntity, "erros", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, newResponse)

		return
	}

	loginUser, err := h.userService.LoginUser(input)
	if err != nil {

		errorMessage := gin.H{
			"errors": err.Error(),
		}

		newResponse := helper.APIResponse("Login is failed", http.StatusUnprocessableEntity, "erros", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, newResponse)

		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		newResponse := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, newResponse)
		return
	}

	userFormat := user.NewUserResponse(loginUser, token)

	newResponse := helper.APIResponse("Login success", http.StatusOK, "success", userFormat)
	c.JSON(http.StatusOK, newResponse)
}

func (h *userHandler) IsEmailAvailable(c *gin.Context) {
	input := user.CheckEmailInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}

		newResponse := helper.APIResponse("Login is failed", http.StatusUnprocessableEntity, "erros", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, newResponse)

		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {

		errorMessage := gin.H{
			"errors": err.Error(),
		}

		newResponse := helper.APIResponse("Login is failed", http.StatusUnprocessableEntity, "erros", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, newResponse)

		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	newResponse := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, newResponse)
}
