package session

import (
	"context"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

func CreateUserSession(ctx context.Context, params sqlc.CreateSessionParams) error {

	q := sqlc.New(core.DB())

	err := q.CreateSession(ctx, params)

	if err != nil {
		return err
	}

	return nil
}
