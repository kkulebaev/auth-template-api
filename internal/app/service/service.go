package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	id       int
	username string
	password string
}

var users = []User{
	{
		id:       1,
		username: "Lorem",
		password: "Ipsum",
	},
	{
		id:       2,
		username: "Lorem2",
		password: "Ipsum2",
	},
	{
		id:       3,
		username: "Lorem3",
		password: "Ipsum3",
	},
	{
		id:       4,
		username: "Lorem4",
		password: "Ipsum4",
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
