CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  email text      NOT NULL,
  password text      NOT NULL
);

CREATE TABLE search_history (
  id        BIGSERIAL PRIMARY KEY,
  user_id   BIGINT REFERENCES users(id) ON DELETE CASCADE,
  keyword   TEXT NOT NULL,
  search_time TIMESTAMPTZ DEFAULT NOW(),
  UNIQUE(user_id, keyword)
);

CREATE TABLE websites (
  id              BIGSERIAL PRIMARY KEY,
  search_history_id BIGINT REFERENCES search_history(id) ON DELETE CASCADE,
  url             TEXT NOT NULL,
  UNIQUE(search_history_id, url)
);

CREATE TABLE emails (
  id           BIGSERIAL PRIMARY KEY,
  website_id   BIGINT REFERENCES websites(id) ON DELETE CASCADE,
  email        TEXT NOT NULL,
  subject      TEXT,
  date         TIMESTAMPTZ,
  UNIQUE(website_id, email)
);

CREATE TABLE user_messages (
  id      BIGSERIAL PRIMARY KEY,
  name    TEXT NOT NULL,
  email   TEXT NOT NULL,
  message TEXT NOT NULL
);
