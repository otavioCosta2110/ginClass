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

func CreateClass(class models.Class) (error){
    if len(class.Teachers) < 1 || class.Name == "" {
      return errors.New("Missing fields")
    }

    class.ID = uuid.NewString()

    users := append(class.Teachers, class.Students...)

    err := repositories.CreateClass(class, users)
    if err != nil{
      return err
    }

    return nil
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

func GetAllClasses() (*[]models.Class, error){
  classes, err := repositories.GetAllClasses()
  if err != nil {
    return nil, err
  }

  return classes, nil
}

func DeleteClass(classID string) (*models.Class, error){
  isClassDeleted, err := repositories.IsClassDeleted(classID)
  if err != nil {
    return nil, err
  }


  if !isClassDeleted {
    class, err := repositories.DeleteClass(classID)

    if err != nil {
      return nil, err
    }
    return class, nil
  }else{
    return nil, errors.New("Class already deleted")
  }
}
