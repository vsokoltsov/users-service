package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vsokoltsov/users-service/app/api"
	"github.com/vsokoltsov/users-service/app/utils"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	dbCon := os.Getenv("DB_CON")
	utils.InitDB(dbCon)
	router := api.InitRouter(mux.NewRouter())
	router.Use(loggingMiddleware)
	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln(err)
	}
}
