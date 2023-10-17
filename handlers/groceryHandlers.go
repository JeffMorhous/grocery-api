package handlers

import (
  "net/http"
  "go-gin-api/data"
  "go-gin-api/models"

  "github.com/gin-gonic/gin"
)

func AddGroceryItem(c *gin.Context) {
  var newItem models.GroceryItem

  if err := c.ShouldBindJSON(&newItem); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Assign a new ID and append to the database
  newItem.ID = data.IDCounter
  data.GroceryDB = append(data.GroceryDB, newItem)
  data.IDCounter++

  c.JSON(http.StatusCreated, newItem)
}
