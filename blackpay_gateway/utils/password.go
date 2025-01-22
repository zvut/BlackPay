package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed version of the given password using bcrypt.
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	return string(hashedPassword)
}

// CheckPasswordHash compares a hashed password with its plaintext version.
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
