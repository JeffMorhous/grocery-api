package main

import (
  "go-gin-api/handlers"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // Register the route to our handler function
  r.POST("/groceries", handlers.AddGroceryItem)

  // New GET routes
  r.GET("/groceries", handlers.ListGroceryItems)
  r.GET("/groceries/:id", handlers.GetGroceryItem)

  r.Run() // default listens on :8080
}
