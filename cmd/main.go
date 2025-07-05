package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/shivansh-source/go-rest-auth-api/config"
    "github.com/shivansh-source/go-rest-auth-api/routes"
)

func main() {
    config.InitDB()

    router := mux.NewRouter()
    routes.RegisterAuthRoutes(router)
    routes.RegisterProtectedRoutes(router)
    routes.RegisterPostRoutes(router)

    log.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", router)
    // Handle any errors that occur during server startup
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }     
}
