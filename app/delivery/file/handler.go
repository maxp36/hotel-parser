package file

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/maxp36/hotel-parser/app"
)

// fileHandler represent the file handler for parsing files
type fileHandler struct {
	Dir    string
	Parser app.Parser
	wg     *sync.WaitGroup
	errs   chan error
}

// NewFileHandler inits file handler for parsing files
func NewFileHandler(dir string, p app.Parser) app.Handler {
	return &fileHandler{
		Dir:    dir,
		Parser: p,
		wg:     &sync.WaitGroup{},
		errs:   make(chan error),
	}
}

func (h *fileHandler) Handle() error {

	h.wg = &sync.WaitGroup{}
	h.errs = make(chan error)

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
			go h.handleJSON(p)
			h.wg.Add(1)
		case ".csv":
			go h.handleCSV(p)
			h.wg.Add(1)
		case ".xml":
			go h.handleXML(p)
			h.wg.Add(1)
		}
	}

	// select {
	// case err := <-h.errs:
	// 	return err
	// case h.wg.Wait():

	// }
	h.wg.Wait()

	return <-h.errs
}

func (h *fileHandler) handleJSON(path string) {

	defer h.wg.Done()
	defer log.Println("done 1")

	file, err := os.Open(path)
	if err != nil {
		h.errs <- err
		return
	}

	reader := bufio.NewReader(file)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("handleJSON 1")
				break
			}
			log.Println("handleJSON 2")
			h.errs <- err
			return
		}

		if err := h.Parser.ParseJSON(data); err != nil {
			log.Println("handleJSON 3")
			h.errs <- err
			return
		}
		log.Println("handleJSON 4")
	}
}

func (h *fileHandler) handleCSV(path string) {

	defer h.wg.Done()
	defer log.Println("done 2")

	file, err := os.Open(path)
	if err != nil {
		h.errs <- err
		return
	}

	reader := csv.NewReader(file)

	columns, err := reader.Read()
	if err != nil {
		h.errs <- err
		return
	}

	for {
		data, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			h.errs <- err
			return
		}

		if err := h.Parser.ParseCSV(columns, data); err != nil {
			h.errs <- err
			return
		}
	}
}

func (h *fileHandler) handleXML(path string) {

	defer h.wg.Done()
	defer log.Println("done 3")

	file, err := os.Open(path)
	if err != nil {
		h.errs <- err
		return
	}

	if err := h.Parser.ParseXML(file); err != nil {
		h.errs <- err
		return
	}
}
