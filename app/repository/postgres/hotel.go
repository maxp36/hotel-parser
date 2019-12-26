package postgres

import (
	"time"

	"github.com/maxp36/hotel-parser/app/models"
)

func (r *repository) AddHotel(hotel *models.HotelRaw) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	var id int64

	{
		ins := `insert into hotels
		(
			created_at, 
			updated_at, 
			name, 
			description,
			country_code,
			city,
			address,
			latitude, 
			longitude, 
			rating
		) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		on conflict on constraint hotels_name_country_code_city_unique do update
		set
			updated_at = $2,
			description = coalesce( $4, hotels.description ),
			address = coalesce( $7, hotels.address ),
			latitude = coalesce( $8, hotels.latitude ),
			longitude = coalesce( $9, hotels.longitude ),
			rating = coalesce( $10, hotels.rating )
		returning id
		;`

		err = tx.Get(&id, ins,
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
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
	}

	{
		ins := `insert into hotel_images
		(
			hotel_id,
			orig_url
		) 
		values ($1, $2) 
		on conflict on constraint hotel_images_hotel_id_orig_url_unique do nothing
		;`
		for _, image := range hotel.Images {

			_, err := tx.Exec(ins,
				id,
				image,
			)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
