package models

import "time"

// Book represents a book in the library.
type Book struct {
    ID       string    `json:"book_id" gorm:"primary_key"`
    Name     string    `json:"name"`
    Author   string    `json:"author"`
    PostedBy string    `json:"posted_by"`
    TillDate time.Time `json:"till_date"`
}