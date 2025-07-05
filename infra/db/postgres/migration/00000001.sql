-- Create url_shortening table with foreign key reference to users
CREATE TABLE url_shortening (
  id varchar(255) PRIMARY KEY,
  id_user varchar(255) NOT NULL,
  url_original varchar(255) NOT NULL,
  url_shortened varchar(255) NOT NULL UNIQUE,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  
  FOREIGN KEY (id_user) REFERENCES users(id)
);
