CREATE TABLE materials (
  id VARCHAR(255) primary key,
  content VARCHAR(255)
)
CREATE TABLE post_material (
  post_id varchar(255),
  material_id varchar(255)
)
