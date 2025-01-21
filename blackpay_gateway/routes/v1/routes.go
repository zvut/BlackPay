package v1

import (
	v1 "blackpay_gateway/controllers/v1"
	"blackpay_gateway/middleware"

	"github.com/gorilla/mux"
)

func SetupV1Routes(r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()

	// Public routes
	api.HandleFunc("/register", v1.Register).Methods("POST")
	api.HandleFunc("/login", v1.Login).Methods("POST")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.Authenticate)
	protected.HandleFunc("/home", v1.HomeHandler).Methods("GET")
	protected.HandleFunc("/logout", v1.LogoutHandler).Methods("POST")
}
