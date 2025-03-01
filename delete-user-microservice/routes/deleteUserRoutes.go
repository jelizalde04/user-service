package routes

import (
    "user-microservice/controllers"

    "github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, controller *controllers.UserController) {
    router.HandleFunc("/users/{id}", controller.DeleteUserHandler).Methods("DELETE")
}
