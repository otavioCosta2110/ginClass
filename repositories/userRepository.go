package repositories

import (
  "errors"
  "otaviocosta2110/ginClass/database"
  "otaviocosta2110/ginClass/models"

)

func UserByEmail(email string) (*models.User, error) {
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

func AddUser(teacherEmail string, classID string) (error) {
  println(classID)
  _, err := database.DB.Exec("UPDATE classes SET teachers = array_append(teachers, $1) WHERE id = $2", teacherEmail, classID)

  if err != nil {
    return err
  }
  return nil
}
