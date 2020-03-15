package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buffer bytes.Buffer
	buffer.WriteString(`{Response: "success", Message: "users service"}`)
	json.NewEncoder(w).Encode(buffer.String())
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln(err)
	}
}
