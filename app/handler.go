package app

// Handler represent the parser's files handler
type Handler interface {
	Handle() error
}
