package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open("mysql", "your_connection_string")
    if err != nil {
        panic("failed to connect to database")
    }
    // db.AutoMigrate(&models.Book{}, &models.Borrow{})
    return db
}