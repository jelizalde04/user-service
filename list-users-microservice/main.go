package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "user-microservice/controllers"
    "user-microservice/routes"
    "user-microservice/services"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load()
    InitDB()

    userService := services.NewUserService(DB)
    userController := controllers.NewUserController(userService)

    router := mux.NewRouter()
    routes.RegisterUserRoutes(router, userController)

    port := os.Getenv("PORT")
    fmt.Println("Servidor corriendo en el puerto:", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
