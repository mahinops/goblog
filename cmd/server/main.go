package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mokhlesurr031/goblog/config"
	userHttp "github.com/mokhlesurr031/goblog/internal/user/delivery/http"
	userRepo "github.com/mokhlesurr031/goblog/internal/user/repository"
	userUsecase "github.com/mokhlesurr031/goblog/internal/user/usecase"
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

	// Run migrations
	if err := db.Migrate(db.DefaultDB()); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create a Gin router
	r := gin.Default()

	// Initialize repository, usecase, and handler
	userRepo := userRepo.NewUserRepository(db.DefaultDB())
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHttp.NewUserHandler(r, userUsecase)

	// Start the Gin server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
