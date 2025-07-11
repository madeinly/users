// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: user_session.sql

package userQuery

import (
	"context"
)

const cleanupExpiredSessions = `-- name: CleanupExpiredSessions :exec
DELETE FROM user_sessions 
WHERE expires_at < CURRENT_TIMESTAMP
   OR last_accessed_at < datetime('now', '-1 year')
`

func (q *Queries) CleanupExpiredSessions(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanupExpiredSessionsStmt, cleanupExpiredSessions)
	return err
}

const createSession = `-- name: CreateSession :exec
INSERT INTO user_sessions (
    id,
    user_id,
    token,
    session_data,
    expires_at
) VALUES (
    ?1, ?2, ?3, ?4, ?5
)
`

type CreateSessionParams struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Token       string `json:"token"`
	SessionData string `json:"session_data"`
	ExpiresAt   string `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.exec(ctx, q.createSessionStmt, createSession,
		arg.ID,
		arg.UserID,
		arg.Token,
		arg.SessionData,
		arg.ExpiresAt,
	)
	return err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM user_sessions
WHERE id = ?1
`

func (q *Queries) DeleteSession(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteSessionStmt, deleteSession, id)
	return err
}

const getSessionByToken = `-- name: GetSessionByToken :one
SELECT id, user_id, token, session_data, created_at, expires_at, last_accessed_at FROM user_sessions
WHERE token = ?1 LIMIT 1
`

func (q *Queries) GetSessionByToken(ctx context.Context, token string) (UserSession, error) {
	row := q.queryRow(ctx, q.getSessionByTokenStmt, getSessionByToken, token)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.SessionData,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.LastAccessedAt,
	)
	return i, err
}

const getSessionByUserID = `-- name: GetSessionByUserID :one
SELECT id, user_id, token, session_data, created_at, expires_at, last_accessed_at FROM user_sessions
WHERE user_id = ?1 Limit 1
`

func (q *Queries) GetSessionByUserID(ctx context.Context, userID string) (UserSession, error) {
	row := q.queryRow(ctx, q.getSessionByUserIDStmt, getSessionByUserID, userID)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.SessionData,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.LastAccessedAt,
	)
	return i, err
}

const updateSessionData = `-- name: UpdateSessionData :exec
UPDATE user_sessions
SET 
    session_data = ?1
WHERE id = ?2
`

type UpdateSessionDataParams struct {
	SessionData string `json:"session_data"`
	ID          string `json:"id"`
}

func (q *Queries) UpdateSessionData(ctx context.Context, arg UpdateSessionDataParams) error {
	_, err := q.exec(ctx, q.updateSessionDataStmt, updateSessionData, arg.SessionData, arg.ID)
	return err
}

const updateSessionToken = `-- name: UpdateSessionToken :one
UPDATE user_sessions
SET 
    token = ?1,
    expires_at = ?2
WHERE user_id = ?3
RETURNING id, user_id, token, session_data, created_at, expires_at, last_accessed_at
`

type UpdateSessionTokenParams struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
	UserID    string `json:"user_id"`
}

func (q *Queries) UpdateSessionToken(ctx context.Context, arg UpdateSessionTokenParams) (UserSession, error) {
	row := q.queryRow(ctx, q.updateSessionTokenStmt, updateSessionToken, arg.Token, arg.ExpiresAt, arg.UserID)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.SessionData,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.LastAccessedAt,
	)
	return i, err
}
