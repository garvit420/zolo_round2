package main

import (
    "github.com/garvit420/zolo_round2/pkg/database"
    "github.com/garvit420/zolo_round2/pkg/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    db := database.InitDB()
    defer db.Close()

    router := gin.Default()
    router.Use(gin.Recovery(), gin.Logger())

    handlers.RegisterBookRoutes(router)
    handlers.RegisterBorrowRoutes(router)

    router.Run("localhost:8080")
}