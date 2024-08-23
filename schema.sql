-- Create users table
CREATE TABLE users (
  id       BIGSERIAL PRIMARY KEY,
  name     TEXT NOT NULL,
  email    TEXT NOT NULL,
  password TEXT NOT NULL,
  credits  INTEGER DEFAULT 100
);

-- Create search_history table
CREATE TABLE search_history (
  id         BIGSERIAL PRIMARY KEY,
  user_id    BIGINT REFERENCES users(id) ON DELETE CASCADE,
  keyword    TEXT NOT NULL,
  search_time TIMESTAMPTZ DEFAULT NOW()
);

-- Create websites table
CREATE TABLE websites (
  id               BIGSERIAL PRIMARY KEY,
  search_history_id BIGINT REFERENCES search_history(id) ON DELETE CASCADE,
  url              TEXT NOT NULL,
  UNIQUE(search_history_id, url)
);

-- Create emails table
CREATE TABLE emails (
  id         BIGSERIAL PRIMARY KEY,
  website_id BIGINT REFERENCES websites(id) ON DELETE CASCADE,
  email      TEXT NOT NULL,
  date       TIMESTAMPTZ,
  UNIQUE(website_id, email)
);

-- Create user_messages table
CREATE TABLE user_messages (
  id      BIGSERIAL PRIMARY KEY,
  name    TEXT NOT NULL,
  email   TEXT NOT NULL,
  message TEXT NOT NULL
);
