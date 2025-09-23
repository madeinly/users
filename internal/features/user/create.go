package user

import (
	"context"
	"fmt"

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

func Create(ctx context.Context, param CreateUserParams) (string, error) {

	tx, err := core.DB().BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := sqlc.New(tx)

	if _, err := query.CreateUser(ctx, sqlc.CreateUserParams{
		ID:       param.UserID,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Role:     param.Role,
		Status:   param.Status,
	}); err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return param.UserID, nil
}
