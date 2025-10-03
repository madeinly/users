package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

func GetByID(ctx context.Context, userID string) (sqlc.User, error) {

	query := sqlc.New(core.DB())

	u, err := query.GetUserByID(ctx, userID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {

		return sqlc.User{}, err
	}

	return u, nil

}
