-- Create a new message
-- name: CreateMessage :one
INSERT INTO user_messages (name, email, message)
VALUES ($1, $2, $3)
RETURNING id, name, email, message;

-- Get a message by ID
-- name: GetMessageByID :one
SELECT id, name, email, message
FROM user_messages
WHERE id = $1;

-- Get messages by user email
-- name: GetMessagesByEmail :many
SELECT id, name, email, message
FROM user_messages
WHERE email = $1
ORDER BY id;

-- Update a message by ID
-- name: UpdateMessage :one
UPDATE user_messages
SET name = $2, email = $3, message = $4
WHERE id = $1
RETURNING id, name, email, message;

-- Delete a message by ID
-- name: DeleteMessage :exec
DELETE FROM user_messages
WHERE id = $1;

-- List all messages
-- name: ListMessages :many
SELECT id, name, email, message
FROM user_messages
ORDER BY id;

-- Create a new user
-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password;

-- Get a user by ID
-- name: GetUserByID :one
SELECT id, name, email, password
FROM users
WHERE id = $1;

-- Get a user by email
-- name: GetUserByEmail :one
SELECT id, name, email, password
FROM users
WHERE email = $1;

-- Update a user's name, email, and password by ID
-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, password = $4
WHERE id = $1
RETURNING id, name, email, password;

-- Delete a user by ID
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- List all users
-- name: ListUsers :many
SELECT id, name, email, password
FROM users
ORDER BY id;

-- Create a new search history entry
-- name: CreateSearchHistory :one
INSERT INTO search_history (user_id, keyword)
VALUES ($1, $2)
RETURNING id, user_id, keyword, search_time;

-- Get search history by user ID
-- name: GetSearchHistoryByUserID :many
SELECT id, user_id, keyword, search_time
FROM search_history
WHERE user_id = $1
ORDER BY search_time DESC;

-- Get search history by keyword
-- name: GetSearchHistoryByKeyword :many
SELECT id, user_id, keyword, search_time
FROM search_history
WHERE user_id = $1 AND keyword = $2
ORDER BY search_time DESC;

-- Delete search history by ID
-- name: DeleteSearchHistory :exec
DELETE FROM search_history
WHERE id = $1;

-- List all search history entries
-- name: ListSearchHistory :many
SELECT id, user_id, keyword, search_time
FROM search_history
ORDER BY search_time DESC;

-- Create a new website entry for a search
-- name: CreateWebsite :one
INSERT INTO websites (search_history_id, url)
VALUES ($1, $2)
RETURNING id, search_history_id, url;

-- Get websites by search history ID
-- name: GetWebsitesBySearchHistoryID :many
SELECT id, search_history_id, url
FROM websites
WHERE search_history_id = $1;

-- Delete websites by search history ID
-- name: DeleteWebsitesBySearchHistoryID :exec
DELETE FROM websites
WHERE search_history_id = $1;

-- Create a new email entry for a website
-- name: CreateEmail :one
INSERT INTO emails (website_id, email, subject, date)
VALUES ($1, $2, $3, $4)
RETURNING id, website_id, email, subject, date;

-- Get emails by website ID
-- name: GetEmailsByWebsiteID :many
SELECT id, website_id, email, subject, date
FROM emails
WHERE website_id = $1;

-- Delete emails by website ID
-- name: DeleteEmailsByWebsiteID :exec
DELETE FROM emails
WHERE website_id = $1;
