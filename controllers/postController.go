package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/services"

	"github.com/gin-gonic/gin"
)

func GetPostByClass(c *gin.Context){

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
