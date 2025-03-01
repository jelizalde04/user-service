package controllers

import (
    "encoding/json"
    "net/http"
    "user-microservice/models"
    "user-microservice/services"
)

type UserController struct {
    Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Solicitud inválida", http.StatusBadRequest)
        return
    }

    userID, err := c.Service.CreateUser(user)
    if err != nil {
        http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Usuario creado exitosamente",
        "user_id": userID,
    })
}
