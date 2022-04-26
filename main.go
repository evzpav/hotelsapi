package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	h := NewHandler()
	r := mux.NewRouter()
	r.HandleFunc("/hotels", h.getHotels).Methods(http.MethodGet)
	r.HandleFunc("/hotels", createHotels).Methods(http.MethodPost)
	// r.HandleFunc("/hotels/{id:[0-9]+}", updateHotel).Methods(http.MethodPut)
	// r.HandleFunc("/hotels/{id:[0-9]+}", deleteHotel).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)
}

type handler struct {}

func NewHandler()*handler {
	return &handler{}
}

func (h *handler) getHotels(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	city := queryParams.Get("city")

	fmt.Println(city)

}

func createHotels(w http.ResponseWriter, r *http.Request) {

}

// func updateHotel(w http.ResponseWriter, r *http.Request) {

// }

// func deleteHotel(w http.ResponseWriter, r *http.Request) {

// }
