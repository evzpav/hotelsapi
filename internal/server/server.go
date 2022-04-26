package server

import (
	"hotelsapi/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	handler *handler.Handler
}

func New(r *mux.Router, handler *handler.Handler) Server {
	return Server{
		router:  r,
		handler: handler,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Routes() {
	s.router.HandleFunc("/hotels", s.handler.GetHotels).Methods(http.MethodGet)
	s.router.HandleFunc("/hotels", s.handler.CreateHotels).Methods(http.MethodPost)
}
