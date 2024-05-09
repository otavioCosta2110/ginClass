CREATE TABLE classes (
  id VARCHAR(255) primary key,
  name VARCHAR(255)
)

CREATE TABLE user_class (
  user_id varchar(255),
  class_id varchar(255),
  
  primary key(user_id, class_id)
)
