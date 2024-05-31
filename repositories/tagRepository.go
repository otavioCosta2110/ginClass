package repositories

import (
	"database/sql"
	"errors"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"

	"github.com/google/uuid"
)

func CreateTags (tagContent string, tx *sql.Tx) (string, error) {
  
  tagID, tagExists, err := TagExists(tagContent, tx)

  if err != nil {
    tx.Rollback()
    return "", errors.New("Error checking if tag exists")
  }

  if !tagExists {
    _, err = tx.Exec("INSERT INTO tags (id, name) values ($1, $2)", tagID, tagContent)
    if err != nil {
      tx.Rollback()
      return "", errors.New("Error creating tags")
    }
  }

  return tagID, nil
}

func TagExists(tagContent string, tx *sql.Tx) (string, bool, error) {
  var id string
  err := tx.QueryRow("SELECT id FROM tags WHERE name = $1", tagContent).Scan(&id)

  if err != nil {
    if err == sql.ErrNoRows {
      print("Tag doesnt exists")
      return uuid.NewString(), false, nil
    }
    return "", false, err
  }

  return id, true, nil
}

func GetTagByID(id string) (*models.Tag, error){
  rows, err := database.DB.Query("SELECT id, name FROM tags WHERE id = $1", id)
  
  if err != nil{
    return nil, errors.New("Error getting tag")
  }

  defer rows.Close()
  var tag models.Tag

  for rows.Next() {

    if err := rows.Scan(&tag.ID, &tag.Content); err != nil {
      return nil, err
    }
  
    return &tag, nil
    
  }

  return nil, nil
}

func GetAllTags()([]string, error) {
  rows, err := database.DB.Query("SELECT name FROM tags")
  
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  var tags []string

  for rows.Next() {
    var tag string
    if err := rows.Scan(&tag); err != nil {
      return nil, err
    }
    tags = append(tags, tag)
  }
  return tags, nil
}
