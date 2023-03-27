package common

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Error while creating password hash")
	}
	return string(pw)
}
