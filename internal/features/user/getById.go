package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

type User struct {
	ID                string
	Role              string
	Username          string
	Email             string
	Password          string
	Status            string
	PasswordUpdatedAt string
	CreatedAt         string
	UpdatedAt         string
	UserStatus        string
	LastLogin         string
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
		UserStatus:        u.Status,
		LastLogin:         s.LastAccessedAt,
	}, nil

}
