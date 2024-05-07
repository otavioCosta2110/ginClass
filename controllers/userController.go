package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/services"
	"github.com/gin-gonic/gin"
)

func GetUserByEmail(c *gin.Context){
  email := c.Param("email")
  user, err := services.GetUserByEmail(email)
  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting user by email"})
  }
  c.IndentedJSON(http.StatusOK, &user)
}

func GetAllUsers(c *gin.Context) {
  users, err := services.GetAllUsers()
  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting users"})
  }

  c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
  var user models.User
  c.BindJSON(&user)
  createdUser, err := services.CreateUser(user)

  if err != nil{
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
  }

  c.IndentedJSON(http.StatusOK, gin.H{"message": "User " + createdUser.Name +" created!"})
}
