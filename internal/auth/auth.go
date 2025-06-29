package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/queries/userQuery"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {

	passwordByte, err := bcrypt.GenerateFromPassword(
		[]byte(p),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	p = string(passwordByte)

	return p, nil

}

// func ValidateToken(tokenString string) (*Claims, error) {

// 	claims := &Claims{}

// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return JwtSecret, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, err
// 	}

// 	return claims, nil
// }

func ValidateCredentials(email string, password string) (bool, string) {
	ctx := context.Background()

	// Initialize queries with your database connection
	queries := userQuery.New(core.DB()) // Assuming db.Connection is *sql.DB

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
