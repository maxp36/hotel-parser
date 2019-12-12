package models

import "time"

// Hotel is base datamodel of hotel
type Hotel struct {
	ID        int64      `json:"id" db:"id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`

	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	CountryCode string  `json:"country_code" db:"country_code"`
	City        string  `json:"city" db:"city"`
	Address     string  `json:"address" db:"address"`
	Latitude    float64 `json:"latitude" db:"latitude"`
	Longitude   float64 `json:"longitude" db:"longitude"`
	Rating      float64 `json:"rating" db:"rating"`
}
