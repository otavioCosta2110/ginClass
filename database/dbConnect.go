package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func init(){
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  host:= os.Getenv("DATABASE_HOST")
  dbName:= os.Getenv("DATABASE_NAME")
  user:= os.Getenv("DATABASE_USER")
  password:= os.Getenv("DATABASE_PASSWORD")

  db, err := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", host, dbName, user, password))

  if err != nil {
    panic(err)
  }
  DB = db
}
