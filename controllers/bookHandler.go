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

// get a book
func FindBook(ctx *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// update a book
func UpdateBook(ctx *gin.Context) {
	// find a book first
	var book models.Book
	if err := models.DB.Where("id=?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Model(&book).Updates(&updateBook)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// delete a book
func DeleteBook(ctx *gin.Context) {
	// find a book first
	var book models.Book
	if err := models.DB.Where("id=?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
