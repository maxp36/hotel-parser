package postgres

func (r *repository) AddHotelImage(hotelID int64, origURL string) error {
	ins := `insert into hotel_images
	(
		hotel_id, 
		orig_url
	) 
	values ($1, $2) 
	;`
	_, err := r.DB.Exec(ins, hotelID, origURL)
	if err != nil {
		return err
	}

	return nil
}
