package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"try-gin.com/controllers"
	"try-gin.com/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/books", controllers.FindBooks)

	r.Run()
}