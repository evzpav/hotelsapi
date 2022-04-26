package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Hotel struct {
	Name          string  `json:"name"`
	City          string  `json:"city"`
	NrOfEmployees int     `json:"nrOfEmployees"`
	Revenue       float64 `json:"revenue"`
	Active        bool    `json:"active"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hotels", getHotels).Methods(http.MethodGet)
	r.HandleFunc("/hotels", createHotels).Methods(http.MethodPost)
	http.ListenAndServe(":7070", r)
}

func getHotels(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	city := queryParams.Get("city")

	fmt.Println("city query param: ", city)

	// TODO db select

	hotels := []Hotel{
		{
			Name:          "Sheraton",
			City:          "São Paulo",
			NrOfEmployees: 300,
			Revenue:       456456.9,
			Active:        true,
		},
		{
			Name:          "Novotel",
			City:          "Florianópolis",
			NrOfEmployees: 244,
			Revenue:       345456.9,
			Active:        true,
		},
	}

	bs, err := json.Marshal(hotels)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func createHotels(w http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var hotel Hotel
	if err := json.Unmarshal(bs, &hotel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("Hotel: %+v\n", hotel)

	//TODO db insert

	w.WriteHeader(http.StatusOK)
	// w.Write(bs)
}
