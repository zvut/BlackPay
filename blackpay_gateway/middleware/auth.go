package middleware

import (
	"blackpay_gateway/config"
	"blackpay_gateway/utils"

	"encoding/json"
	"net/http"
)

// jsonResponse sends a standardized JSON response
func jsonResponse(w http.ResponseWriter, statusCode int, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Authenticate verifies the JWT and CSRF tokens for protected routes.
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract JWT token from the HttpOnly cookie
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			jsonResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized: Missing or invalid token"})
			return
		}

		// Validate the JWT token
		userID, err := utils.ValidateJWT(cookie.Value)
		if err != nil {
			jsonResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized: Invalid token"})
			return
		}

		// Extract CSRF token from headers
		csrfToken := r.Header.Get("X-CSRF-Token")
		if csrfToken == "" {
			jsonResponse(w, http.StatusForbidden, map[string]string{"error": "Forbidden: Missing CSRF token"})
			return
		}

		// Retrieve stored CSRF token from database
		var storedCSRFToken string
		row := config.DB.QueryRow("SELECT csrf FROM auth_user WHERE id = ?", userID)
		if err := row.Scan(&storedCSRFToken); err != nil {
			jsonResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized: User not found"})
			return
		}

		// Validate CSRF token
		if err := utils.ValidateCSRFToken(csrfToken, storedCSRFToken); err != nil {
			jsonResponse(w, http.StatusForbidden, map[string]string{"error": "Forbidden: Invalid CSRF token"})
			return
		}

		// Add userID to the context
		r = r.WithContext(utils.AddToContext(r.Context(), "userID", userID))

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
