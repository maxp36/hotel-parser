package parser

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// InitDB inits database connection for hotels' data
func InitDB() *sqlx.DB {
	user := os.Getenv("HOTEL_PARSER_DB_USER")
	pass := os.Getenv("HOTEL_PARSER_DB_PASSWORD")
	host := os.Getenv("HOTEL_PARSER_DB_HOST")
	port := os.Getenv("HOTEL_PARSER_DB_PORT")
	name := os.Getenv("HOTEL_PARSER_DB_NAME")

	source := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pass,
		host,
		port,
		name)

	var err error
	db, err := sqlx.Open("postgres", source)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	log.Println("DB connected!")

	return db
}
