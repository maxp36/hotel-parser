package main

import (
	"flag"

	"github.com/maxp36/hotel-parser/app"
	"github.com/maxp36/hotel-parser/app/delivery/file"
	"github.com/maxp36/hotel-parser/app/repository/postgres"
	"github.com/maxp36/hotel-parser/app/service"
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

	_ = file.NewFileHandler(dir, parser)
}
