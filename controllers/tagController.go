package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/services"
	"github.com/gin-gonic/gin"
)

func GetAllTags(c *gin.Context){
  tags, err := services.GetAllTags()
  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error getting tags"})
  }
  c.IndentedJSON(http.StatusOK, gin.H{"data": tags})
}
