package main

import (
	"hotelsapi/internal/domain"
	"hotelsapi/internal/handler"
	"hotelsapi/internal/server"
	"hotelsapi/internal/storage"
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

func main() {

	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.New(db)

	hotelService := domain.NewHotelService(storage)

	handler := handler.New(hotelService)
	router := mux.NewRouter()

	s := server.New(router, handler)

	s.Routes()

	log.Fatal(http.ListenAndServe(":7070", s))
}
