package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/services"

	"github.com/gin-gonic/gin"
)

func GetPostByClass(c *gin.Context){
  classID := c.Param("class")
  posts, err := services.GetPostByClass(classID)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
  }

  c.IndentedJSON(http.StatusOK, gin.H{"data": posts})
}

func GetPostById(c *gin.Context){
  id := c.Param("id")
  post, err := services.GetPostById(id)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    return
  }

  c.IndentedJSON(http.StatusOK, gin.H{"data": post})
}

func GetMaterialById(c *gin.Context){
  id := c.Param("id")
  material, err := services.GetMaterialById(id)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    return
  }

  c.IndentedJSON(http.StatusOK, gin.H{"data": material})
}

func CreatePost(c *gin.Context){
  var postBody models.Post
  c.BindJSON(&postBody)

  post, err := services.CreatePost(postBody)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
  }

  c.IndentedJSON(http.StatusCreated, gin.H{"message": post})
}
