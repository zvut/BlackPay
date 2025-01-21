package main

import (
	"blackpay_gateway/config"
	v1 "blackpay_gateway/routes/v1"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	config.ConnectDB()

	// Initialize router
	r := mux.NewRouter()
	v1.SetupV1Routes(r)

	// Start the server
	port := config.GetEnv("APP_PORT")
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
