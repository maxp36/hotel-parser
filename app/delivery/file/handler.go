package file

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

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

	fpaths := make([]string, 0)

	err := filepath.Walk(h.Dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Mode().IsRegular() {
				fpaths = append(fpaths, path)
			}

			return nil
		})
	if err != nil {
		return err
	}

	for _, p := range fpaths {
		switch filepath.Ext(p) {
		case ".json":
			h.handleJSON(p)
		case ".csv":
			h.handleCSV(p)
		case ".xml":
			h.handleXML(p)
		}
	}

	return nil
}

func (h *fileHandler) handleJSON(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if err := h.Parser.ParseJSON(data); err != nil {
			return err
		}
	}

	return nil
}

func (h *fileHandler) handleCSV(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)

	columns, err := reader.Read()
	if err != nil {
		return err
	}

	for {
		data, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if err := h.Parser.ParseCSV(columns, data); err != nil {
			return err
		}
	}

	return nil
}

func (h *fileHandler) handleXML(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	if err := h.Parser.ParseXML(file); err != nil {
		return err
	}

	return nil
}
