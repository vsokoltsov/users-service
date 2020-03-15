package main

import (
	"github.com/vsokoltsov/users-service/app/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := api.InitRouter(mux.NewRouter())
	router.Use(loggingMiddleware)
	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln(err)
	}
}
