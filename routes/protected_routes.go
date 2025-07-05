package routes

import (
	"github.com/gorilla/mux"
	"github.com/shivansh-source/go-rest-auth-api/controllers"
	"github.com/shivansh-source/go-rest-auth-api/middleware"
)

func RegisterProtectedRoutes(router *mux.Router) {
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/protected", controllers.ProtectedEndpoint).Methods("GET")
}
// This function registers the protected routes that require JWT authentication.
// It uses the JWTMiddleware to ensure that only authenticated users can access these routes.