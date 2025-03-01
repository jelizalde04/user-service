package routes

import (
    "net/http"
    "user-microservice/controllers"

    "github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, controller *controllers.UserController) {
    router.HandleFunc("/users", controller.CreateUserHandler).Methods("POST")
}
