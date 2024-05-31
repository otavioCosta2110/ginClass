package services


import (
	"otaviocosta2110/ginClass/repositories"

)

func GetAllTags() ([]string, error) {
  tags, err := repositories.GetAllTags()

  if err != nil {
    return nil, err;
  }
  return tags, nil
}

