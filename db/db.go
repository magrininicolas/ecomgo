package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/magrininicolas/ecomgo/config"
)

func NewPostgreSQLStorage() (*sqlx.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.PublicHost, config.Envs.Port, config.Envs.DBUser, config.Envs.DBPasswd, config.Envs.DBName)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
