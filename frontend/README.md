# Stock Analyzer Platform - Frontend

Vue 3 + TypeScript + Pinia + Tailwind CSS frontend application.

## 🚀 Quick Start

### Prerequisites
- Node.js 18+ and npm

### Installation

```bash
# Install dependencies
npm install
```

### Development

```bash
# Run development server
npm run dev
```

The app will be available at http://localhost:5173

### Build for Production

```bash
# Build
npm run build

# Preview production build
npm run preview
```

## 📁 Project Structure

```
frontend/
├── src/
│   ├── assets/          # Static assets and global styles
│   ├── components/      # Reusable Vue components
│   │   ├── StockCard.vue
│   │   └── RecommendationCard.vue
│   ├── views/           # Page components
│   │   ├── HomeView.vue
│   │   └── RecommendationsView.vue
│   ├── stores/          # Pinia stores
│   │   └── stockStore.ts
│   ├── services/        # API services
│   │   └── api.ts
│   ├── types/           # TypeScript type definitions
│   │   └── stock.ts
│   ├── router/          # Vue Router configuration
│   │   └── index.ts
│   ├── App.vue          # Root component
│   └── main.ts          # Application entry point
├── public/              # Public static files
├── index.html          # HTML entry point
└── package.json        # Dependencies and scripts
```

## 🎨 Features

### Home Dashboard
- **Stock List**: View all stocks with pagination
- **Search**: Search stocks by ticker or company name
- **Sync**: Fetch latest data from external API
- **Real-time Stats**: Total stocks, displayed count, status

### Recommendations
- **AI-Powered**: Smart recommendations based on multiple factors
- **Top 3 Picks**: Highlighted best investment opportunities
- **Score System**: 0-100 score for each recommendation
- **Detailed Insights**: Reasons for each recommendation

### Components

#### StockCard
Displays individual stock information:
- Ticker and company name
- Target price (from → to)
- Action and brokerage
- Rating with color coding
- Timestamp

#### RecommendationCard
Premium card for top recommendations:
- Score visualization
- Gradient background
- Detailed reasoning
- Key metrics

## 🔌 API Integration

The frontend communicates with the backend API:

```typescript
// Available API methods
stocksApi.getStocks(limit, offset)
stocksApi.getStockById(id)
stocksApi.searchStocks(query)
stocksApi.syncStocks()
stocksApi.getRecommendations(limit)
stocksApi.healthCheck()
```

## 🎨 Styling

- **Tailwind CSS** for utility-first styling
- **Custom Color Palette** with primary blue theme
- **Responsive Design** for mobile, tablet, and desktop
- **Smooth Animations** for better UX

## 🛠️ Development

### Adding a New Component

```bash
# Create component file
touch src/components/MyComponent.vue
```

```vue
<template>
  <div>My Component</div>
</template>

<script setup lang="ts">
// Component logic
</script>
```

### Using the Store

```vue
<script setup lang="ts">
import { useStockStore } from '@/stores/stockStore'

const stockStore = useStockStore()

// Fetch stocks
await stockStore.fetchStocks()

// Access state
console.log(stockStore.stocks)
console.log(stockStore.loading)
</script>
```

## 🧪 Testing

```bash
# Run linter
npm run lint
```

## 📝 Environment Variables

Create a `.env` file:

```env
VITE_API_URL=http://localhost:3000
```

## 🎯 TODO

- [ ] Add stock detail modal
- [ ] Implement filters (by rating, brokerage, etc.)
- [ ] Add sorting options
- [ ] Charts and visualizations
- [ ] Export to CSV
- [ ] Dark mode
- [ ] Unit tests
- [ ] E2E tests

## 📚 Tech Stack

- **Vue 3** - Progressive JavaScript framework
- **TypeScript** - Type-safe development
- **Pinia** - State management
- **Vue Router** - Routing
- **Tailwind CSS** - Styling
- **Axios** - HTTP client
- **Vite** - Build tool

---

Built with ❤️ for the Stock Analyzer Platform
