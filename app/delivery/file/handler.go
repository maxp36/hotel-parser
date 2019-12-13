package file

import (
	"io/ioutil"

	"github.com/maxp36/hotel-parser/app"
)

// fileHandler represent the file handler for parsing files
type fileHandler struct {
	Dir    string
	Parser app.Parser
}

// NewFileHandler inits file handler for parsing files
func NewFileHandler(dir string, p app.Parser) app.Handler {
	return &fileHandler{
		Dir:    dir,
		Parser: p,
	}
}

func (h *fileHandler) Handle() error {

	return nil
}

func (h *fileHandler) handleDirs(root string) error {
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			go h.handleDirs(entry.Name())
		}

		// TODO: handle ffile
	}

	return nil
}
