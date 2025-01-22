package v1

import (
	"blackpay_gateway/config"
	"blackpay_gateway/utils"
	"encoding/json"
	"net/http"
	"time"
)

// LogoutHandler handles user logout by invalidating tokens.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Extract userID from the context
	userID := utils.GetFromContext(r.Context(), "userID")

	// Clear tokens in the database
	_, err := config.DB.Exec("UPDATE auth_user SET token = NULL, csrf = NULL WHERE id = ?", userID)
	if err != nil {
		http.Error(w, "Error logging out", http.StatusInternalServerError)
		return
	}

	// Clear the JWT cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})

	json.NewEncoder(w).Encode(map[string]string{"message": "Successfully logged out"})
}
