package service

import "github.com/maxp36/hotel-parser/app"

import "encoding/json"

import "bytes"

import "github.com/maxp36/hotel-parser/app/models"

import "log"

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

	var hotel models.HotelRaw

	dec := json.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&hotel)
	if err != nil {
		return err
	}

	log.Printf("%#v", hotel)

	// id, err := s.R.AddHotel(&hotel)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *parser) ParseCSV(data []byte) error {
	return nil
}

func (s *parser) ParseXML(data []byte) error {
	return nil
}
