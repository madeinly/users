package session

import (
	"context"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

func UpdateUserSession(userID string, token string, expiresAt string) error {

	ctx := context.Background()

	q := sqlc.New(core.DB())

	//uses userID to find the session of the user and then updates the token and the expiration
	_, err := q.UpdateSessionToken(ctx, sqlc.UpdateSessionTokenParams{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		return err
	}

	return nil

}
