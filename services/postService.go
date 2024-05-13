package services

import (
	"errors"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"

	"github.com/google/uuid"
)


func CreatePost(post models.Post)(*models.Post, error){
  if  len(post.Teachers) < 0 || post.Content == "" || post.Name == "" {
    return nil, errors.New("Missing Fields")
  }
  post.ID = uuid.NewString()

  err := repositories.CreatePost(post)

  if err != nil{
    return nil, err
  }

  return &post, nil

}
