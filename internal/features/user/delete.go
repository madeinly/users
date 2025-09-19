package user

import (
	"context"
	"fmt"

	core "github.com/madeinly/core/v1"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

func Delete(ctx context.Context, userID string) error {

	tx, err := core.DB().BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	q := sqlc.New(tx)

	err = q.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	err = q.DeleteSession(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	err = q.DeleteMetas(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete metas: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil

}
