package controllers

import (
	"encoding/json"
	"net/http"

	"create-user-microservice/models"
	"create-user-microservice/services"
	"create-user-microservice/utils"
)

// CreateUserHandler handles user creation requests
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Encrypt sensitive data
	user.Password = utils.Encrypt(user.Password)

	// Call user service
	createdUser, err := services.CreateUser(user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}
