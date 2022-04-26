package domain

type Hotel struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	City          string  `json:"city"`
	NrOfEmployees int     `json:"nrOfEmployees"`
	Revenue       float64 `json:"revenue"`
	Active        bool    `json:"active"`
}

type HotelFilter struct {
	City string
}

type HotelService interface {
	CreateHotel(*Hotel) error
	GetHotels(HotelFilter) ([]*Hotel, error)
}

type HotelStorage interface {
	InsertHotel(*Hotel) error
	GetHotels(HotelFilter) ([]*Hotel, error)
}

type hotelService struct {
	storage HotelStorage
}

func NewHotelService(storage HotelStorage) *hotelService {
	return &hotelService{
		storage: storage,
	}
}

func (hs *hotelService) CreateHotel(hotel *Hotel) error {
	return hs.storage.InsertHotel(hotel)
}

func (hs *hotelService) GetHotels(hf HotelFilter) ([]*Hotel, error) {
	return hs.storage.GetHotels(hf)
}
