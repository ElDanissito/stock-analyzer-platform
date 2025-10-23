package repository

import (
	"database/sql"
	"fmt"

	"github.com/ElDanissito/stock-analyzer-platform/backend/internal/models"
)

type StockRepository struct {
	db *Database
}

func NewStockRepository(db *Database) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) Create(stock *models.Stock) error {
	query := `
		INSERT INTO stocks (id, ticker, company, target_from, target_to, action, brokerage, rating_from, rating_to, time, last_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (ticker, time) DO UPDATE SET
			target_from = EXCLUDED.target_from,
			target_to = EXCLUDED.target_to,
			action = EXCLUDED.action,
			brokerage = EXCLUDED.brokerage,
			rating_from = EXCLUDED.rating_from,
			rating_to = EXCLUDED.rating_to,
			last_updated = EXCLUDED.last_updated
	`

	_, err := r.db.DB.Exec(query,
		stock.ID, stock.Ticker, stock.Company, stock.TargetFrom, stock.TargetTo,
		stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo,
		stock.Time, stock.LastUpdated)

	return err
}

func (r *StockRepository) GetAll(limit, offset int) ([]models.Stock, error) {
	query := `
		SELECT id, ticker, company, target_from, target_to, action, brokerage, rating_from, rating_to, time, last_updated, created_at
		FROM stocks
		ORDER BY time DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanStocks(rows)
}

func (r *StockRepository) GetByID(id string) (*models.Stock, error) {
	query := `
		SELECT id, ticker, company, target_from, target_to, action, brokerage, rating_from, rating_to, time, last_updated, created_at
		FROM stocks
		WHERE id = $1
	`

	var stock models.Stock
	err := r.db.DB.QueryRow(query, id).Scan(
		&stock.ID, &stock.Ticker, &stock.Company, &stock.TargetFrom, &stock.TargetTo,
		&stock.Action, &stock.Brokerage, &stock.RatingFrom, &stock.RatingTo,
		&stock.Time, &stock.LastUpdated, &stock.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("stock not found")
	}

	return &stock, err
}

func (r *StockRepository) Search(query string) ([]models.Stock, error) {
	searchQuery := `
		SELECT id, ticker, company, target_from, target_to, action, brokerage, rating_from, rating_to, time, last_updated, created_at
		FROM stocks
		WHERE ticker ILIKE $1 OR company ILIKE $1
		ORDER BY time DESC
	`

	rows, err := r.db.DB.Query(searchQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanStocks(rows)
}

func (r *StockRepository) Count() (int, error) {
	var count int
	err := r.db.DB.QueryRow("SELECT COUNT(*) FROM stocks").Scan(&count)
	return count, err
}

func (r *StockRepository) scanStocks(rows *sql.Rows) ([]models.Stock, error) {
	var stocks []models.Stock

	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(
			&stock.ID, &stock.Ticker, &stock.Company, &stock.TargetFrom, &stock.TargetTo,
			&stock.Action, &stock.Brokerage, &stock.RatingFrom, &stock.RatingTo,
			&stock.Time, &stock.LastUpdated, &stock.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, rows.Err()
}
