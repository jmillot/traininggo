package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/application", getApplications).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getApplications(w http.ResponseWriter, r *http.Request) {

}
