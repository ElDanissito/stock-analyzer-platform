# Stock Analyzer Platform - Backend

Backend API built with Go for the Stock Analyzer Platform.

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

3. Edit `.env` with your configuration:
```env
PORT=8080
ENV=development
DATABASE_URL=postgresql://user:password@localhost:26257/stocks?sslmode=disable
STOCK_API_URL=https://api.karenai.click/swechallenge/list
STOCK_API_KEY=your_api_key_here
ALLOWED_ORIGINS=http://localhost:5173
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

## 🔌 API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### Stocks
- `GET /api/stocks` - Get all stocks (with pagination)
  - Query params: `limit` (default: 50), `offset` (default: 0)
- `GET /api/stocks/:id` - Get stock by ID
- `GET /api/stocks/search?q=query` - Search stocks by symbol or name
- `POST /api/sync` - Sync stocks from external API (async)

### Recommendations
- `GET /api/recommendations` - Get recommended stocks
  - Query params: `limit` (default: 10)

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
| `PORT` | Server port | `8080` |
| `ENV` | Environment (development/production) | `development` |
| `DATABASE_URL` | PostgreSQL connection string | - |
| `STOCK_API_URL` | External stock API URL | - |
| `STOCK_API_KEY` | API key for authentication | - |
| `ALLOWED_ORIGINS` | CORS allowed origins | `*` |

## 🏗️ Architecture

The backend follows a clean architecture pattern:

1. **API Layer** (`internal/api`): HTTP handlers and routing
2. **Service Layer** (`internal/services`): Business logic
3. **Repository Layer** (`internal/repository`): Database operations
4. **Models** (`internal/models`): Data structures

## 📦 Dependencies

- `gin-gonic/gin` - HTTP web framework
- `lib/pq` - PostgreSQL driver
- `joho/godotenv` - Environment variable management
- `stretchr/testify` - Testing toolkit
