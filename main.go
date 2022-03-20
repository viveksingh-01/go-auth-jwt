package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8000"
	}

	app := gin.New()
	app.Use(gin.Logger())

	app.GET("api/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello there!")
	})

	app.Run(":" + port)
}