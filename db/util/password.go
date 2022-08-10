package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//HashedPassword return bcrypt password hash
func HashPassword(password string) (string, error) {
	hashedPasword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}

	return string(hashedPasword), nil
}

//CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(password))
}
