package repositories

import (
  "database/sql"
  "errors"
  "otaviocosta2110/ginClass/database"
  "otaviocosta2110/ginClass/models"
)

func GetClassByTeacher(teacherEmail string) (*[]models.Class, error){
  teacher, err := GetUserByEmail(teacherEmail)

  if err != nil {
    return nil, errors.New("Error Getting Class")
  }

  rows, err := database.DB.Query("SELECT class_id FROM user_class WHERE user_id = $1", teacher.ID)
  if err != nil {
    return nil, errors.New("Error Getting Class")
  }

  defer rows.Close()
  var class models.Class
  var classes []models.Class

  for rows.Next() {
    if err := rows.Scan(&class.ID); err != nil {
      return nil, errors.New("Error Scanning classes")
    }

    nameRow, err := database.DB.Query("SELECT name FROM classes WHERE id = $1", class.ID)
    if err != nil {
      return nil, errors.New("Error Getting Class")
    }
    nameRow.Next()
    if err := nameRow.Scan(&class.Name); err != nil {
      return nil, errors.New("Error Scanning classes")
    }
    classes = append(classes, class)
  }
  println(classes)
  return &classes, nil
}

func GetClassByID(classID string) (*models.Class, error){
  rows, err := database.DB.Query("SELECT id, name FROM classes WHERE id = $1", classID)

  if err != nil {
    return nil, errors.New("Error getting class")
  }

  defer rows.Close()
  var class models.Class

  for rows.Next() {

    if err := rows.Scan(&class.ID, &class.Name); err != nil {
      return nil, errors.New("Error Scanning class")
    }
  
    return &class, nil
    
  }

  return nil, nil

}

func IsClassDeleted(id string) (bool, error) {
  var deleted sql.NullTime

  err := database.DB.QueryRow("SELECT deleted_at FROM classes WHERE id = $1", id).Scan(&deleted)
  println("ID: ", id, " Deleted: ", deleted.Valid)

  if err != nil {
    println("error: ", err.Error())
    return false, errors.New("Error checking if class is deleted")
  }

  return deleted.Valid, nil
}

func AddUser(teacherID string, classID string) (error) {
  println(classID)
  _, err := database.DB.Exec("INSERT INTO user_class (user_id, class_id) values ($1, $2)", teacherID, classID)

  if err != nil {
    return errors.New("Error inserting user into database")
  }
  return nil
}


func CreateClass(class models.Class, users []string) (error){
  tx, err := database.DB.Begin()
  if err != nil {
    tx.Rollback()
    return errors.New("Error connecting to database")
  }

  // defer func() {
  //   if r := recover(); r != nil {
  //     tx.Rollback()
  //     return errors.New("Error connecting to database")
  //   }
  // }()

  _, err = tx.Exec("INSERT INTO classes (id, name) values ($1, $2)", class.ID, class.Name)
  if err != nil {
    tx.Rollback()
    return errors.New("Error creating class")
  }


  for _, userEmail := range users {
    user, err := GetUserByEmail(userEmail)
    if user == nil {
      tx.Rollback()
      return errors.New("User with email " + userEmail + " does not exist")
    }
    if err != nil {
      tx.Rollback()
      return errors.New("Error getting teacher ID")
    }
    _, err = tx.Exec("INSERT INTO user_class (user_id, class_id) values ($1, $2)", user.ID, class.ID)
    if err != nil {
      tx.Rollback()
      return errors.New("Error creating class")
    }
  }

  for _, tagContent := range class.Tags {
    tagID, err := CreateTags(tagContent, tx)
    if err != nil {
      tx.Rollback()
      return errors.New("Error creating tag")
    }

    _, err = tx.Exec("INSERT INTO class_tag (class_id, tag_id) values ($1, $2)", class.ID, tagID)
    if err != nil {
      tx.Rollback()
      return errors.New("Error creating class_tag")
    }
  }

  err = tx.Commit()
  if err != nil {
    tx.Rollback()
    return errors.New("Error committing transaction")
  }
  return nil
}

func GetAllClasses() (*[]models.Class, error) {
  rows, err := database.DB.Query("SELECT id, name FROM classes")

  if err != nil {
    return nil, errors.New("Error getting classes")
  }

  defer rows.Close()

  var classes []models.Class

  for rows.Next() {
    var class models.Class
    if err := rows.Scan(&class.ID, &class.Name); err != nil{
      return nil, errors.New("Error scanning classes")
    }
    userRow, err := database.DB.Query("SELECT user_id FROM user_class WHERE class_id = $1", class.ID)
    if err != nil{
      return nil, errors.New("Error getting users")
    }
    for userRow.Next() {
      var userID string
      if err := userRow.Scan(&userID); err != nil{
        return nil, errors.New("Error scanning users")
      }
      user, err := GetUserByID(userID)
      if err != nil {
        return nil, err
      }

      if user.IsTeacher{
        class.Teachers = append(class.Teachers, user.Email)
      }else {
        class.Students = append(class.Teachers, user.Email)
      }
    }
    tagsRow, err := database.DB.Query("SELECT tag_id FROM class_tag WHERE class_id = $1", class.ID)
    if err != nil{
      return nil, errors.New("Error getting users")
    }
    for tagsRow.Next() {
      var tagID string
      if err := tagsRow.Scan(&tagID); err != nil{
        return nil, errors.New("Error scanning tags")
      }
      tag, err := GetTagByID(tagID)
      if err != nil {
        return nil, err
      }
      class.Tags = append(class.Tags, tag.Content)
    }

    if err != nil {
      return nil, errors.New("Error getting classes")
    }
    classes = append(classes, class)
  }

  if err := rows.Err(); err != nil {
    return nil, errors.New("Error iterating classes")
  }

  return &classes, nil
}

func DeleteClass(classID string)(*models.Class, error){

  _, err := database.DB.Exec("UPDATE classes SET deleted_at = NOW() WHERE id = $1", classID)

  if err != nil {
    return nil, errors.New("Error deleting class")
  } 
  class, err := GetClassByID(classID)

  if err != nil {
    return nil, err
  }

  return class, nil
}
