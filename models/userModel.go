package models

type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  Password string `json:"password"`
  IsTeacher bool `json:"isteacher"`
}

type UserLogin struct {
  Email string `json:"email"`
  Password string `json:"password"`
}
