package v1

import (
	"github.com/gorilla/mux"
)

// InitNamespace defines /v1 REST API routes
// for the given router
func InitNamespace(subrouter *mux.Router) {
	subrouter.HandleFunc("/users", getUsers).Methods("GET")
}
