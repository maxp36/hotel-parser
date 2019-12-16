package file

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
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
}

// NewFileHandler inits file handler for parsing files
func NewFileHandler(dir string, p app.Parser) app.Handler {
	return &fileHandler{
		Dir:    dir,
		Parser: p,
	}
}

func (h *fileHandler) Handle() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go h.handleDirs(wg, h.Dir)

	wg.Wait()
}

func (h *fileHandler) handleDirs(wg *sync.WaitGroup, root string) {
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			go h.handleDirs(wg, entry.Name())
		}

		switch filepath.Ext(entry.Name()) {
		case "json":
			h.handleJSON(entry.Name())
		case "csv":
			h.handleCSV(entry.Name())
		case "xml":
			h.handleXML(entry.Name())
		}
	}
	wg.Done()
}

func (h *fileHandler) handleJSON(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	reader := bufio.NewReader(file)

	dec := json.NewDecoder(file)

	for {
		var data []byte
		err := dec.Decode(data)
		if err != nil {
			log.Println(err)
		}

		log.Printf("%s\n", data)
		go h.Parser.ParseJSON(data)
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
