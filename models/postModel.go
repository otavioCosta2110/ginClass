package models

type Post struct {
  ID string `json:"id"`
  Name string `json:"name"`
  ClassID string `json:"classid"`
  Content string `json:"content"`
  Material []string `json:"material"`
  Tags []string `json:"tags"`
  Teachers []string `json:"teachers"`
}
