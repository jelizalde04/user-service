package controllers

import (
	"encoding/json"
	"net/http"

	"list-users-microservice/services"
)

// ListUsersHandler handles user listing requests
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
