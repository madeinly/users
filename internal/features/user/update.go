package user

import (
	"context"
	"fmt"

	core "github.com/madeinly/core/v1"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

type UpdateUserParams struct {
	ID       string
	Username string
	Email    string
	Status   string
	Password string
	Role     string
}

func Update(ctx context.Context, args UpdateUserParams) error {

	tx, err := core.DB().Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := sqlc.New(tx)

	if args.Username != "" {
		query.UpdateUserUsername(ctx, sqlc.UpdateUserUsernameParams{
			Username: args.Username,
			ID:       args.ID,
		})
	}

	if args.Email != "" {
		query.UpdateUserEmail(ctx, sqlc.UpdateUserEmailParams{
			Email: args.Email,
			ID:    args.ID,
		})
	}

	if args.Status != "" {
		query.UpdateUserStatus(ctx, sqlc.UpdateUserStatusParams{
			Status: args.Status,
			ID:     args.ID,
		})
	}

	if args.Password != "" {
		query.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
			Password: args.Password,
			ID:       args.ID,
		})
	}

	if args.Role != "" {
		query.UpdateUserRole(ctx, sqlc.UpdateUserRoleParams{
			Role: args.Role,
			ID:   args.ID,
		})
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
