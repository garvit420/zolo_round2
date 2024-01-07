package handlers

import (
    "net/http"
    "github.com/garvit420/zolo_round2/pkg/models"

    "github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine) {
    router.GET("/api/v1/booky", getBooks)
    router.POST("/api/v1/booky", postBooks)
}

// getBooks handles the retrieval of all books.
func getBooks(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// postBooks handles the addition of a new book.
func postBooks(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}

	if err := db.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}