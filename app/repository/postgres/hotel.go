package postgres

import "time"

import "github.com/maxp36/hotel-parser/app/models"

func (r *repository) AddHotel(hotel *models.HotelRaw) (id int64, err error) {
	ins := `insert into hotels
	(
		created_at, 
		updated_at, 
		name, 
		description,
		countryCode,
		city,
		address,
		latitude, 
		longitude, 
		rating
	) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	;`
	res, err := r.DB.Exec(ins,
		time.Now(),
		time.Now(),
		hotel.Name,
		hotel.Description,
		hotel.CountryCode,
		hotel.City,
		hotel.Address,
		hotel.Latitude,
		hotel.Longitude,
		hotel.Rating,
	)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
