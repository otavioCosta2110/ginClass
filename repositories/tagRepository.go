package repositories

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTags (tagContent string, tx *sql.Tx, c *gin.Context) (string, error) {
  
  tagID, tagExists, err := TagExists(tagContent, tx)

  println("tagID", tagID)
  if err != nil {
    tx.Rollback()
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error checking if tag exists"})
    return "", err
  }

  if !tagExists {
    _, err = tx.Exec("INSERT INTO tags (id, name) values ($1, $2)", tagID, tagContent)
    if err != nil {
      tx.Rollback()
      print("error", err.Error())
      c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating tag"})
      return "", err
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
