package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/madeinly/core"
)

type ResToken struct {
	Token string `json:"token"`
}

var settings = core.Settings()
var jwtSecret = []byte(settings.JWTSalt) // Change this!

type Claims struct {
	UserID       string `json:"user_id"`
	SessionToken string `json:"session_token"`
	Role         string `json:"user_roleID"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, sessionToken, role string) (string, error) {

	claims := &Claims{
		UserID:           userID,
		SessionToken:     sessionToken,
		Role:             role,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
