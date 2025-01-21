package v1

import (
	"blackpay_gateway/utils"
	"encoding/json"
	"net/http"
)

// HomeHandler is a protected route that requires authentication.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract userID from the context
	userID := utils.GetFromContext(r.Context(), "userID")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the protected home page!",
		"userID":  userID.(string),
	})
}
