package services

import (
	"errors"
	"log"

	"delete-user-microservice/models"
)

// Mock user database
var users = []models.User{
	{ID: "1", Name: "John Doe", Email: "john@example.com"},
	{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
}

// DeleteUser deletes a user from the mock database
func DeleteUser(userID string) error {
	for i, user := range users {
		if user.ID == userID {
			// Remove user from slice
			users = append(users[:i], users[i+1:]...)
			log.Println("User deleted:", userID)
			return nil
		}
	}
	return errors.New("User not found")
}
