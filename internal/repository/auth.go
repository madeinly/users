package repository

import (
	"context"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/queries/userQuery"
	"golang.org/x/crypto/bcrypt"
)

// Returns the user ID or error (no sql row is a possible error)
func (repo *sqliteRepo) ValidateCredentials(email string, password string) (userQuery.User, error) {
	ctx := context.Background()

	queries := userQuery.New(core.DB())

	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		return userQuery.User{}, err
	}

	// Compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return userQuery.User{}, err
	}
	return user, nil
}
