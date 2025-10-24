# Stock Analyzer Platform - Backend

Backend service built with Go for the Stock Analyzer Platform.

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or higher
- CockroachDB running locally or cloud instance
- API credentials for the stock data source

### Installation

1. Install dependencies:
```bash
go mod download
```

2. Create `.env` file:
```bash
cp .env.example .env
```

3. Create and edit your `.env` with local credentials (do not commit secrets):
```env
# Example keys only – fill with your own values
PORT=3000
ENV=development
DATABASE_URL=postgresql://<user>:<password>@<host>:26257/stocks?sslmode=disable
ALLOWED_ORIGINS=*
```

4. Run migrations:
```bash
go run cmd/migrate/main.go
```

5. Start the server:
```bash
go run cmd/server/main.go
```

## 📁 Project Structure

```
backend/
├── cmd/
│   ├── server/       # Main application entry point
│   └── migrate/      # Database migration tool
├── internal/
│   ├── api/          # HTTP handlers and routing
│   ├── config/       # Configuration management
│   ├── models/       # Data models
│   ├── repository/   # Database layer
│   └── services/     # Business logic
├── pkg/
│   └── utils/        # Shared utilities
└── tests/            # Test files
```

## �️ Database Schema

The service persists analyst actions and price targets. Current schema (CockroachDB/PostgreSQL compatible):

```sql
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
```

Notes:
- `id` is a deterministic hash of `ticker+time` to avoid duplicates of the same event
- `UNIQUE(ticker, time)` guarantees idempotent syncs
- Indexes support search by ticker and time ordering

## 🧪 Testing

Run tests:
```bash
go test ./... -v
```

Run tests with coverage:
```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🛠️ Development

### Using Make commands

```bash
make run          # Run the application
make build        # Build the application
make test         # Run tests
make test-coverage # Run tests with coverage
make clean        # Clean build artifacts
make migrate      # Run database migrations
make deps         # Install dependencies
make fmt          # Format code
```

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `3000` |
| `ENV` | Environment (development/production) | `development` |
| `DATABASE_URL` | CockroachDB/PostgreSQL connection string | – |
| `ALLOWED_ORIGINS` | CORS allowed origins | `*` |

Security: credentials and any provider keys are configured locally via environment variables but are intentionally not documented here. Do not commit secrets.

## 🏗️ Architecture

The backend follows a clean architecture pattern:

1. **API Layer** (`internal/api`): HTTP handlers and routing
2. **Service Layer** (`internal/services`): Business logic
3. **Repository Layer** (`internal/repository`): Database operations
4. **Models** (`internal/models`): Data structures

Recommendation logic lives in `internal/services/recommendation_service.go` and scores signals from analyst actions, ratings, and target price changes. The sync page limit is user-configurable from the UI; the backend enforces safe defaults.

## 📦 Dependencies

- `gin-gonic/gin` - HTTP web framework
- `lib/pq` - PostgreSQL driver
- `joho/godotenv` - Environment variable management
- `stretchr/testify` - Testing toolkit
