package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	return string(hashedPasswd), nil
}

// CheckPassword compares if provided password is correct or not
func CheckPassword(password string, hashedPasswd string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPasswd), []byte(password), 
	)
}