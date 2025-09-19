package session

import (
	"context"

	core "github.com/madeinly/core/v1"
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

func CreateUserSession(ctx context.Context, params sqlc.CreateSessionParams) error {

	q := sqlc.New(core.DB())

	err := q.CreateSession(ctx, params)

	if err != nil {
		return err
	}

	return nil
}

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
