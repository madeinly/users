package models

import (
	"context"
	"time"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/queries/userQuery"
)

type userSessionRepo struct {
	ID             string    `json:"id" db:"id"`
	UserID         string    `json:"user_id" db:"user_id"`
	Token          string    `json:"token" db:"token"`
	SessionData    string    `json:"session_data" db:"session_data"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	ExpiresAt      time.Time `json:"expires_at" db:"expires_at"`
	LastAccessedAt time.Time `json:"last_accessed_at" db:"last_accessed_at"`
}

func NewUserRepo() userSessionRepo {

	return userSessionRepo{}
}

func (us *userSessionRepo) CreateSession() {

	query := userQuery.New(core.DB())

	query.CreateSession(context.Background(), userQuery.CreateSessionParams{
		ID:          us.ID,
		UserID:      us.UserID,
		Token:       us.Token,
		SessionData: us.SessionData,
		ExpiresAt:   us.CreatedAt,
	})

}
