package file

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

	"github.com/maxp36/hotel-parser/app"
	"github.com/maxp36/wgext"
)

// fileHandler represent the file handler for parsing files
type fileHandler struct {
	Dir    string
	Parser app.Parser
	wg     *wgext.WaitGroup
}

// NewFileHandler inits file handler for parsing files
func NewFileHandler(dir string, p app.Parser) app.Handler {
	return &fileHandler{
		Dir:    dir,
		Parser: p,
	}
}

func (h *fileHandler) Handle() error {

	h.wg = wgext.NewWaitGroup()

	fpaths := make([]string, 0)

	// collect all file paths
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

	// parse all necessary files
	for _, p := range fpaths {
		switch filepath.Ext(p) {
		case ".xml":
			h.wg.Add(1)
			go h.handleXML(p)
		case ".json":
			h.wg.Add(1)
			go h.handleJSON(p)
		case ".csv":
			h.wg.Add(1)
			go h.handleCSV(p)
		}
	}

	return h.wg.Wait()
}

func (h *fileHandler) handleJSON(path string) {

	file, err := os.Open(path)
	if err != nil {
		h.wg.Fail(err)
		return
	}

	reader := bufio.NewReader(file)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			h.wg.Fail(err)
			return
		}

		h.wg.Add(1)
		go func(wg *wgext.WaitGroup, p app.Parser, d []byte) {
			err := p.ParseJSON(d)
			if err != nil && err != io.EOF {
				wg.Fail(err)
				return
			}
			wg.Done()
		}(h.wg, h.Parser, data)
	}

	h.wg.Done()
}

func (h *fileHandler) handleCSV(path string) {

	file, err := os.Open(path)
	if err != nil {
		h.wg.Fail(err)
		return
	}

	reader := csv.NewReader(file)

	// read csv columns
	columns, err := reader.Read()
	if err != nil {
		h.wg.Fail(err)
		return
	}

	for {
		data, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			h.wg.Fail(err)
			return
		}

		h.wg.Add(1)
		go func(wg *wgext.WaitGroup, p app.Parser, cs, d []string) {
			if err := p.ParseCSV(cs, d); err != nil {
				wg.Fail(err)
				return
			}
			wg.Done()
		}(h.wg, h.Parser, columns, data)
	}

	h.wg.Done()
}

func (h *fileHandler) handleXML(path string) {

	file, err := os.Open(path)
	if err != nil {
		h.wg.Fail(err)
		return
	}

	if err := h.Parser.ParseXML(file); err != nil {
		h.wg.Fail(err)
		return
	}

	h.wg.Done()
}
