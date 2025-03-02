package routes

import (
	"net/http"

	"create-user-microservice/controllers"
	"create-user-microservice/middleware"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers the user-related endpoints
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")

	// Secure routes with JWT middleware
	router.Handle("/secure/users", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateUserHandler))).Methods("POST")
}
