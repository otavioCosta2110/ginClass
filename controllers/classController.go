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

  teacher, err := repositories.GetUserByEmail(body.TeacherEmail)

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

    if len(class.Teachers) < 1 || class.Name == "" {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Missing Fields"})
      return
    }

    class.ID = uuid.NewString()

    users := append(class.Teachers, class.Students...)

    tx, err := database.DB.Begin()
    if err != nil {
      tx.Rollback()
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
      return
    }

    defer func() {
      if r := recover(); r != nil {
        tx.Rollback()
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Panic occurred. Transaction rolled back."})
      }
    }()

    _, err = tx.Exec("INSERT INTO classes (id, name) values ($1, $2)", class.ID, class.Name)
    if err != nil {
      tx.Rollback()
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
      return
    }

    for _, userEmail := range  users {
        user, err := repositories.GetUserByEmail(userEmail)
        if user == nil {
            tx.Rollback()
            c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User with email " + userEmail + " does not exist"})
            return
        }
        if err != nil {
            tx.Rollback()
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting teacher ID"})
            return
        }
        _, err = tx.Exec("INSERT INTO user_class (user_id, class_id) values ($1, $2)", user.ID, class.ID)
        if err != nil {
            tx.Rollback()
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class"})
            return
        }
    }
    
    for _, tagContent := range class.Tags {
      tagID, err := repositories.CreateTags(tagContent, tx, c)
      if err != nil {
        tx.Rollback()
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating tag"})
        return
      }

      _, err = tx.Exec("INSERT INTO class_tag (class_id, tag_id) values ($1, $2)", class.ID, tagID)
      if err != nil {
        tx.Rollback()
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating class_tag"})
        return
      }
    }

    err = tx.Commit()
    if err != nil {
        tx.Rollback()
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error committing transaction"})
        return
    }

    c.IndentedJSON(http.StatusCreated, class)
}

func DeleteClass(c *gin.Context) {
  classID := c.Param("id")

  isClassDeleted, err := repositories.IsClassDeleted(classID)
  println("IsClassDeleted: ", isClassDeleted)
  if err != nil {
    println("error: ", err.Error())
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting class"})
    return
  }

  println("IsClassDeleted: ", isClassDeleted)

  if !isClassDeleted {
    _, err := database.DB.Exec("UPDATE classes SET deleted_at = NOW() WHERE id = $1", classID)

    if err != nil {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting class"})
      return
    }
  }else{
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Class already deleted"})
  }


  c.IndentedJSON(http.StatusNoContent, gin.H{})
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
