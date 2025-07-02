package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/queries/userQuery"
	"golang.org/x/crypto/bcrypt"
)

func (repo *sqliteRepo) ValidateCredentials(email string, password string) (bool, string) {
	ctx := context.Background()

	queries := userQuery.New(core.DB())

	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, ""
		}
		log.Printf("Database error: %v", err)
		return false, ""
	}

	// Compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, ""
	}

	return true, user.ID
}
