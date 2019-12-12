package parser

// Handler represent the parser's files handler
type Handler interface {
	Handle(dir string) error
}
