package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/madeinly/core"
)

type ResToken struct {
	Token string `json:"token"`
}

var settings = core.Settings()
var jwtSecret = []byte(settings.JWTSalt) // Change this!

type Claims struct {
	SessionToken string `json:"session_token"`
	jwt.RegisteredClaims
}

func GenerateToken(sessionToken, role string) (string, error) {

	claims := &Claims{
		SessionToken:     sessionToken,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println(err.Error())
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
