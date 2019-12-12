package service

import "github.com/maxp36/hotel-parser/parser"

type jsonParser struct {
	R parser.Repository
}

// NewJSONParser creates new service objects that implements Service interface
func NewJSONParser(r parser.Repository) parser.Parser {
	return &jsonParser{
		R: r,
	}
}

func (s *jsonParser) Parse(data []byte) error {
	return nil
}
