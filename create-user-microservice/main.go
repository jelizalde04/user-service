package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"create-user-microservice/controllers"
	"create-user-microservice/middleware"
	"create-user-microservice/routes"
	"create-user-microservice/utils"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the router
	r := mux.NewRouter()

	// Enable CORS
	r.Use(middleware.CORSHandler)

	// User routes
	routes.RegisterUserRoutes(r)

	// Swagger documentation route
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger/"))))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "2001"
	}

	fmt.Println("Create User Microservice is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
