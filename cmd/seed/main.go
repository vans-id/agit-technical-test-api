package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vans-id/agit-technical-test-api.git/migration"
	"github.com/vans-id/agit-technical-test-api.git/pkg/pg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := pg.SetupDB()
	if err != nil {
		log.Fatalln(err)
	}
	migration.Seed(db)
}
