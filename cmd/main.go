package main

import (
	"fmt"
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
	port        string = os.Getenv("PORT")
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

	fmt.Printf("Server running on http://localhost:%s\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, s))

}