package main

import (
	"github.com/vsokoltsov/users-service/app/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := api.InitRouter(mux.NewRouter())
	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln(err)
	}
}
