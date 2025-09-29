package flows

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
	"github.com/madeinly/users/internal/features/auth"
	"github.com/madeinly/users/internal/features/session"
	"github.com/madeinly/users/internal/features/user"
)

type ValidateCredentialsParams struct {
	Email    string
	Username string
	Password string
}

func ValidateCredentials(ctx context.Context, params ValidateCredentialsParams) (string, error) {

	//========================= fast errors

	fastErrors := core.Validate()

	if params.Email == "" && params.Username == "" {
		fastErrors.Add("authentication", "--agregar-status-code--", "email and username can't both be empty")
	}

	if params.Password == "" {
		fastErrors.Add("authentication", "--agregar-status-code--", "password can't be empty")
	}

	if fastErrors.HasErrors() {
		return "", fastErrors
	}

	if params.Email != "" {
		fastErrors.Validate(params.Email, user.EmailRules)
	}

	if params.Username != "" {
		fastErrors.Validate(params.Username, user.UsernameRules)
	}

	fastErrors.Validate(params.Password, user.PasswordRules)

	if fastErrors.HasErrors() {
		return "", fastErrors
	}

	//=================================

	u, err := auth.ValidateCredentials(params.Email, params.Password)

	if err != nil {
		return "", ErrInvalidCredentials
	}

	sessionToken := uuid.New().String()

	expirationTime := time.Now().Add(2 * time.Hour)

	token, err := auth.GenerateToken(sessionToken, u.Role)

	if err != nil {
		return "", ErrServerFailure
	}

	err = session.UpdateUserSession(u.ID, sessionToken, expirationTime.Format("2006-01-02 15:04:05"))

	if err != nil && errors.Is(err, sql.ErrNoRows) {

		err = session.CreateUserSession(ctx, sqlc.CreateSessionParams{
			ID:          uuid.New().String(),
			UserID:      u.ID,
			Token:       sessionToken,
			SessionData: "[]",
			ExpiresAt:   expirationTime.Format("2006-01-02 15:04:05"),
		})

	}

	if err != nil {
		return "", ErrServerFailure
	}

	return token, nil

}
