-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    password
) VALUES (
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
    u.id,
    u.email,
    u.username,
    u.password,
    ur.role_id,
    um.meta_value AS user_status
FROM
    users u
INNER JOIN
    user_roles ur ON u.id = ur.user_id
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
    u.id,
    u.username,
    u.email,
    u.password,
    ur.role_id,
    um.meta_value AS status_name,
    u.created_at
FROM
    users u
INNER JOIN user_roles ur ON u.id = ur.user_id
LEFT JOIN users_meta um ON u.id = um.user_id AND um.meta_key = 'user_status'
WHERE
    (:username = '' OR u.username LIKE '%' || :username || '%') AND
    (:role_id = 0 OR ur.role_id = :role_id) AND 
    (:status = '' OR COALESCE(um.meta_value, 'active') = :status)
LIMIT :limit OFFSET :offset;