package services

import (
	"errors"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"

	"github.com/google/uuid"
)

func GetClassByTeacher(teacherEmail string) (*[]models.Class, error) {
  classes, err := repositories.GetClassByTeacher(teacherEmail)

  if err != nil {
    println(err.Error())
    return nil, errors.New("Error getting classes")
  }

  return classes, err
}

func CreateClass(class models.Class) (*models.Class, error){
    if len(class.Teachers) < 1 || class.Name == "" {
      return nil, errors.New("Missing fields")
    }

    class.ID = uuid.NewString()

    users := append(class.Teachers, class.Students...)

    repositories.CreateClass(class, users)

    return &class, nil
}

func AddTeacher(body models.AddUser) (error){
  teacher, err := repositories.GetUserByEmail(body.TeacherEmail)

  if err != nil {
    return err
  }

  if !teacher.IsTeacher {
    return errors.New("User is not a teacher!")
  }

  err = repositories.AddUser(teacher.ID, body.ClassID)
  if err != nil {
    return err
  }
  return nil

}
