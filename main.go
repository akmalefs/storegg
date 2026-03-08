package main

import (
	"os"
	"storegg-backend/auth"
	"storegg-backend/config"
	"storegg-backend/handler"
	"storegg-backend/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.InitDB()

	secret := os.Getenv("JWT_SECRET")
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService(secret)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	users := api.Group("/users")
	{
		users.POST("/register", userHandler.RegisterUser)
		users.POST("/login", userHandler.LoginUser)
		users.POST("/email_checkers", userHandler.IsEmailAvailable)
	}

	router.Run()
}
