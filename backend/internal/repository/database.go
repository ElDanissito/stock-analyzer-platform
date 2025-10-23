package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(databaseURL string) (*Database, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Successfully connected to database")
	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}

func (d *Database) InitSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS stocks (
		id VARCHAR(255) PRIMARY KEY,
		ticker VARCHAR(50) NOT NULL,
		company VARCHAR(255) NOT NULL,
		target_from VARCHAR(50),
		target_to VARCHAR(50),
		action VARCHAR(100),
		brokerage VARCHAR(255),
		rating_from VARCHAR(50),
		rating_to VARCHAR(50),
		time TIMESTAMP NOT NULL,
		last_updated TIMESTAMP NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(ticker, time)
	);

	CREATE INDEX IF NOT EXISTS idx_stocks_ticker ON stocks(ticker);
	CREATE INDEX IF NOT EXISTS idx_stocks_time ON stocks(time);
	CREATE INDEX IF NOT EXISTS idx_stocks_last_updated ON stocks(last_updated);
	`

	_, err := d.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating schema: %w", err)
	}

	log.Println("Database schema initialized successfully")
	return nil
}
