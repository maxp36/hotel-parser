package main

import (
	"flag"
	"log"

	"github.com/maxp36/hotel-parser/app"
	"github.com/maxp36/hotel-parser/app/delivery/file"
	"github.com/maxp36/hotel-parser/app/repository/postgres"
	"github.com/maxp36/hotel-parser/app/service"

	_ "github.com/lib/pq"
)

var dir string

func init() {
	flag.StringVar(&dir, "dir", "raw", "path to the directory which contains files for parsing")

	app.InitConfig()
}

func main() {

	db := app.InitDB()

	repo := postgres.NewRepository(db)

	parser := service.NewParser(repo)

	handler := file.NewFileHandler(dir, parser)

	if err := handler.Handle(); err != nil {
		log.Println(err)
		return
	}
	log.Println("Done!")
}
