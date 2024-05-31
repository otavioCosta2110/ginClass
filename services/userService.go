package services

import (
	"errors"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(email string) (*models.User, error) {
  user, err := repositories.GetUserByEmail(email)

  if err != nil {
    return nil, err;
  }
  return user, nil
}

func GetAllUsers() (*[]string , error) {
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
  hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost);

  if err != nil {
    return nil, errors.New("Error hashing password")
  }

  hashedPassword := string(hashedPasswordBytes)
  user.Password = hashedPassword

  foundUser, _ := repositories.GetUserByEmail(user.Email);

  if foundUser != nil {
    return nil, errors.New("User Already Exists")
  }

  err = repositories.CreateUser(user)

  if err != nil {
    return nil, err
  }

  return &user, nil
}

func Login(userEmail string, password string) (*models.User, error) {
  user, err := GetUserByEmail(userEmail)

  if err != nil {
    return nil, err
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if err != nil {
    return nil, errors.New("Invalid Password")
  }

  return user, nil
}

