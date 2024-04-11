package models

type Class struct { 
  ID string `json:"id"`
  Name string `json:"name"`
  Teachers []string `json:"teachers"`
  Students []string `json:"students"`
  Posts []string `json:"posts"`
}
