# 📈 Stock Analyzer Platform

A full-stack stock analysis system that retrieves, stores, and analyzes stock market data to provide intelligent investment recommendations.

## 🎯 Project Overview

This platform is designed to help investors make informed decisions by:
- Fetching real-time stock data from external APIs
- Storing and managing stock information efficiently
- Providing an intuitive interface for browsing and analyzing stocks
- Recommending the best investment opportunities based on data analysis

## 🛠️ Tech Stack

### Backend
- **Go (Golang)** - High-performance API server
- **CockroachDB** - Distributed SQL database for scalability and reliability

### Frontend
- **Vue 3** - Progressive JavaScript framework
- **TypeScript** - Type-safe development
- **Pinia** - State management
- **Tailwind CSS** - Utility-first styling

## ✨ Features

### Core Functionality
- ✅ Real-time stock data fetching from external API
- ✅ Persistent storage in CockroachDB
- ✅ RESTful API for data access
- ✅ Search and filter stocks
- ✅ Sort by various metrics
- ✅ Detailed stock information views

### Smart Recommendations
- 🤖 Intelligent algorithm to recommend best stocks
- 📊 Data analysis based on market trends
- 💡 Investment insights

### Quality Assurance
- 🧪 Comprehensive unit tests
- ✅ Test coverage for critical components

## 📋 Requirements

- **Go** 1.21 or higher
- **Node.js** 18+ and npm/yarn
- **CockroachDB** (local or cloud instance)
- **Git**

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/ElDanissito/stock-analyzer-platform.git
cd stock-analyzer-platform
```

### 2. Backend Setup

```bash
cd backend
go mod download
```

Create a `.env` file with your configuration (do not include secrets in Git):
```env
DATABASE_URL=postgresql://<user>@<host>:<port>/<database>
PORT=8080
```

Run the backend:
```bash
go run cmd/server/main.go
```

### 3. Frontend Setup

```bash
cd frontend
npm install
```

Optionally, create a `.env` file for the frontend to point to your backend base URL (no defaults are included here to avoid leaking local endpoints).

Run the development server:
```bash
npm run dev
```

### 4. Database Setup

Initialize CockroachDB and run migrations:
```bash
cd backend
go run cmd/migrate/main.go
```

Optional: start CockroachDB locally with Docker (Windows PowerShell):

If you prefer running CockroachDB via Docker, this repo includes helper scripts under `scripts/`:

```powershell
# From the repository root (PowerShell)
./scripts/start-services.ps1   # Start CockroachDB
./scripts/init-db.ps1          # Create the 'stocks' database

# View logs or stop services when needed
./scripts/view-logs.ps1
./scripts/stop-services.ps1
```

CockroachDB Admin UI: http://localhost:8080

## 🧪 Running Tests

### Backend Tests
```bash
cd backend
go test ./... -v
```

### Frontend Tests
```bash
cd frontend
npm run test
```

## 📁 Project Structure

```
stock-analyzer-platform/
├── backend/
│   ├── cmd/              # Application entrypoints
│   ├── internal/         # Private application code
│   │   ├── api/          # API handlers
│   │   ├── models/       # Data models
│   │   ├── services/     # Business logic
│   │   └── repository/   # Database layer
│   ├── pkg/              # Public libraries
│   └── tests/            # Test files
├── frontend/
│   ├── src/
│   │   ├── components/   # Vue components
│   │   ├── views/        # Page views
│   │   ├── stores/       # Pinia stores
│   │   ├── services/     # API services
│   │   └── types/        # TypeScript types
│   └── tests/            # Test files
└── Documentation.md      # Technical documentation
```

## 🔄 Development Workflow

1. **Data Collection**: Backend fetches stock data from external API
2. **Storage**: Data is persisted in CockroachDB
3. **API Layer**: Go backend exposes RESTful endpoints
4. **Frontend**: Vue 3 application consumes API and displays data
5. **Analysis**: Recommendation engine processes data and provides insights

## 🎨 UI Features

- **Dashboard**: Overview of all stocks
- **Search & Filter**: Find specific stocks quickly
- **Sort Options**: Order by price, volume, performance, etc.
- **Detail View**: Comprehensive information about each stock
- **Recommendations**: AI-powered investment suggestions

## Security & Privacy

- This repository intentionally omits provider endpoints and API keys
- Configure any credentials locally via environment variables
- Keep .env files out of version control


## 📄 License

This project is part of a technical challenge and is for educational purposes.

## 👨‍💻 Author

**Daniel Rodriguez**
- GitHub: [@ElDanissito](https://github.com/ElDanissito)

## 🙏 Acknowledgments

This project was developed as part of a technical assessment to demonstrate:
- Full-stack development capabilities
- Clean architecture principles
- Modern web technologies
- Problem-solving skills
- Code quality and testing practices

---

**Note**: This is a portfolio project demonstrating technical skills in modern web development.