package session

import (
	"context"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

func GetSessionByUserID(userID string) sqlc.UserSession {

	ctx := context.Background()

	q := sqlc.New(core.DB())

	session, err := q.GetSessionByUserID(ctx, userID)

	if err != nil {
		core.Log("error on getting session by user ID", err.Error())

		return sqlc.UserSession{}
	}

	return session

}
