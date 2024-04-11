package controllers

import (
  "net/http"
  "otaviocosta2110/ginClass/database"
  "otaviocosta2110/ginClass/models"
  "otaviocosta2110/ginClass/repositories"

  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
)

func GetUserByEmail(c *gin.Context){
  email := c.Param("email")
  println(email)
  user, err := repositories.UserByEmail(email)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting user by email"})
    panic(err)
  }
  c.IndentedJSON(http.StatusCreated, &user)
}

func GetAllUsers(c *gin.Context) {
  rows, err := database.DB.Query("SELECT id, name FROM users")

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": "Error getting users" })
    return 
  }

  defer rows.Close()

  var users []models.User

  for rows.Next() {
    var user models.User
    if err := rows.Scan(&user.ID, &user.Name); err != nil {
      c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": "Error scanning users" })
      return
    }
    users = append(users, user)
  }

  if err := rows.Err(); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": "Error iterating users" })
    return
  }

  c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
  var user models.User
  c.BindJSON(&user)

  user.ID = uuid.New().String()

  foundUser, _ := repositories.UserByEmail(user.Email);

  if foundUser != nil {
    println(foundUser)
    c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User Already Exists"})
    return
  }

  _, err := database.DB.Exec("INSERT INTO users (id, name, email, password, isteacher) VALUES ($1, $2, $3, $4, $5)", user.ID, user.Name, user.Email, user.Password, user.IsTeacher)

  if err != nil {
    println(err.Error())
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
    return
  }

  c.IndentedJSON(http.StatusCreated, user)
}
