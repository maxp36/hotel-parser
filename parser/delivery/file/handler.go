package file

import "github.com/maxp36/hotel-parser/parser"

// fileHandler represent the file handler for parsing files
type fileHandler struct {
	Parser parser.Parser
}

// NewFileHandler inits file handler for parsing files
func NewFileHandler(p parser.Parser) parser.Handler {
	return &fileHandler{
		Parser: p,
	}
}

func (h *fileHandler) Handle(dir string) error {

	return nil
}
