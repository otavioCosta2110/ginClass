package repositories

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTags (tagContent string, tx *sql.Tx, c *gin.Context) (string, error) {
  tagID := uuid.NewString();
  _, err := tx.Exec("INSERT INTO tags (id, name) values ($1, $2)", tagID, tagContent)

  if err != nil {
    tx.Rollback()
    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating tag"})
    return "", err
  }

  println("peeeeeenisiasidjoaisdiojio",tagID)
  return tagID, nil
}
