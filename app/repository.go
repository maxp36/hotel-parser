package app

// Repository represent the parser's Postgres repository contract.
type Repository interface {
	AddHotel(name, description, countryCode, city, address string, latitude, longitude, rating float64) (id int64, err error)
	AddHotelImage(hotelID int64, origURL string) error
}
