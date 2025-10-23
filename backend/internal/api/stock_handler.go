package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	stockService          *services.StockService
	recommendationService *services.RecommendationService
}

func NewStockHandler(stockService *services.StockService, recommendationService *services.RecommendationService) *StockHandler {
	return &StockHandler{
		stockService:          stockService,
		recommendationService: recommendationService,
	}
}

func (h *StockHandler) GetStocks(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	stocks, err := h.stockService.GetAllStocks(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.stockService.GetTotalCount()

	c.JSON(http.StatusOK, gin.H{
		"data":   stocks,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (h *StockHandler) GetStockByID(c *gin.Context) {
	id := c.Param("id")

	stock, err := h.stockService.GetStockByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

func (h *StockHandler) SearchStocks(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	stocks, err := h.stockService.SearchStocks(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": stocks,
	})
}

type SyncRequest struct {
	Pages int `json:"pages"`
}

func (h *StockHandler) SyncStocks(c *gin.Context) {
	var req SyncRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Si no hay body o está mal formado, usar valor por defecto
		req.Pages = 0
	}

	go func() {
		if err := h.stockService.FetchAndStoreStocks(req.Pages); err != nil {
			// Log error but don't block the response
			log.Printf("Sync error: %v", err)
		}
	}()

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Stock synchronization started",
		"pages":   req.Pages,
	})
}

func (h *StockHandler) GetRecommendations(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	recommendations, err := h.recommendationService.GetRecommendations(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": recommendations,
	})
}
