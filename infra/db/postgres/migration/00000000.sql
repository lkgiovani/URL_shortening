-- Drop tables in correct order (dependent tables first)
DROP TABLE IF EXISTS url_shortening;
DROP TABLE IF EXISTS users;

-- Create users table first (referenced table)
CREATE TABLE users (
  id varchar(255) PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);