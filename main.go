package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Elagoht/go-pass/controllers"
	"github.com/Elagoht/go-pass/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Warning: .env file not found")
	}

	// Initialize the database, If there is no database, it will be created
	db.InitDB()

	// Create a new router from the gorilla/mux package
	router := mux.NewRouter()

	// Initialize the account controller
	accountController := controllers.NewAccountController()

	// Define the endpoints for the API
	router.HandleFunc("/accounts", accountController.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts", accountController.GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", accountController.GetAccount).Methods("GET")
	router.HandleFunc("/accounts/{id}", accountController.UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", accountController.DeleteAccount).Methods("DELETE")

	// Get server configuration from environment variables
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	// Start the server
	serverAddr := host + ":" + port
	log.Printf("Server is running on http://%s", serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatal("Server could not be started:", err)
	}
}
