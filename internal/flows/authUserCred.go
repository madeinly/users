package flows

import (
	"context"
	"time"

	"github.com/google/uuid"
	core "github.com/madeinly/core/v1"
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

	errors := core.Validate()

	if params.Email == "" && params.Username == "" {
		errors.Add("authentication", "--agregar-status-code--", "email and username can't both be empty")
	}

	if params.Password == "" {
		errors.Add("authentication", "--agregar-status-code--", "password can't be empty")
	}

	if errors.HasErrors() {
		return "", errors
	}

	if params.Email != "" {
		errors.Validate(params.Email, user.EmailRules)
	}

	if params.Username != "" {
		errors.Validate(params.Username, user.UsernameRules)
	}

	errors.Validate(params.Password, user.PasswordRules)

	if errors.HasErrors() {
		return "", errors
	}

	//=================================

	//NOTE: maybe could not use a repo and instead add all in here

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

	if u.ID == "" {

		err = session.CreateUserSession(ctx, sqlc.CreateSessionParams{
			ID:          uuid.New().String(),
			UserID:      u.ID,
			Token:       sessionToken,
			SessionData: "[]",
			ExpiresAt:   expirationTime.Format("2006-01-02 15:04:05"),
		})

	} else {
		err = session.UpdateUserSession(u.ID, sessionToken, expirationTime.Format("2006-01-02 15:04:05"))
	}

	if err != nil {
		return "", err
	}

	return token, nil

}
