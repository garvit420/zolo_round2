package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Book represents a book in the library.
type Book struct {
	ID       string    `json:"book_id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	PostedBy string    `json:"posted_by"`
	TillDate time.Time `json:"till_date"`
}

// Borrow represents a record of a book being borrowed.
type Borrow struct {
	ID              string      `json:"borrow_id" gorm:"primary_key"`
	BookID          string      `json:"book_id"`
	BorrowStartTime time.Time   `json:"borrow_start_time"`
	BorrowEndTime   time.Time   `json:"borrow_end_time"`
	Returned        bool        `json:"returned"`
}

var db *gorm.DB

func main() {
	var err error
    // Directly specifying the database connection details
    db, err = gorm.Open("mysql", "avnadmin:AVNS_9tXBtUpZnwTq-vievbj@tcp(mysql-zolo-zolo.a.aivencloud.com:12993)/defaultdb?parseTime=true")
    if err != nil {
        panic("failed to connect to database")
    }
	defer db.Close()

	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())

	// Book routes
	router.GET("/api/v1/booky", getBooks)
	router.POST("/api/v1/booky", postBooks)

	// Borrow routes
	router.POST("/api/v1/booky/:book_id/borrow", postBorrow)
	router.GET("/api/v1/booky/borrows", getBorrows)
	router.PUT("/api/v1/booky/:book_id/borrow/:borrow_id", putReturnBorrow)

	router.Run("localhost:8080")
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
