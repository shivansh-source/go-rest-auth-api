package routes

import (
    "github.com/gorilla/mux"
    "github.com/shivansh-source/go-rest-auth-api/controllers"
)

func RegisterAuthRoutes(router *mux.Router) {
    router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
