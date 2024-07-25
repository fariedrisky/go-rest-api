package routes

import (
	"github.com/fariedrisky/go-restful-mysql/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes sets up the application routes.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/auth/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/auth/login", controllers.Login).Methods("POST")

	// Protected routes
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")

	return router
}
