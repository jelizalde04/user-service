package services

import (
    "database/sql"
    "errors"
    "user-microservice/models"
)

type UserService struct {
    DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{DB: db}
}

func (s *UserService) CreateUser(user models.User) (int, error) {
    var userID int
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    err := s.DB.QueryRow(query, user.Name, user.Email).Scan(&userID)
    if err != nil {
        return 0, errors.New("error al crear usuario")
    }
    return userID, nil
}
