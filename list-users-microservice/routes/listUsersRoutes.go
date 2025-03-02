package routes

import (
	"net/http"

	"list-users-microservice/controllers"
	"list-users-microservice/middleware"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers the user-related endpoints
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.ListUsersHandler).Methods("GET")

	// Secure route with JWT authentication
	router.Handle("/secure/users", middleware.AuthMiddleware(http.HandlerFunc(controllers.ListUsersHandler))).Methods("GET")
}
