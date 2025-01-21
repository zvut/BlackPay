package v1

import (
	"blackpay_gateway/config"
	"blackpay_gateway/models"
	"blackpay_gateway/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Helper function for sending JSON responses
func jsonResponse(w http.ResponseWriter, statusCode int, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Register handles user registration.
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Hash the password
	hashedPassword := utils.HashPassword(user.Password)

	//GENERATE UUID
	user_uuid := utils.GenerateUUID()

	// Save the user to the database
	_, err := config.DB.Exec("INSERT INTO auth_user (id, mobile, password) VALUES (?, ?, ?)", user_uuid, user.Mobile, hashedPassword)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error saving user"})
		return
	}

	jsonResponse(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

// Login handles user authentication and token generation.
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Fetch user credentials from the database
	row := config.DB.QueryRow("SELECT id, password FROM auth_user WHERE mobile = ?", user.Mobile)
	var dbPassword string
	if err := row.Scan(&user.ID, &dbPassword); err != nil {
		fmt.Println(err)
		jsonResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		return
	}

	// Verify the password
	if !utils.CheckPasswordHash(user.Password, dbPassword) {
		jsonResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		return
	}

	// Generate JWT and CSRF tokens
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}
	csrf := utils.GenerateCSRFToken()

	// Store tokens in the database
	_, err = config.DB.Exec("UPDATE auth_user SET token = ?, csrf = ? WHERE id = ?", token, csrf, user.ID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error updating user tokens"})
		return
	}

	// Set the JWT token as an HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	// Respond with the CSRF token
	jsonResponse(w, http.StatusOK, map[string]string{
		"csrf": csrf,
	})
}
