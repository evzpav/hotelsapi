package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hotels", getHotels).Methods(http.MethodGet)
	r.HandleFunc("/hotels", createHotels).Methods(http.MethodPost)
	http.ListenAndServe(":7070", r)
}

func getHotels(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	city := queryParams.Get("city")

	fmt.Println(city)

}

func createHotels(w http.ResponseWriter, r *http.Request) {

}
