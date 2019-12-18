package service

import (
	"bytes"
	"encoding/json"

	"github.com/maxp36/hotel-parser/app"
	"github.com/maxp36/hotel-parser/app/models"
)

type parser struct {
	R app.Repository
}

// NewParser creates new service objects that implements Service interface
func NewParser(r app.Repository) app.Parser {
	return &parser{
		R: r,
	}
}

func (s *parser) ParseJSON(data []byte) error {

	var hotel models.HotelJSON

	dec := json.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&hotel)
	if err != nil {
		return err
	}

	err = s.R.AddHotel(hotel.ToHotelRaw())
	if err != nil {
		return err
	}

	return nil
}

func (s *parser) ParseCSV(columns, data []string) error {

	var hotel models.HotelCSV

	m := make(map[string]string)
	for i, c := range columns {
		m[c] = data[i]
	}

	mar, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(mar, &hotel)
	if err != nil {
		return err
	}

	dbHotel, err := hotel.ToHotelRaw()
	if err != nil {
		return err
	}

	err = s.R.AddHotel(dbHotel)
	if err != nil {
		return err
	}

	return nil
}

func (s *parser) ParseXML(data []byte) error {
	return nil
}
