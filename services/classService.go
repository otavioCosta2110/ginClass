package services

import (
	"errors"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"
)

func GetClassByTeacher(teacherEmail string) (*[]models.Class, error) {
  classes, err := repositories.GetClassByTeacher(teacherEmail)

  if err != nil {
    println(err.Error())
    return nil, errors.New("Error getting classes")
  }

  return classes, err
}
