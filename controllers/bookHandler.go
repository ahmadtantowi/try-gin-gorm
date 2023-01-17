package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"try-gin.com/models"
)

// list book
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// add book
func CreateBook(ctx *gin.Context) {
	// validate input
	var input CreateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	ctx.JSON(http.StatusCreated, gin.H{"data": book})
}
