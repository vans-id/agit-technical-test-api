package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vans-id/agit-technical-test-api.git/internal/routes"
	"github.com/vans-id/agit-technical-test-api.git/pkg/pg"
	"github.com/vans-id/agit-technical-test-api.git/shared/constants"
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

	r := routes.NewRouter(routes.GetConfig(db))
	server := http.Server{
		Addr:    constants.APP_PORT,
		Handler: r,
	}
	server.ListenAndServe()
}
