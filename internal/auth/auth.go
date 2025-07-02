package auth

import (
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
