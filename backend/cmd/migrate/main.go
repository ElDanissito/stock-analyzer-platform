package main

import (
	"log"

	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/config"
	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/repository"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := repository.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := db.InitSchema(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
}
