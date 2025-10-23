package models

import "time"

// Stock represents a stock entity
type Stock struct {
	ID          string    `json:"id" db:"id"`
	Ticker      string    `json:"ticker" db:"ticker"`
	Company     string    `json:"company" db:"company"`
	TargetFrom  string    `json:"target_from" db:"target_from"`
	TargetTo    string    `json:"target_to" db:"target_to"`
	Action      string    `json:"action" db:"action"`
	Brokerage   string    `json:"brokerage" db:"brokerage"`
	RatingFrom  string    `json:"rating_from" db:"rating_from"`
	RatingTo    string    `json:"rating_to" db:"rating_to"`
	Time        time.Time `json:"time" db:"time"`
	LastUpdated time.Time `json:"last_updated" db:"last_updated"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// APIResponse represents the response from external API
type APIResponse struct {
	Items    []APIStockItem `json:"items"`
	NextPage string         `json:"next_page"`
}

// APIStockItem represents a single stock item from the API
type APIStockItem struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

// StockRecommendation represents a recommended stock
type StockRecommendation struct {
	Stock  *Stock  `json:"stock"`
	Score  float64 `json:"score"`
	Reason string  `json:"reason"`
}
