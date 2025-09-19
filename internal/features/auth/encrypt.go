package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {

	passwordByte, err := bcrypt.GenerateFromPassword(
		[]byte(p),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return ""
	}

	p = string(passwordByte)

	return p

}
