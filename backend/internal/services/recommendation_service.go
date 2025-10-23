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

	var recommendations []models.StockRecommendation

	for _, stock := range stocks {
		score, reason := s.calculateScore(&stock)
		recommendations = append(recommendations, models.StockRecommendation{
			Stock:  &stock,
			Score:  score,
			Reason: reason,
		})
	}

	// Sort by score descending
	sort.Slice(recommendations, func(i, j int) bool {
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

	// Factor 1: Target price increase (weight: 40%)
	if targetTo > targetFrom && targetFrom > 0 {
		changePerc := ((targetTo - targetFrom) / targetFrom) * 100
		score += math.Min(changePerc*4, 40)
		reasons = append(reasons, "target price increased")
	}

	// Factor 2: Rating improvement (weight: 30%)
	ratingScore := s.getRatingScore(stock.RatingFrom, stock.RatingTo)
	score += ratingScore
	if ratingScore > 15 {
		reasons = append(reasons, "rating upgraded")
	} else if ratingScore > 0 {
		reasons = append(reasons, "positive rating")
	}

	// Factor 3: Action type (weight: 20%)
	actionScore := s.getActionScore(stock.Action)
	score += actionScore
	if actionScore > 10 {
		reasons = append(reasons, "positive analyst action")
	}

	// Factor 4: Brokerage reputation (weight: 10%)
	if s.isTopBrokerage(stock.Brokerage) {
		score += 10
		reasons = append(reasons, "top-tier brokerage")
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

	if scoreTo > scoreFrom {
		return 30 // Upgraded
	} else if scoreTo >= 4 {
		return 20 // Strong positive rating
	} else if scoreTo == 3 {
		return 10 // Neutral rating
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

func (s *RecommendationService) isTopBrokerage(brokerage string) bool {
	topBrokerages := []string{
		"Goldman Sachs",
		"Morgan Stanley",
		"JP Morgan",
		"Wells Fargo",
		"Bank of America",
		"Citigroup",
		"Barclays",
		"Credit Suisse",
		"UBS",
		"Deutsche Bank",
	}

	brokerageLower := strings.ToLower(brokerage)
	for _, top := range topBrokerages {
		if strings.Contains(brokerageLower, strings.ToLower(top)) {
			return true
		}
	}
	return false
}
