package main

import (
  "go-gin-api/handlers"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // Register the route to our handler function
  r.POST("/groceries", handlers.AddGroceryItem)

  r.Run() // default listens on :8080
}
