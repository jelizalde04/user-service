package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "user-microservice/services"

    "github.com/gorilla/mux"
)

type UserController struct {
    Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

func (c *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteUser(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Usuario eliminado exitosamente",
    })
}
