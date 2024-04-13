package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetClassByTeacher(c *gin.Context) {
  teacherEmail := c.Param("teacheremail")
  println(teacherEmail)

  classes, err := repositories.ClassByTeacher(teacherEmail)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting class by teacher email"})
    panic(err)
  }

  c.IndentedJSON(http.StatusOK, &classes)
}

func AddTeacher(c *gin.Context) {
  var body models.AddUser

  c.BindJSON(&body)

  teacher, err := repositories.UserByEmail(body.TeacherEmail)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting user!"})
    panic(err)
  }

  if !teacher.IsTeacher {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User is not a teacher!"})
    return
  }

  repositories.AddUser(body.TeacherEmail, body.ClassID)

}

func CreateClass(c *gin.Context) {
  var class models.Class

  c.BindJSON(&class)

  class.ID = uuid.NewString()

  tx, err := database.DB.Begin()
  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
    return
  }

  defer func(){
    if err != nil {
      tx.Rollback()
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
      return
    }
    tx.Commit()
  }()

  _, err = database.DB.Exec("INSERT INTO classes (id, name) values ($1, $2)", class.ID, class.Name)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
    return
  }

  for _, teacher := range class.Teachers {
    teacherID, err := repositories.UserByEmail(teacher)

    if teacherID == nil {
      c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Teacher with email " + teacher + " does not exist"})
      return
    }
    if err != nil {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting teacher ID"})
      return
    }

    _, err = tx.Exec("INSERT INTO user_class (user_id, class_id) values ($1, $2)", teacherID.ID, class.ID)

    if err != nil {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
      return
    }
  }

  c.IndentedJSON(http.StatusCreated, class.Tags)
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
