package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func panicRoute(c *gin.Context) {
	panic("foo!")
}

func Run() {
	router := gin.Default()

	router.GET("/ping", pingRoute)
	router.GET("/panic", panicRoute)

	router.Run()
}
