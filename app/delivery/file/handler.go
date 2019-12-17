package file

import (
	"bufio"
	"io"
	"log"
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

func (h *fileHandler) Handle() {
	// wg := &sync.WaitGroup{}

	fpaths := make([]string, 0)

	err := filepath.Walk(h.Dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Handle()", err)
			}
			if info.Mode().IsRegular() {
				fpaths = append(fpaths, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	log.Println(fpaths)

	for _, p := range fpaths {
		switch filepath.Ext(p) {
		case ".json":
			h.handleJSON(p)
		case ".csv":
			h.handleCSV(p)
		case ".xml":
			h.handleXML(p)
		default:
			continue
		}
	}

	// wg.Add(1)
	// go h.handleDirs(wg, h.Dir)

	// wg.Wait()
}

func (h *fileHandler) handleJSON(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Println("handleJSON 1: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("EOF: ", err)
				// break
			}
			log.Println("handleJSON 2: ", err)
			// break
		}

		// log.Printf("handleJSON 3: %s\n", data)
		h.Parser.ParseJSON(data)
	}
}

func (h *fileHandler) handleCSV(path string) {

	// file, err := os.Open(path)
	// if err != nil {
	// 	log.Println(err)
	// }

	// r := csv.NewReader(file)
	// r.R

	// _ = h.Parser.ParseCSV()
}

func (h *fileHandler) handleXML(path string) {

	// _ = h.Parser.ParseXML()
}
