package handler

import (
	"encoding/json"
	"fmt"
	"hotelsapi/internal/domain"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	hotelService domain.HotelService
}

func New(hotelService domain.HotelService) *Handler {
	return &Handler{
		hotelService: hotelService,
	}
}

func (h *Handler) GetHotels(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	city := queryParams.Get("city")

	fmt.Printf("City query param: %s\n", city)

	hotels, err := h.hotelService.GetHotels(domain.HotelFilter{City: city})
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

	var hotel *domain.Hotel
	if err := json.Unmarshal(bs, hotel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("Hotel: %+v\n", hotel)

	err = h.hotelService.CreateHotel(hotel)
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
