package postgres

import "time"

func (r *repository) AddHotel(name, description, countryCode, city, address string, latitude, longitude, rating float64) (id int64, err error) {
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
	res, err := r.DB.Exec(ins, time.Now(), time.Now(), name, description, countryCode, city, address, latitude, longitude, rating)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
