package api

import (
	"github.com/gorilla/mux"
	v1 "github.com/vsokoltsov/users-service/app/api/v1"
)

// InitRouter initializes base router applications
func InitRouter(router *mux.Router) *mux.Router {
	var app = router.PathPrefix("/api").Subrouter()
	v1.InitNamespace(app.PathPrefix("/v1").Subrouter())
	return router
}
