package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/abelmalu/CafeteriaAccessControl/config"
)

type App struct {
	Config *config.Config
	DB     *sql.DB
	Server *http.Server
}

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {

		log.Fatalf("Failed to connect to DB: %v", err)
	}

}
