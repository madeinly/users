package auth

import (
	"context"

	core "github.com/madeinly/core/v1"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
	"golang.org/x/crypto/bcrypt"
)

// Returns the user or error (no sql row is a possible error)
func ValidateCredentials(email string, password string) (sqlc.User, error) {
	ctx := context.Background()

	queries := sqlc.New(core.DB())

	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		return sqlc.User{}, err
	}

	// Compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}
