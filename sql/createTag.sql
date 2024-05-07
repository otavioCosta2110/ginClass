CREATE TABLE tags (
  id VARCHAR(255) primary key,
  name VARCHAR(255)
)

CREATE TABLE class_tag (
  class_id VARCHAR(255),
  tag_id VARCHAR(255),
  primary key(class_id, tag_id)
)
