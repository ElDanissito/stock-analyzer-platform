package services

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/models"
)

type RecommendationService struct {
	stockService *StockService
}

func NewRecommendationService(stockService *StockService) *RecommendationService {
	return &RecommendationService{
		stockService: stockService,
	}
}

func (s *RecommendationService) GetRecommendations(limit int) ([]models.StockRecommendation, error) {
	stocks, err := s.stockService.GetAllStocks(1000, 0)
	if err != nil {
		return nil, err
	}

	// Group by ticker and keep only the most recent
	tickerMap := make(map[string]models.Stock)
	for _, stock := range stocks {
		existing, exists := tickerMap[stock.Ticker]
		if !exists || stock.Time.After(existing.Time) {
			tickerMap[stock.Ticker] = stock
		}
	}

	// Calculate scores for unique tickers
	var recommendations []models.StockRecommendation
	for _, stock := range tickerMap {
		// IMPORTANT: copy loop variable before taking address to avoid pointer aliasing
		stockCopy := stock
		score, reason := s.calculateScore(&stockCopy)
		// Only include stocks with meaningful scores
		if score > 0 {
			recommendations = append(recommendations, models.StockRecommendation{
				Stock:  &stockCopy,
				Score:  score,
				Reason: reason,
			})
		}
	}

	// Sort by score descending
	sort.Slice(recommendations, func(i, j int) bool {
		if recommendations[i].Score == recommendations[j].Score {
			// Secondary sort by ticker for stable ordering
			return recommendations[i].Stock.Ticker < recommendations[j].Stock.Ticker
		}
		return recommendations[i].Score > recommendations[j].Score
	})

	// Return top N recommendations
	if len(recommendations) > limit {
		recommendations = recommendations[:limit]
	}

	return recommendations, nil
}

func (s *RecommendationService) calculateScore(stock *models.Stock) (float64, string) {
	var score float64
	var reasons []string

	// Extract price change from target_from and target_to
	targetFrom := s.parsePrice(stock.TargetFrom)
	targetTo := s.parsePrice(stock.TargetTo)
	var changePerc float64
	if targetFrom > 0 {
		changePerc = ((targetTo - targetFrom) / targetFrom) * 100
	}

	// Factor 1: Target price increase (weight: 40%)
	if changePerc > 0 {
		score += math.Min(changePerc*4, 40)
		reasons = append(reasons, "target price increased")
	}

	// Factor 2: Rating improvement (weight: 30%)
	ratingScore := s.getRatingScore(stock.RatingFrom, stock.RatingTo)
	score += ratingScore
	if ratingScore >= 30 {
		reasons = append(reasons, "rating upgraded")
	} else if ratingScore >= 20 {
		reasons = append(reasons, "strong rating initiated")
	} else if ratingScore >= 10 {
		reasons = append(reasons, "positive rating maintained")
	}

	// Factor 3: Action type (weight: 20%)
	actionScore := s.getActionScore(stock.Action)
	score += actionScore
	if actionScore > 10 {
		reasons = append(reasons, "positive analyst action")
	}

	// Bonus: Momentum synergy (up to +10)
	// Reward strong alignment when there is a big target hike and an upgrade action/rating
	actionLower := strings.ToLower(stock.Action)
	if changePerc >= 50 && (ratingScore >= 30) && (strings.Contains(actionLower, "upgraded") || strings.Contains(actionLower, "raised")) {
		score += 10
		reasons = append(reasons, "strong multi-signal momentum")
	} else if changePerc >= 25 && (ratingScore >= 30 || strings.Contains(actionLower, "upgraded") || strings.Contains(actionLower, "raised")) {
		score += 5
		reasons = append(reasons, "strong momentum")
	}

	// Normalize score to 0-100
	score = math.Min(score, 100)
	score = math.Max(score, 0)

	reason := "Good fundamentals"
	if len(reasons) > 0 {
		reason = strings.Join(reasons, ", ")
	}

	return score, reason
}

func (s *RecommendationService) parsePrice(priceStr string) float64 {
	// Remove $ and any whitespace
	cleaned := strings.TrimSpace(strings.ReplaceAll(priceStr, "$", ""))
	price, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0
	}
	return price
}

func (s *RecommendationService) getRatingScore(ratingFrom, ratingTo string) float64 {
	ratings := map[string]int{
		"Strong Buy":     5,
		"Buy":            4,
		"Outperform":     4,
		"Hold":           3,
		"Market Perform": 3,
		"Neutral":        3,
		"Equal Weight":   3,
		"Underperform":   2,
		"Underweight":    2,
		"Sell":           1,
	}

	scoreFrom := ratings[ratingFrom]
	scoreTo := ratings[ratingTo]

	// Only give points if there was an actual upgrade or strong positive rating
	if scoreTo > scoreFrom {
		return 30 // Upgraded
	} else if scoreTo >= 4 && scoreFrom == 0 {
		// New coverage with strong rating (no previous rating)
		return 20
	} else if scoreTo == scoreFrom && scoreTo >= 4 {
		// Maintained strong rating (reiterated)
		return 10
	}
	return 0
}

func (s *RecommendationService) getActionScore(action string) float64 {
	action = strings.ToLower(action)
	if strings.Contains(action, "raised") || strings.Contains(action, "upgraded") {
		return 20
	} else if strings.Contains(action, "initiated") && !strings.Contains(action, "lowered") {
		return 15
	} else if strings.Contains(action, "reiterated") {
		return 10
	}
	return 5
}
