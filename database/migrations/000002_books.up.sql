-- +migrate Up
CREATE TABLE books (
  id  BIGSERIAL PRIMARY KEY,
  name text     NOT NULL,
  author_id INT,
  FOREIGN KEY (author_id) REFERENCES authors(id)
);
