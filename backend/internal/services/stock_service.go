package services

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/models"
	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/repository"
)

const (
	// DEFAULT_MAX_PAGES es el límite por defecto de páginas a sincronizar (~200 stocks)
	DEFAULT_MAX_PAGES = 20
	// ABSOLUTE_MAX_PAGES es el límite máximo permitido (~1000 stocks)
	ABSOLUTE_MAX_PAGES = 100
)

type StockService struct {
	repo       *repository.StockRepository
	apiURL     string
	apiKey     string
	httpClient *http.Client
}

func NewStockService(repo *repository.StockRepository, apiURL, apiKey string) *StockService {
	return &StockService{
		repo:   repo,
		apiURL: apiURL,
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *StockService) FetchAndStoreStocks(maxPages int) error {
	// Validar y aplicar límites
	if maxPages <= 0 {
		maxPages = DEFAULT_MAX_PAGES
	}
	if maxPages > ABSOLUTE_MAX_PAGES {
		maxPages = ABSOLUTE_MAX_PAGES
	}

	nextPage := ""
	totalFetched := 0
	pageCount := 0

	log.Printf("Starting stock synchronization (max %d pages)...", maxPages)

	for {
		stocks, next, err := s.fetchStocksFromAPI(nextPage)
		if err != nil {
			return fmt.Errorf("error fetching stocks: %w", err)
		}

		for _, stock := range stocks {
			if err := s.repo.Create(&stock); err != nil {
				log.Printf("Error storing stock %s: %v", stock.Ticker, err)
				continue
			}
		}

		totalFetched += len(stocks)
		pageCount++
		log.Printf("Fetched and stored %d stocks (total: %d, page: %d/%d)", len(stocks), totalFetched, pageCount, maxPages)

		if next == "" {
			log.Printf("Successfully fetched all available stocks: %d total", totalFetched)
			break
		}

		if pageCount >= maxPages {
			log.Printf("Reached maximum sync pages limit (%d pages, %d stocks)", maxPages, totalFetched)
			break
		}

		nextPage = next
	}

	log.Printf("Sync completed: %d stocks stored from %d pages", totalFetched, pageCount)
	return nil
}

func (s *StockService) fetchStocksFromAPI(nextPage string) ([]models.Stock, string, error) {
	url := s.apiURL
	if nextPage != "" {
		url = fmt.Sprintf("%s?next_page=%s", s.apiURL, nextPage)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	var apiResponse models.APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, "", err
	}

	stocks := s.parseStocksFromResponse(apiResponse.Items)
	return stocks, apiResponse.NextPage, nil
}

func (s *StockService) parseStocksFromResponse(items []models.APIStockItem) []models.Stock {
	var stocks []models.Stock

	for _, item := range items {
		// Parse time
		parsedTime, err := time.Parse(time.RFC3339, item.Time)
		if err != nil {
			log.Printf("Error parsing time for %s: %v", item.Ticker, err)
			parsedTime = time.Now()
		}

		// Generate unique ID using ticker + time
		id := s.generateID(item.Ticker, item.Time)

		stock := models.Stock{
			ID:          id,
			Ticker:      item.Ticker,
			Company:     item.Company,
			TargetFrom:  item.TargetFrom,
			TargetTo:    item.TargetTo,
			Action:      item.Action,
			Brokerage:   item.Brokerage,
			RatingFrom:  item.RatingFrom,
			RatingTo:    item.RatingTo,
			Time:        parsedTime,
			LastUpdated: time.Now(),
		}
		stocks = append(stocks, stock)
	}

	return stocks
}

func (s *StockService) generateID(ticker, timestamp string) string {
	hash := sha256.Sum256([]byte(ticker + timestamp))
	return fmt.Sprintf("%x", hash[:16])
}

func (s *StockService) GetAllStocks(limit, offset int) ([]models.Stock, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *StockService) GetStockByID(id string) (*models.Stock, error) {
	return s.repo.GetByID(id)
}

func (s *StockService) SearchStocks(query string) ([]models.Stock, error) {
	return s.repo.Search(query)
}

func (s *StockService) GetTotalCount() (int, error) {
	return s.repo.Count()
}
