package app

import "github.com/maxp36/hotel-parser/app/models"

// Repository represent the parser's Postgres repository contract.
type Repository interface {
	AddHotel(hotel *models.HotelRaw) error
}
