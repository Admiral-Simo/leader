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
