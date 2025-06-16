package repo

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/madeinly/users/internal/queries/userQuery"

	"github.com/madeinly/core"

	"golang.org/x/crypto/bcrypt"
)

func Auth(email string, password string) (bool, string) {
	ctx := context.Background()

	queries := userQuery.New(core.DB())

	// Get user by email using sqlc generated query
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
