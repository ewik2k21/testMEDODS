CREATE TABLE users
(
  id serial not null unique,
  username varchar(255) not null,
  email varchar(255) not null unique,
  password_hash varchar(255) not null, 
  refresh_token varchar(255)
);