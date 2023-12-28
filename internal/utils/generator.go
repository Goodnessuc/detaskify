package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares the password with the hashed password
func ComparePassword(password, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func ResetToken(email string) string {
	namespace := uuid.NameSpaceDNS

	// Generate a version 5 UUID using the namespace and text
	id := uuid.NewSHA1(namespace, []byte(email))

	// Return the UUID as a string
	return id.String()
}
