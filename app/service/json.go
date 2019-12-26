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

	dbHotel, ok := hotel.ToHotelRaw()
	if !ok {
		return nil
	}

	if err := s.R.AddHotel(dbHotel); err != nil {
		return err
	}

	return nil
}
