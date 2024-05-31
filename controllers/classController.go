package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/models"
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
      c.IndentedJSON(http.StatusInternalServerError, err.Error())
      return
    }


    c.IndentedJSON(http.StatusCreated, class)
}

func DeleteClass(c *gin.Context) {
  classID := c.Param("id")

  class, err := services.DeleteClass(classID)

  if err != nil{
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
  }

  c.IndentedJSON(http.StatusOK, gin.H{"message": class})

}

func GetAllClasses(c *gin.Context) {
  classes, err := services.GetAllClasses()
  if err != nil{
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    return
  }

  c.IndentedJSON(http.StatusOK, gin.H{"data": classes})
}
