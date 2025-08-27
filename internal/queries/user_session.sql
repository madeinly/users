-- name: CreateSession :exec
INSERT INTO user_sessions (
    id,
    user_id,
    token,
    session_data,
    expires_at
) VALUES (
    ?1, ?2, ?3, ?4, ?5
);

-- name: GetSessionByUserID :one
SELECT * FROM user_sessions
WHERE user_id = ?1 Limit 1;


-- name: GetSessionBySessionToken :one
SELECT * FROM user_sessions
WHERE token = ?1 Limit 1;


-- name: UpdateSessionToken :one
UPDATE user_sessions
SET 
    token = ?1,
    expires_at = ?2
WHERE user_id = ?3
RETURNING *;


-- name: UpdateSessionData :exec
UPDATE user_sessions
SET 
    session_data = ?1
WHERE id = ?2;

-- name: GetSessionByToken :one
SELECT * FROM user_sessions
WHERE token = ?1 LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM user_sessions
WHERE id = ?1;

-- name: CleanupExpiredSessions :exec
DELETE FROM user_sessions 
WHERE expires_at < CURRENT_TIMESTAMP
   OR last_accessed_at < datetime('now', '-1 year');