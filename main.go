package main

import (
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
	"project.com/gweb/gweb"
)

func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "sumbit")
}

func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func main() {
	fmt.Print("hello!")
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	// {
	v1.POST("/login", loginEndpoint)
	v1.POST("/submit", submitEndpoint)
	v1.POST("/read", readEndpoint)
	// }

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	fmt.Print("hello!")
	fmt.Print(gweb.DATABASES["type"])
	fmt.Print("hello!")
	router.Run(":8080")
}
