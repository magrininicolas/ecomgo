package main

import (
	"log"

	"github.com/magrininicolas/ecomgo/cmd/api"
	"github.com/magrininicolas/ecomgo/db"
)

func main() {
	db, err := db.NewPostgreSQLStorage()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewApiServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
