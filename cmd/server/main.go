package main

import (
	"log"

	"github.com/mokhlesurr031/goblog/config"
	"github.com/mokhlesurr031/goblog/pkg/db"
)

func main() {
	// Load configuration
	config.LoadDBEnvs()

	log.Println("Connecting database")
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer db.CloseDB()

}
