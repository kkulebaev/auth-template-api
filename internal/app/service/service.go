package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "All is OK")
}

func LogIn(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "LogIn")
}