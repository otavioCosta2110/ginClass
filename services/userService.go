package services

import (
	"errors"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"

	"github.com/google/uuid"
)

func GetUserByEmail(email string) (*models.User, error) {
  user, err := repositories.GetUserByEmail(email)

  if err != nil {
    return nil, err;
  }
  return user, nil
}

func GetAllUsers() (*[]models.User, error) {
  users, err := repositories.GetAllUsers();

  if err != nil {
    return nil, err
  }
  return users, nil

}

func CreateUser(user models.User) (*models.User, error) {
  if user.Email == ""|| user.Name == "" || user.Password == "" {
    return nil, errors.New("Missing fields")
  }

  user.ID = uuid.New().String()

  foundUser, _ := repositories.GetUserByEmail(user.Email);

  if foundUser != nil {
    return nil, errors.New("User Already Exists")
  }

  err := repositories.CreateUser(user)

  if err != nil {
    return nil, err
  }

  return &user, nil
}
