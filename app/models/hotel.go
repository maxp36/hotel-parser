package models

import (
	"strconv"
)

// Hotel is the base datamodel of a hotel
type HotelRaw struct {
	Name        string
	Description string
	CountryCode string
	City        string
	Address     string
	Latitude    float64
	Longitude   float64
	Rating      float64
	Images      []string
}

// HotelJSON is the base datamodel of a hotel from JSON data file
type HotelJSON struct {
	CountryCode string `json:"country_code"`
	EN          struct {
		Address     string `json:"address"`
		City        string `json:"city"`
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"en"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Rating    struct {
		Total float64 `json:"total"`
	} `json:"rating"`
	Images []struct {
		OrigURL string `json:"orig_url"`
	} `json:"images"`
}

func (h HotelJSON) ToHotelRaw() *HotelRaw {
	images := make([]string, 0)
	for _, image := range h.Images {
		images = append(images, image.OrigURL)
	}

	return &HotelRaw{
		Name:        h.EN.Name,
		Description: h.EN.Description,
		CountryCode: h.CountryCode,
		City:        h.EN.City,
		Address:     h.EN.Address,
		Latitude:    h.Latitude,
		Longitude:   h.Longitude,
		Rating:      h.Rating.Total / 2,
		Images:      images,
	}
}

// HotelCSV is the base datamodel of a hotel from CSV data file
type HotelCSV struct {
	Name        string `json:"hotel_name"`
	Description string `json:"overview"`
	CountryCode string `json:"countryisocode"`
	City        string `json:"city"`
	Address     string `json:"addressline1"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Rating      string `json:"star_rating"`
	Photo1      string `json:"photo1"`
	Photo2      string `json:"photo2"`
	Photo3      string `json:"photo3"`
	Photo4      string `json:"photo4"`
	Photo5      string `json:"photo5"`
}

func (h HotelCSV) ToHotelRaw() (hotel *HotelRaw, err error) {

	latitude, err := strconv.ParseFloat(h.Latitude, 64)
	if err != nil {
		return nil, err
	}

	longitude, err := strconv.ParseFloat(h.Longitude, 64)
	if err != nil {
		return nil, err
	}

	rating, err := strconv.ParseFloat(h.Rating, 64)
	if err != nil {
		return nil, err
	}

	images := []string{
		h.Photo1,
		h.Photo2,
		h.Photo3,
		h.Photo4,
		h.Photo5,
	}

	return &HotelRaw{
		Name:        h.Name,
		Description: h.Description,
		CountryCode: h.CountryCode,
		City:        h.City,
		Address:     h.Address,
		Latitude:    latitude,
		Longitude:   longitude,
		Rating:      rating,
		Images:      images,
	}, nil
}
