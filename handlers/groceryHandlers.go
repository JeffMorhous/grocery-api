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

// Fetch a specific grocery item by ID
func GetGroceryItem(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  for _, item := range data.GroceryDB {
    if item.ID == id {
      c.JSON(http.StatusOK, item)
      return
    }
  }

  c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// List all grocery items
func ListGroceryItems(c *gin.Context) {
  c.JSON(http.StatusOK, data.GroceryDB)
}

