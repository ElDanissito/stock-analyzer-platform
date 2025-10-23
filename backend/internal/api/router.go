package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *StockHandler, allowedOrigins string) *gin.Engine {
	router := gin.Default()

	// CORS middleware
	router.Use(corsMiddleware(allowedOrigins))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Stocks routes
		api.GET("/stocks", handler.GetStocks)
		api.GET("/stocks/:id", handler.GetStockByID)
		api.GET("/stocks/search", handler.SearchStocks)
		api.POST("/sync", handler.SyncStocks)

		// Recommendations route
		api.GET("/recommendations", handler.GetRecommendations)
	}

	return router
}

func corsMiddleware(allowedOrigins string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// Check if origin is allowed
		if allowedOrigins == "*" || contains(strings.Split(allowedOrigins, ","), origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.TrimSpace(s) == item {
			return true
		}
	}
	return false
}
