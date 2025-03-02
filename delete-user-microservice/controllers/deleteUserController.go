package controllers

import (
	"encoding/json"
	"net/http"

	"delete-user-microservice/services"

	"github.com/gorilla/mux"
)

// DeleteUserHandler handles user deletion requests
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := services.DeleteUser(userID)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
