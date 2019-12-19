package service

import (
	"github.com/maxp36/hotel-parser/app"
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
