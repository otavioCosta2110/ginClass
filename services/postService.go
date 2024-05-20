package services

import (
	"errors"
	"otaviocosta2110/ginClass/models"
	"otaviocosta2110/ginClass/repositories"

	"github.com/google/uuid"
)

func GetPostByClass(classID string) ([]models.Post, error){
  posts, err := repositories.GetPostByClass(classID)

  if err != nil {
    return nil, errors.New("Error getting posts")
  }

  return posts, nil
}

func GetPostById(id string) (*models.Post, error){
  post, err := repositories.GetPostById(id)

  if err != nil {
    return nil, errors.New("Error getting posts")
  }

  return post, nil
}

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
