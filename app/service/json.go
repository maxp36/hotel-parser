package service

import (
	"bytes"
	"encoding/json"

	"github.com/maxp36/hotel-parser/app/models"
)

func (s *parser) ParseJSON(data []byte) error {

	var hotel models.HotelJSON

	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&hotel); err != nil {
		return err
	}

	if err := s.R.AddHotel(hotel.ToHotelRaw()); err != nil {
		return err
	}

	return nil
}
