package repositories

import (
	"database/sql"
	"errors"

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

  print("Tag exists")
  return id, true, nil
}
