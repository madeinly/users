-- name: AddUserMeta :one
INSERT INTO users_meta (
    user_id,
    meta_key,
    meta_value
) VALUES (
    ?, ?, ?
)
ON CONFLICT (user_id, meta_key) 
DO UPDATE SET meta_value = EXCLUDED.meta_value
RETURNING *;


-- name: UpdateUserMeta :one
UPDATE users_meta
SET meta_value = ?
WHERE user_id = ?
RETURNING *;



-- name: DeleteMetas :exec
DELETE FROM users_meta
WHERE user_id = ?;