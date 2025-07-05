package routes

import (
	"github.com/gorilla/mux"
	"github.com/shivansh-source/go-rest-auth-api/controllers"
	"github.com/shivansh-source/go-rest-auth-api/middleware"
)

func RegisterPostRoutes(router *mux.Router) {
	posts := router.PathPrefix("/api/posts").Subrouter()
	posts.Use(middleware.JWTMiddleware)

	posts.HandleFunc("", controllers.CreatePost).Methods("POST")
	posts.HandleFunc("", controllers.GetAllPosts).Methods("GET")
	posts.HandleFunc("/{id}", controllers.UpdatePost).Methods("PUT")
	posts.HandleFunc("/{id}", controllers.DeletePost).Methods("DELETE")
}
