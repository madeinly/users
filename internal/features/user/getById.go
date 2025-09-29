package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

type User struct {
	ID                string `json:"user_id"`
	Role              string `json:"user_role"`
	Username          string `json:"user_username"`
	Email             string `json:"user_email"`
	Password          string `json:"-"`
	Status            string `json:"user_status"`
	PasswordUpdatedAt string `json:"-"`
	CreatedAt         string `json:"user_createdAt"`
	UpdatedAt         string `json:"user_updatedAt"`
	LastLogin         string `json:"user_lastLoginAT"`
}

func GetByID(ctx context.Context, userID string) (User, error) {

	query := sqlc.New(core.DB())

	u, err := query.GetUserByID(ctx, userID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {

		return User{}, err
	}

	s, err := query.GetSessionByUserID(ctx, userID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return User{}, err
	}

	return User{
		ID:                u.ID,
		Role:              u.Role,
		Username:          u.Username,
		Email:             u.Email,
		Password:          u.Password,
		Status:            u.Status,
		PasswordUpdatedAt: u.PasswordUpdatedAt,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
		LastLogin:         s.LastAccessedAt,
	}, nil

}
