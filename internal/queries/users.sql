-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    password,
    role,
    status
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
) RETURNING *;

-- name: UserExists :one
SELECT CAST(EXISTS (
    SELECT 1 FROM users 
    WHERE id = ?
) AS BOOLEAN) AS user_exists;


-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = ? LIMIT 1;


-- name: GetUser :one
SELECT
    u.*,
    um.meta_value AS user_status
FROM
    users u
INNER JOIN
    users_meta um ON u.id = um.user_id AND um.meta_key = 'user_status'
WHERE
    u.id = ?;

-- name: UpdateUser :one
UPDATE users
SET
    username = ?,
    email = ?,
    password = ?,
    role = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;


-- name: UpdateUserPassword :exec
UPDATE users
SET
    password = ?,
    password_updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserLastLogin :exec
UPDATE users
SET
    last_login = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;



-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: GetUsers :many
SELECT
    u.*
FROM
    users u
WHERE
    (@username = '' OR u.username LIKE '%' || @username || '%' ) AND
    (@status = '' OR u.status = @status) AND
    (@role = '' OR u.role = @role)
LIMIT @limit OFFSET @offset;