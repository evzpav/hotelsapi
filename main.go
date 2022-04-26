package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresURL string = os.Getenv("POSTGRES_URL")
)

type Hotel struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	City          string  `json:"city"`
	NrOfEmployees int     `json:"nrOfEmployees"`
	Revenue       float64 `json:"revenue"`
	Active        bool    `json:"active"`
}

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func main() {
	r := mux.NewRouter()
	fmt.Println(postgresURL)
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	h := NewHandler(db)

	r.HandleFunc("/hotels", h.GetHotels).Methods(http.MethodGet)
	r.HandleFunc("/hotels", h.CreateHotels).Methods(http.MethodPost)

	fmt.Printf("Server running on http://localhost:7070\n")

	log.Fatal(http.ListenAndServe(":7070", r))
}

func (h *Handler) GetHotels(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	city := queryParams.Get("city")

	fmt.Println(city)

	var hotels []*Hotel
	err := h.db.Table("hotels").
		Find(&hotels).
		Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
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

func (h *Handler) CreateHotels(w http.ResponseWriter, r *http.Request) {
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

	err = h.db.Table("hotels").Create(&hotel).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	respBS, err := json.Marshal(hotel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBS)
}
