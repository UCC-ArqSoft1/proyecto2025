package services

import "backend/domain"

func GetHotelByID(id int) domain.Hotel {
	location := domain.Location{
		Country: "Argentina",
		City:    "Cordoba",
		Street:  "Av. Colon",
		Number:  100,
		ZipCode: 5000,
	}

	hotel := domain.Hotel{
		ID:         id,
		Name:       "SHERATON 2025",
		RoomsCount: 20,
		Rating:     5.00,
		Location:   location,
	}
	return hotel
}
