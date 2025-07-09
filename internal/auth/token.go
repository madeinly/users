package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ResToken struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiration"`
}

// Define your secret key (should be in environment variables in production)
var jwtSecret = []byte("7996a6f6308f3825593597349e53a30c") // Change this!

// Claims struct now embeds jwt.RegisteredClaims for standard JWT fields
type Claims struct {
	UserID       string `json:"user_id"`
	SessionToken string `json:"session_token"`
	Role         string `json:"user_roleID"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token with custom claims
func GenerateToken(userID, sessionToken, role string, expiresAt time.Time) (string, error) {
	// Set expiration time

	// Create the claims
	claims := &Claims{
		UserID:       userID,
		SessionToken: sessionToken,
		Role:         role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken validates the JWT token and returns the claims
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
