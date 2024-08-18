CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  email text      NOT NULL,
  password text      NOT NULL
);

CREATE TABLE user_messages (
  id      BIGSERIAL PRIMARY KEY,
  name    TEXT NOT NULL,
  email   TEXT NOT NULL,
  message TEXT NOT NULL
);
