package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateClass(c *gin.Context) {
  var class models.Class

  c.BindJSON(&class)

  class.ID = uuid.NewString()

  teachersArray := pq.Array(class.Teachers)
  studentsArray := pq.Array(class.Students)
  postsArray := pq.Array(class.Posts)

  _, err := database.DB.Exec("INSERT INTO classes (id, name, teachers, students, posts) values ($1, $2, $3, $4, $5)", class.ID, class.Name, teachersArray, studentsArray, postsArray)
  if err != nil {
    println(err.Error(), class.Teachers)
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
    return
  }
  c.IndentedJSON(http.StatusCreated, class)
}

func GetAllClasses(c *gin.Context) {
  rows, err := database.DB.Query("SELECT id, name FROM classes")

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting classes"})
    return 
  }

  defer rows.Close()

  var classes []models.Class

  for rows.Next() {
    var class models.Class
    if err := rows.Scan(&class.ID, &class.Name); err != nil{
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error scanning classes"})
      return
    }
    classes = append(classes, class)
  }

  if err := rows.Err(); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error iterating classes"})
    return
  }

  c.IndentedJSON(http.StatusOK, classes)
}
