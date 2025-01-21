package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTSecret is used for signing the JWT token
var JWTSecret = []byte(os.Getenv("JWT_SECRET")) // Replace with value from `.env`

// GenerateJWT generates a new JWT token with the given user ID
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(), // Token expiration (1 day)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateJWT validates the given JWT token and extracts the user ID if valid
func ValidateJWT(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token uses the correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	// Check if there was an error or if the token is invalid
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract claims and retrieve the user ID
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userID, ok := claims["userID"].(string); ok {
			return userID, nil
		}
		return "", errors.New("user ID not found in token")
	}

	return "", errors.New("invalid token claims")
}
