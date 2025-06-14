-- name: UpdateUserRole :one
UPDATE user_roles
SET role_id = ?
WHERE user_id = ?
RETURNING *;


-- name: AddUserRole :exec
INSERT INTO 
user_roles(user_id, role_id)
values(?,?);

