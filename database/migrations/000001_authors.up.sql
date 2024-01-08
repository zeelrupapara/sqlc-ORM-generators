-- +migrate Up
CREATE TABLE authors (
  id   INT PRIMARY KEY,
  name text      NOT NULL,
  bio  text      NOT NULL
);
