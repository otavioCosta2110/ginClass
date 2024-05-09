package controllers

import (
  "net/http"
  "otaviocosta2110/ginClass/database"
  "otaviocosta2110/ginClass/models"
  "otaviocosta2110/ginClass/repositories"

  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
)

func GetPostByClass(c *gin.Context){

}

func CreatePost(c *gin.Context){
  var post models.Post

  c.BindJSON(&post)

  if  len(post.Teachers) < 0 || post.Content == "" || post.Name == "" {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Missing Fields"})
    return
  }
  post.ID =uuid.NewString()

  tx, err := database.DB.Begin()

  if err != nil {
    tx.Rollback()
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
    return
  }

  defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Panic occurred. Transaction rolled back."})
    }
  }()
  _, err = database.DB.Query("INSERT INTO posts (id, name, class_id, content) values ($1, $2, $3, $4)", post.ID, post.Name, post.ClassID, post.Content)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
    panic(err)
  }

  for _, tagContent := range post.Tags {
    tagID, err := repositories.CreateTags(tagContent, tx)

    if err != nil {
      tx.Rollback()
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating tag"})
      return
    }
    _, err = tx.Exec("INSERT INTO post_tag (post_id, tag_id) values ($1, $2)", post.ID, tagID)

    if err != nil {
      tx.Rollback()
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating post_tag"})
      return
    }

  }

  for _, userEmail := range post.Teachers {
    user, err := repositories.GetUserByEmail(userEmail)
    _, err = database.DB.Query("INSERT INTO user_post (user_id, post_id) values ($1, $2)", user.ID, post.ID)

    if err != nil {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
      panic(err)
    }
  }

  for _, material := range post.Material {
    materialID := uuid.NewString()
    _, err = database.DB.Query("INSERT INTO materials (id, content) values ($1, $2)", materialID, material)
    println(material)

    if err != nil {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
      panic(err)
    }

    _, err = database.DB.Query("INSERT INTO post_material (post_id, material_id) values ($1, $2)", post.ID, materialID)

    if err != nil {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
      panic(err)
    }
  }

  err = tx.Commit()
  if err != nil {
    tx.Rollback()
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error committing transaction"})
    return
  }

  c.IndentedJSON(http.StatusCreated, post)
}
