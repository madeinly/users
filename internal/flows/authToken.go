package flows

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
	"github.com/madeinly/users/internal/features/auth"
)

func AuthenticateWithToken(jwtToken string) (bool, error) {

	claims, err := auth.ParseToken(jwtToken)

	if err != nil {
		return false, err
	}

	sessionToken := claims.SessionToken

	q := sqlc.New(core.DB())

	session, err := q.GetSessionByToken(context.Background(), sessionToken)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		core.Log(err.Error(), "could not fetch session from user")
		return false, ErrServerFailure
	}

	exp, err := time.ParseInLocation("2006-01-02 15:04:05", session.ExpiresAt, time.Local)

	if err != nil {
		core.Log(err.Error(), "could not parse the expiration of the session")
		return false, ErrServerFailure
	}

	now := time.Now()
	if now.After(exp) {
		return false, ErrSessionExpired
	}

	return true, nil

}
