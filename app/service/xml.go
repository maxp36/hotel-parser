package service

import (
	"encoding/xml"
	"io"

	"github.com/maxp36/hotel-parser/app/models"
)

func (s *parser) ParseXML(r io.Reader) error {

	dec := xml.NewDecoder(r)

	for {
		token, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch el := token.(type) {
		case xml.StartElement:

			if el.Name.Local == models.HotelXMLName {

				var hotel models.HotelXML

				if err := dec.DecodeElement(&hotel, &el); err != nil {
					return err
				}

				if err := s.R.AddHotel(hotel.ToHotelRaw()); err != nil {
					return err
				}
			}

		}

	}

	return nil
}
