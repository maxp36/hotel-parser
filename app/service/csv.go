package service

import (
	"encoding/json"

	"github.com/maxp36/hotel-parser/app/models"
)

func (s *parser) ParseCSV(columns, data []string) error {

	var hotel models.HotelCSV

	m := make(map[string]string)
	for i, c := range columns {
		m[c] = data[i]
	}

	// maybe, this's not the best solution
	mar, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(mar, &hotel); err != nil {
		return err
	}
	// maybe, this's not the best solution

	dbHotel, err := hotel.ToHotelRaw()
	if err != nil {
		return err
	}

	if err := s.R.AddHotel(dbHotel); err != nil {
		return err
	}

	return nil
}
