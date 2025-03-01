package controllers

import (
    "encoding/json"
    "net/http"
    "user-microservice/services"
)

type UserController struct {
    Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

func (c *UserController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    users, err := c.Service.GetAllUsers()
    if err != nil {
        http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}
