package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func GetPostByClass(c *gin.Context){

}

func CreatePost(c *gin.Context){
  var post models.Post

  c.BindJSON(&post)

  post.ID =uuid.NewString()

  postTags := pq.Array(post.Tags)
  postTeachers := pq.Array(post.Teachers)

  _, err := database.DB.Query("INSERT INTO posts (id, name, class_id, content, material, tags, teachers) values ($1, $2, $3, $4, $5, $6, $7)", post.ID, post.Name, post.ClassID, post.Content, post.Material, postTags, postTeachers)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
    panic(err)
  }

  c.IndentedJSON(http.StatusCreated, post)
}
