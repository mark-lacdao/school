package routes

import (
	"net/http"
	"school/internal/handlers"

	"github.com/gorilla/mux"
)

// InitializeRoutes initializes the routes for the application
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	return r
}
