package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{
		ID:       1,
		Username: "Lorem",
		Password: "Ipsum",
	},
	{
		ID:       2,
		Username: "Lorem2",
		Password: "Ipsum2",
	},
	{
		ID:       3,
		Username: "Lorem3",
		Password: "Ipsum3",
	},
	{
		ID:       4,
		Username: "Lorem4",
		Password: "Ipsum4",
	},
}

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "All is OK")
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func LogIn(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "LogIn")
}
