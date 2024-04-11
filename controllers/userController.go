package controllers

import (
	"net/http"
	"otaviocosta2110/ginClass/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
  {ID: "1", Name: "Otavio", Email: "otavio@email.com", Password: "123"},
}

func GetUsers(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
  var user models.User
  c.BindJSON(&user)
  users = append(users, user)
  c.IndentedJSON(http.StatusCreated, user)
}
