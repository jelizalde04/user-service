package services

import (
	"errors"
	"log"

	"list-users-microservice/models"
)

// Mock user database
var users = []models.User{
	{ID: "1", Name: "John Doe", Email: "john@example.com"},
	{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
}

// GetAllUsers returns a list of all users
func GetAllUsers() ([]models.User, error) {
	if len(users) == 0 {
		return nil, errors.New("No users found")
	}
	log.Println("Users retrieved")
	return users, nil
}
