package user

import (
	"context"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

type CreateUserParams struct {
	UserID   string
	Username string
	Email    string
	Password string
	Role     string
	Status   string
}

func Create(ctx context.Context, params CreateUserParams) error {

	appDB := core.DB()

	query := sqlc.New(appDB)

	_, err := query.CreateUser(ctx, sqlc.CreateUserParams{
		ID:       params.UserID,
		Username: params.Username,
		Email:    params.Email,
		Password: params.Password,
		Role:     params.Role,
		Status:   params.Status,
	})

	if err != nil {
		return err
	}

	return nil
}
