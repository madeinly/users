package repository

import (
	"context"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/queries/userQuery"
)

func (repo *sqliteRepo) GetSessionByUserID(userID string) userQuery.UserSession {

	ctx := context.Background()

	q := userQuery.New(repo.db)

	session, err := q.GetSessionByUserID(ctx, userID)

	if err != nil {
		core.Log("error on getting session by user ID", err.Error())

		return userQuery.UserSession{}
	}

	return session

}

func (repo *sqliteRepo) CreateUserSession(ctx context.Context, params userQuery.CreateSessionParams) error {

	q := userQuery.New(repo.db)

	err := q.CreateSession(ctx, params)

	if err != nil {
		return err
	}

	return nil
}

func (repo *sqliteRepo) UpdateUserSession(userID string, token string, expiresAt string) error {

	ctx := context.Background()

	q := userQuery.New(repo.db)

	//uses userID to find the session of the user and then updates the token and the expiration
	_, err := q.UpdateSessionToken(ctx, userQuery.UpdateSessionTokenParams{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		return err
	}

	return nil

}
