package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// GenerateCSRFToken generates a random CSRF token
func GenerateCSRFToken() string {
	b := make([]byte, 16) // 16 bytes = 128 bits
	rand.Read(b)
	return hex.EncodeToString(b)
}

// ValidateCSRFToken compares the provided CSRF token with the stored one
func ValidateCSRFToken(providedToken, storedToken string) error {
	if providedToken == "" {
		return errors.New("missing CSRF token")
	}
	if providedToken != storedToken {
		return errors.New("invalid CSRF token")
	}
	return nil
}
