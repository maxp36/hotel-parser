package models

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

// Hotel is the base datamodel of a hotel
type HotelRaw struct {
	Name        string
	Description *string
	CountryCode string
	City        string
	Address     *string
	Latitude    *float64
	Longitude   *float64
	Rating      *float64
	Images      []string
}

// HotelJSON is the base datamodel of a hotel from JSON data file
type HotelJSON struct {
	CountryCode *string `json:"country_code,omitempty"`
	EN          *struct {
		Address     *string `json:"address,omitempty"`
		City        *string `json:"city,omitempty"`
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	} `json:"en,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Rating    *struct {
		Total *float64 `json:"total,omitempty"`
	} `json:"rating,omitempty"`
	Images []struct {
		OrigURL *string `json:"orig_url,omitempty"`
	} `json:"images,omitempty"`
}

// If Name, CountryCode and City are empty ok == false will be returned
func (h *HotelJSON) ToHotelRaw() (hotel *HotelRaw, ok bool) {

	hotel = new(HotelRaw)

	if h.CountryCode != nil && *h.CountryCode != "" {
		hotel.CountryCode = *h.CountryCode
	} else {
		return nil, false
	}

	if h.EN != nil {
		if h.EN.Address != nil && *h.EN.Address != "" {
			hotel.Address = h.EN.Address
		}

		if h.EN.City != nil && *h.EN.City != "" {
			hotel.City = *h.EN.City
		} else {
			return nil, false
		}

		if h.EN.Description != nil && *h.EN.Description != "" {
			description := strip.StripTags(*h.EN.Description)
			description = strings.Replace(description, "\n", " ", -1)
			description = strings.TrimSpace(description)
			hotel.Description = &description
		}

		if h.EN.Name != nil && *h.EN.Name != "" {
			hotel.Name = *h.EN.Name
		} else {
			return nil, false
		}
	}

	if h.Latitude != nil && validateLatitude(*h.Latitude) {
		hotel.Latitude = h.Latitude
	}

	if h.Longitude != nil && validateLongitude(*h.Longitude) {
		hotel.Longitude = h.Longitude
	}

	if h.Rating != nil && h.Rating.Total != nil {
		rating := math.Round(*h.Rating.Total / 2)
		if validateRating(rating) {
			hotel.Rating = &rating
		}
	}

	images := make([]string, 0)
	for _, image := range h.Images {
		if image.OrigURL != nil && *image.OrigURL != "" {
			images = append(images, *image.OrigURL)
		}
	}
	hotel.Images = images

	return hotel, true
}

// HotelCSV is the base datamodel of a hotel from CSV data file
type HotelCSV struct {
	Name        *string `json:"hotel_name,omitempty"`
	Description *string `json:"overview,omitempty"`
	CountryCode *string `json:"countryisocode,omitempty"`
	City        *string `json:"city,omitempty"`
	Address     *string `json:"addressline1,omitempty"`
	Latitude    *string `json:"latitude,omitempty"`
	Longitude   *string `json:"longitude,omitempty"`
	Rating      *string `json:"star_rating,omitempty"`
	Photo1      *string `json:"photo1,omitempty"`
	Photo2      *string `json:"photo2,omitempty"`
	Photo3      *string `json:"photo3,omitempty"`
	Photo4      *string `json:"photo4,omitempty"`
	Photo5      *string `json:"photo5,omitempty"`
}

// If Name, CountryCode and City are empty ok == false will be returned
func (h *HotelCSV) ToHotelRaw() (hotel *HotelRaw, ok bool, err error) {

	hotel = new(HotelRaw)

	if h.Name != nil && *h.Name != "" {
		hotel.Name = *h.Name
	} else {
		return nil, false, nil
	}

	if h.Description != nil && *h.Description != "" {
		hotel.Description = h.Description
	}

	if h.CountryCode != nil && *h.CountryCode != "" {
		hotel.CountryCode = *h.CountryCode
	} else {
		return nil, false, nil
	}

	if h.City != nil && *h.City != "" {
		hotel.City = *h.City
	} else {
		return nil, false, nil
	}

	if h.Address != nil && *h.Address != "" {
		hotel.Address = h.Address
	}

	if h.Latitude != nil && *h.Latitude != "" {
		latitude, err := strconv.ParseFloat(*h.Latitude, 64)
		if err != nil {
			return nil, false, err
		}

		if validateLatitude(latitude) {
			hotel.Latitude = &latitude
		}
	}

	if h.Longitude != nil && *h.Longitude != "" {
		longitude, err := strconv.ParseFloat(*h.Longitude, 64)
		if err != nil {
			return nil, false, err
		}

		if validateLongitude(longitude) {
			hotel.Longitude = &longitude
		}
	}

	if h.Rating != nil && *h.Rating != "" {
		rating, err := strconv.ParseFloat(*h.Rating, 64)
		if err != nil {
			return nil, false, err
		}

		if validateRating(rating) {
			hotel.Rating = &rating
		}
	}

	images := make([]string, 0)
	if h.Photo1 != nil && *h.Photo1 != "" {
		images = append(images, *h.Photo1)
	}
	if h.Photo2 != nil && *h.Photo2 != "" {
		images = append(images, *h.Photo2)
	}
	if h.Photo3 != nil && *h.Photo3 != "" {
		images = append(images, *h.Photo3)
	}
	if h.Photo4 != nil && *h.Photo4 != "" {
		images = append(images, *h.Photo4)
	}
	if h.Photo5 != nil && *h.Photo5 != "" {
		images = append(images, *h.Photo5)
	}
	hotel.Images = images

	return hotel, true, nil
}

var HotelXMLName = "hotel"

// HotelXML is the base datamodel of a hotel from XML data file
type HotelXML struct {
	Name        *string  `xml:"name,omitempty"`
	Description *string  `xml:"descriptions>en,omitempty"`
	CountryCode *string  `xml:"countrytwocharcode,omitempty"`
	City        *string  `xml:"city>en,omitempty"`
	Address     *string  `xml:"address,omitempty"`
	Latitude    *float64 `xml:"latitude,omitempty"`
	Longitude   *float64 `xml:"longitude,omitempty"`
	Rating      *float64 `xml:"stars,omitempty"`
	Images      []struct {
		URL *string `xml:"url,omitempty"`
	} `xml:"photos>photo,omitempty"`
}

// If Name, CountryCode and City are empty ok == false will be returned
func (h *HotelXML) ToHotelRaw() (hotel *HotelRaw, ok bool) {

	hotel = new(HotelRaw)

	if h.Name != nil && *h.Name != "" {
		hotel.Name = *h.Name
	} else {
		return nil, false
	}

	if h.Description != nil && *h.Description != "" {
		r := regexp.MustCompile(`\[.*?\]`)

		description := r.ReplaceAllString(*h.Description, "")
		description = strings.TrimSpace(description)

		hotel.Description = &description
	}

	if h.CountryCode != nil && *h.CountryCode != "" {
		hotel.CountryCode = *h.CountryCode
	} else {
		return nil, false
	}

	if h.City != nil && *h.City != "" {
		hotel.City = *h.City
	} else {
		return nil, false
	}

	if h.Address != nil && *h.Address != "" {
		hotel.Address = h.Address
	}

	if h.Latitude != nil && validateLatitude(*h.Latitude) {
		hotel.Latitude = h.Latitude
	}

	if h.Longitude != nil && validateLongitude(*h.Longitude) {
		hotel.Longitude = h.Longitude
	}

	if h.Rating != nil && validateRating(*h.Rating) {
		hotel.Rating = h.Rating
	}

	images := make([]string, 0)
	for _, image := range h.Images {
		if image.URL != nil && *image.URL != "" {
			images = append(images, *image.URL)
		}
	}
	hotel.Images = images

	return hotel, true
}

func validateLatitude(latitude float64) bool {
	if latitude >= -90.0 && latitude <= 90.0 {
		return true
	}
	return false
}

func validateLongitude(longitude float64) bool {
	if longitude >= -180.0 && longitude <= 180.0 {
		return true
	}
	return false
}

func validateRating(rating float64) bool {
	if rating >= 0 && rating <= 5.0 {
		return true
	}
	return false
}
