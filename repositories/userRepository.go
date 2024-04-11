package repositories

import (
	"errors"
	"otaviocosta2110/ginClass/database"
	"otaviocosta2110/ginClass/models"

)

func UserByEmail(email string) (*models.User, error) {
  rows, err := database.DB.Query("SELECT email FROM users WHERE email = $1", email)

  if err != nil {
    return nil, errors.New("Error getting user")
  }

  defer rows.Close()
  var user models.User

  for rows.Next() {
    if err := rows.Scan(&user.Email); err != nil {
      return nil, errors.New("Error Scanning user")
    }
    return &user, nil
  }
  return nil, nil
}
