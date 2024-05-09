package repositories

import (
  "database/sql"
  "errors"
  "otaviocosta2110/ginClass/database"
  "otaviocosta2110/ginClass/models"
)

func GetClassByTeacher(teacherEmail string) (*[]models.Class, error){
  teacher, err := GetUserByEmail(teacherEmail)

  if err != nil {
    return nil, errors.New("Error Getting Class")
  }

  rows, err := database.DB.Query("SELECT class_id FROM user_class WHERE user_id = $1", teacher.ID)
  if err != nil {
    return nil, errors.New("Error Getting Class")
  }

  defer rows.Close()
  var class models.Class
  var classes []models.Class

  for rows.Next() {
    if err := rows.Scan(&class.ID); err != nil {
      return nil, errors.New("Error Scanning classes")
    }

    nameRow, err := database.DB.Query("SELECT name FROM classes WHERE id = $1", class.ID)
    if err != nil {
      return nil, errors.New("Error Getting Class")
    }
    nameRow.Next()
    if err := nameRow.Scan(&class.Name); err != nil {
      return nil, errors.New("Error Scanning classes")
    }
    classes = append(classes, class)
  }
  println(classes)
  return &classes, nil
}

func IsClassDeleted(id string) (bool, error) {
  var deleted sql.NullTime

  err := database.DB.QueryRow("SELECT deleted_at FROM classes WHERE id = $1", id).Scan(&deleted)
  println("ID: ", id, " Deleted: ", deleted.Valid)

  if err != nil {
    println("error: ", err.Error())
    return false, errors.New("Error checking if class is deleted")
  }

  return deleted.Valid, nil
}
