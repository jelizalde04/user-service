package routes

import (
	"net/http"

	"delete-user-microservice/controllers"
	"delete-user-microservice/middleware"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers the user-related endpoints
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")

	// Secure route with JWT authentication
	router.Handle("/secure/users/{id}", middleware.AuthMiddleware(http.HandlerFunc(controllers.DeleteUserHandler))).Methods("DELETE")
}
