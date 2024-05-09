package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"
	"otaviocosta2110/ginClass/services"

	"github.com/gin-gonic/gin"
)

func GetClassByTeacher(c *gin.Context) {
  teacherEmail := c.Param("teacheremail")

  classes, err := services.GetClassByTeacher(teacherEmail)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
  }

  c.IndentedJSON(http.StatusOK, &classes)
}

func AddTeacher(c *gin.Context) {
  var body models.AddUser

  c.BindJSON(&body)
  
  err := services.AddTeacher(body)
  if err != nil{
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
  }
  c.IndentedJSON(http.StatusOK, gin.H{"message": "User Added!"})
}

func CreateClass(c *gin.Context) {
    var class models.Class

    c.BindJSON(&class)
    err := services.CreateClass(class)
    if err != nil{
      c.IndentedJSON(http.StatusInternalServerError, err)
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
  classes, err := services.GetAllClasses()
  if err != nil{
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    return
  }

  c.IndentedJSON(http.StatusOK, classes)
}
