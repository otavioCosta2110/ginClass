package repositories

import (
	"errors"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"

	"github.com/google/uuid"
)

func CreatePost (post models.Post) (error) {

  tx, err := database.DB.Begin()

  if err != nil {
    tx.Rollback()
    return errors.New("Error connecting to database")
  }

  _, err = database.DB.Query("INSERT INTO posts (id, name, class_id, content) values ($1, $2, $3, $4)", post.ID, post.Name, post.ClassID, post.Content)

  if err != nil {
    return errors.New("Error Creating Post")
  }

  for _, tagContent := range post.Tags {
    tagID, err := CreateTags(tagContent, tx)

    if err != nil {
      tx.Rollback()
      return errors.New("Error Creating Tags")
    }
    _, err = tx.Exec("INSERT INTO post_tag (post_id, tag_id) values ($1, $2)", post.ID, tagID)

    if err != nil {
      tx.Rollback()
      return errors.New("Error Creating post_tag")
    }

  }

  for _, userEmail := range post.Teachers {
    user, err := GetUserByEmail(userEmail)
    if err != nil {
      return err
    }

    _, err = database.DB.Query("INSERT INTO user_post (user_id, post_id) values ($1, $2)", user.ID, post.ID)

    if err != nil {
      return errors.New("Error Creating user_post")
    }
  }

  for _, material := range post.Material {
    materialID := uuid.NewString()
    _, err = database.DB.Query("INSERT INTO materials (id, content) values ($1, $2)", materialID, material)
    println(material)

    if err != nil {
      return errors.New("Error Creating material")
    }

    _, err = database.DB.Query("INSERT INTO post_material (post_id, material_id) values ($1, $2)", post.ID, materialID)

    if err != nil {
      return errors.New("Error Creating post_material")
    }
  }

  err = tx.Commit()
  if err != nil {
    tx.Rollback()
    return errors.New("Error commiting transaction")
  }
  return nil

}
