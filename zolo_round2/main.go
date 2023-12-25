package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Book represents data about a book.
type Book struct {
	ID        string    `json:"book_id"`
	Name      string    `json:"Name"`
	Author    string    `json:"Author"`
	PostedBy  string    `json:"posted_by"`
	TillDate  time.Time `json:"till_date"`
	Genre     []string  `json:"genre"`
}

// Borrow represents data about borrowing a book.
type Borrow struct {
	ID              string    `json:"borrow_id"`
	BookID          string    `json:"book_id"`
	BorrowStartTime time.Time `json:"borrow_Start_Time"`
	BorrowEndTime   time.Time `json:"borrow_End_Time"`
	Returned        bool      `json:"returned"`
}

var (
	books   = []Book{
		{ID: "1", Name: "Blue Train", Author: "John Coltrane", PostedBy: "garvit", TillDate: time.Date(2023, time.December, 27, 23, 59, 0, 0, time.UTC), Genre: []string{"Fiction", "Thriller"}},
		{ID: "2", Name: "Jeru", Author: "Gerry Mulligan", PostedBy: "garvit", TillDate: time.Date(2023, time.December, 30, 23, 59, 0, 0, time.UTC), Genre: []string{"Fiction", "Thriller"}},
		{ID: "3", Name: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan", PostedBy: "garvit", TillDate: time.Date(2023, time.December, 31, 23, 59, 0, 0, time.UTC), Genre: []string{"Fiction", "Thriller"}},
	}
	borrows = []Borrow{
		{ID: "1", BookID: "1", BorrowStartTime: time.Now(), BorrowEndTime: time.Date(2023, time.December, 27, 23, 59, 0, 0, time.UTC), Returned: false},
	}
)

func main() {
	router := gin.Default()

	// Define API endpoints

	// GET request to fetch all books
	router.GET("/api/v1/booky", getBooks)

	// POST request to add a new book
	router.POST("/api/v1/booky", postBooks)

	// POST request to borrow a book
	router.POST("/api/v1/booky/:book_id/borrow", postBorrow)

	// PUT request to return a borrowed book
	router.PUT("/api/v1/booky/:book_id/borrow/:borrow_id", putReturnBorrow)

	// Start the server
	router.Run("localhost:8080")
}

// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds a book from JSON received in the request body.
func postBooks(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new album to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


func postBorrow(c *gin.Context) {
    bookID := c.Param("book_id")
    var newBorrow Borrow

    if err := c.BindJSON(&newBorrow); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrow data"})
        return
    }

    bookFound := false
    for _, b := range books {
        if b.ID == bookID {
            bookFound = true

            // Check if the book is available for borrowing based on current date and till date
            currentDate := time.Now().UTC()
            if currentDate.After(b.TillDate) {
                c.JSON(http.StatusConflict, gin.H{"error": "Book's till date is passed"})
                return
            }

            // Check if the book is already borrowed
            borrowed := false
            for _, borrow := range borrows {
                if borrow.BookID == bookID && !borrow.Returned {
                    borrowed = true
                    break
                }
            }

            if borrowed {
                c.JSON(http.StatusConflict, gin.H{"error": "Book already borrowed"})
                return
            }

            // Check if the provided dates are within valid bounds
            if newBorrow.BorrowEndTime.After(b.TillDate){
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrow dates"})
                return
            }

            // Add the borrow
            borrows = append(borrows, Borrow{
                ID:              newBorrow.ID,
                BookID:          bookID,
                BorrowStartTime: newBorrow.BorrowStartTime,
                BorrowEndTime:   newBorrow.BorrowEndTime,
                Returned:        false,
            })
            c.JSON(http.StatusCreated, gin.H{"message": "Borrow added successfully"})
            return
        }
    }

    if !bookFound {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
    }
}


func putReturnBorrow(c *gin.Context) {
	bookID := c.Param("book_id")
	borrowID := c.Param("borrow_id")

	// Check if the book exists
	bookExists := false
	for _, b := range books {
		if b.ID == bookID {
			bookExists = true
			break
		}
	}
	if !bookExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check if the borrow ID corresponds to the given book ID and mark as returned if not already returned
	borrowFound := false
	for i := range borrows {
		if borrows[i].BookID == bookID && borrows[i].ID == borrowID {
			if borrows[i].Returned {
				c.JSON(http.StatusConflict, gin.H{"error": "Borrow already marked as returned"})
				return
			}
			borrows[i].Returned = true
			c.JSON(http.StatusOK, gin.H{"message": "Borrow returned successfully"})
			borrowFound = true
			return
		}
	}

	if !borrowFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrow ID not found for the given book"})
		return
	}
}



