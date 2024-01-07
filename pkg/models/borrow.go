package models

import "time"

// Borrow represents a record of a book being borrowed.
type Borrow struct {
    ID              string    `json:"borrow_id" gorm:"primary_key"`
    BookID          string    `json:"book_id"`
    BorrowStartTime time.Time `json:"borrow_start_time"`
    BorrowEndTime   time.Time `json:"borrow_end_time"`
    Returned        bool      `json:"returned"`
}