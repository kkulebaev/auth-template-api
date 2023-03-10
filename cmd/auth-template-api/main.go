package main

import (
	"log"
	"auth-template-api/internal/app/service"

	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(cors.Default())
	router.POST("/login", service.LogIn)
	router.GET("/ping", service.Ping)
	router.GET("/users", service.GetUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}