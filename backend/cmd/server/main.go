package main

import (
	"log"

	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/api"
	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/config"
	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/repository"
	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/services"
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

	// Initialize database schema
	if err := db.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Initialize repositories
	stockRepo := repository.NewStockRepository(db)

	// Initialize services
	stockService := services.NewStockService(stockRepo, cfg.StockAPIURL, cfg.StockAPIKey)
	recommendationService := services.NewRecommendationService(stockService)

	// Initialize handlers
	stockHandler := api.NewStockHandler(stockService, recommendationService)

	// Setup router
	router := api.SetupRouter(stockHandler, cfg.AllowedOrigins)

	// Start server
	log.Printf("Server starting on port %s...", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
