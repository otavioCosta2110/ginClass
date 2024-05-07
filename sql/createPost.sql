CREATE TABLE posts (
  id VARCHAR(255) primary key,
  name VARCHAR(255),
  classid VARCHAR(255),
  content VARCHAR(255)
)
CREATE TABLE post_tag (
  post_id varchar(255),
  tag_id varchar(255)
)
CREATE TABLE user_post (
  post_id varchar(255),
  user_id varchar(255)
)
