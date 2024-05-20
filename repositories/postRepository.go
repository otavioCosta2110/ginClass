package repositories

import (
	"errors"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"

	"github.com/google/uuid"
)

func GetPostByClass (classID string) ([]models.Post, error) {
  rows, err := database.DB.Query("SELECT * FROM posts WHERE class_id = $1", classID)

  if err != nil {
    return nil, errors.New("Error getting posts")
  }

  var posts []models.Post

  for rows.Next() {
    var post models.Post
    err := rows.Scan(&post.ID, &post.Name, &post.ClassID, &post.Content)

    if err != nil {
      return nil, errors.New("Error scanning posts")
    }

    posts = append(posts, post)
  }

  return posts, nil
}

func GetPostById (id string) (*models.Post, error) {
  rows, err := database.DB.Query("SELECT * FROM posts WHERE id = $1", id)

  if err != nil {
    return nil, errors.New("Error getting posts")
  }
  var post models.Post

  for rows.Next() {
    err = rows.Scan(&post.ID, &post.Name, &post.ClassID, &post.Content)
    userRow, err := database.DB.Query("SELECT user_id FROM user_post WHERE post_id = $1", post.ID)
    for userRow.Next() {
      var userID string
      err = userRow.Scan(&userID)
      if err != nil {
        return nil, errors.New("Error scanning user_post")
      }
      user, err := GetUserByID(userID)
      if err != nil {
        return nil, err
      }
      post.Teachers = append(post.Teachers, user.Email)
    }
    
    materialRow, err := database.DB.Query("SELECT material_id FROM post_material WHERE post_id = $1", post.ID)
    for materialRow.Next() {
      var materialID string
      err = materialRow.Scan(&materialID)
      if err != nil {
        return nil, errors.New("Error scanning post_material")
      }
      post.Material = append(post.Material, materialID)
    }

    tagRow, err := database.DB.Query("SELECT tag_id FROM post_tag WHERE post_id = $1", post.ID)
    for tagRow.Next() {
      var tagID string
      err = tagRow.Scan(&tagID)
      if err != nil {
        return nil, errors.New("Error scanning post_tag")
      }
      post.Tags = append(post.Tags, tagID)
    }

    if err != nil {
      return nil, errors.New("Error scanning post")
    }
  }

  return &post, nil
}

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
