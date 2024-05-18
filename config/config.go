package config

import (
	"fmt"
	"log"
	"os"
)

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

var (
	db Database
)

func DB() *Database {
	return &db
}

func LoadDBEnvs() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	if host == "" || port == "" || name == "" || username == "" || password == "" {
		log.Fatalf("One or more database required environment variables are missing")
	}

	db = Database{
		Name:     name,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
	fmt.Println(db)
}
