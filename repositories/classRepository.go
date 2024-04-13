package repositories

import (
	"errors"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"
)

func ClassByTeacher(teacherEmail string) (*[]models.Class, error){
  rows, err := database.DB.Query("SELECT name FROM classes WHERE $1 = ANY(teachers)", teacherEmail)
  if err != nil {
    return nil, errors.New("Error Getting Class")
  }

  defer rows.Close()
  var class models.Class
  var classes []models.Class

  for rows.Next() {
    if err := rows.Scan(&class.Name); err != nil {
      return nil, errors.New("Error Scanning classes")
    }
    classes = append(classes, class)
  }
  return &classes, nil
}
