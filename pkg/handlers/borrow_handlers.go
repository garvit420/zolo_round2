package handlers

import (
	"net/http"
    "github.com/garvit420/zolo_round2/pkg/models"

    "github.com/gin-gonic/gin"
)

func RegisterBorrowRoutes(router *gin.Engine) {
    router.POST("/api/v1/booky/:book_id/borrow", postBorrow)
    router.GET("/api/v1/booky/borrows", getBorrows)
    router.PUT("/api/v1/booky/:book_id/borrow/:borrow_id", putReturnBorrow)
}

// postBorrow handles the creation of a borrow record.
func postBorrow(c *gin.Context) {
	var newBorrow Borrow
	bookID := c.Param("book_id")

	if err := c.BindJSON(&newBorrow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrow data"})
		return
	}

	var book Book
	if err := db.First(&book, "id = ?", bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if newBorrow.BorrowEndTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrow end time"})
		return
	}

	newBorrow.BookID = bookID
	if err := db.Create(&newBorrow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to borrow book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Borrow added successfully"})
}

// getBorrows handles the retrieval of all borrows.
func getBorrows(c *gin.Context) {
	var borrows []Borrow
	if err := db.Find(&borrows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching borrow records"})
		return
	}
	c.JSON(http.StatusOK, borrows)
}

// putReturnBorrow handles the return of a borrowed book.
func putReturnBorrow(c *gin.Context) {
	bookID := c.Param("book_id")
	borrowID := c.Param("borrow_id")

	var borrow Borrow
	if err := db.Where("id = ? AND book_id = ?", borrowID, bookID).First(&borrow).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrow record not found"})
		return
	}

	if borrow.Returned {
		c.JSON(http.StatusConflict, gin.H{"error": "Book already returned"})
		return
	}

	borrow.Returned = true
	if err := db.Save(&borrow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update borrow record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}