package storage

import (
	"hotelsapi/internal/domain"

	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) InsertHotel(hotel *domain.Hotel) error {
	return s.db.Table("hotels").Create(&hotel).Error
}

func (s *Storage) GetHotels(hf domain.HotelFilter) ([]*domain.Hotel, error) {
	var hotels []*domain.Hotel
	tx := s.db.Table("hotels")

	if hf.City != "" {
		tx.Where("city = ?", hf.City)
	}
	
	err := tx.
		Find(&hotels).
		Error
	return hotels, err

}
