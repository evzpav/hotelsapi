package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StorageMock struct {
	InsertHotelFn            func(hotel *Hotel) error
	InsertHotelsInvokedCount int
	GetHotelsFn              func(hf HotelFilter) ([]*Hotel, error)
	GetHotelsInvokedCount    int
}

func (sm *StorageMock) InsertHotel(h *Hotel) error {
	sm.InsertHotelsInvokedCount++
	return sm.InsertHotelFn(h)
}

func (sm *StorageMock) GetHotels(hf HotelFilter) ([]*Hotel, error) {
	sm.GetHotelsInvokedCount++
	return sm.GetHotelsFn(hf)
}

var storageMock = &StorageMock{
	InsertHotelFn: func(hotel *Hotel) error {
		hotel.ID = 2
		return nil
	},
	GetHotelsFn: func(hf HotelFilter) ([]*Hotel, error) {
		return []*Hotel{
			{
				ID:            1,
				Name:          "Costão do Santinho",
				City:          "Florianópolis",
				NrOfEmployees: 100,
				Revenue:       2342343.99,
				Active:        true,
			},
		}, nil
	},
}

func Test_CreateHotel(t *testing.T) {

	service := NewHotelService(storageMock)

	hotel := &Hotel{
		Name:          "Novohotel",
		City:          "Florianópolis",
		NrOfEmployees: 120,
		Revenue:       12121.21,
		Active:        true,
	}

	err := service.CreateHotel(hotel)
	assert.Nil(t, err)
	assert.Equal(t, 2, hotel.ID)
	assert.Equal(t, 1, storageMock.InsertHotelsInvokedCount)
	assert.Equal(t, 0, storageMock.GetHotelsInvokedCount)
}

func Test_GetHotels(t *testing.T) {
	service := NewHotelService(storageMock)

	hotels, err := service.GetHotels(HotelFilter{City: "Florianópolis"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(hotels))
	assert.Equal(t, "Costão do Santinho", hotels[0].Name)
	assert.Equal(t, 0, storageMock.InsertHotelsInvokedCount)
	assert.Equal(t, 1, storageMock.GetHotelsInvokedCount)

}
