package util

import (
	"github.com/bnsngltn/go-fiber-boilerplate/config"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword func
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.PasswordHashingCost)

	return string(bytes), err
}

func VerifyPassword(passwordHash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	return err
}
