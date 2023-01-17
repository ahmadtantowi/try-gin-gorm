package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"try-gin.com/models"
)

func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}
