package repositories

import (
  "errors"
  "otaviocosta2110/ginClass/database"
  "otaviocosta2110/ginClass/models"

)

func GetUserByEmail(email string) (*models.User, error) {
  rows, err := database.DB.Query("SELECT id, email, name, isteacher FROM users WHERE email = $1", email)

  if err != nil {
    return nil, errors.New("Error getting user")
  }

  defer rows.Close()
  var user models.User

  for rows.Next() {

    if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.IsTeacher); err != nil {
      return nil, errors.New("Error Scanning user")
    }
  
    return &user, nil
    
  }

  return nil, nil
}

func GetAllUsers() (*[]models.User, error){
  rows, err := database.DB.Query("SELECT id, name FROM users")
  
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  var users []models.User

  for rows.Next() {
    var user models.User
    if err := rows.Scan(&user.ID, &user.Name); err != nil {
      return nil, errors.New("Error getting users")
    }
    users = append(users, user)
  }
  return &users, nil
}

func CreateUser(user models.User) ( error) {

  tx, err := database.DB.Begin()

  if err != nil {
    tx.Rollback()
    return err
  }
  
  _, err = database.DB.Exec("INSERT INTO users (id, name, email, password, isteacher) VALUES ($1, $2, $3, $4, $5)", user.ID, user.Name, user.Email, user.Password, user.IsTeacher)

  if err != nil {
    tx.Rollback()
    return err
  }

  return nil
}

func AddUser(teacherEmail string, classID string) (error) {
  println(classID)
  _, err := database.DB.Exec("UPDATE classes SET teachers = array_append(teachers, $1) WHERE id = $2", teacherEmail, classID)

  if err != nil {
    return err
  }
  return nil
}
