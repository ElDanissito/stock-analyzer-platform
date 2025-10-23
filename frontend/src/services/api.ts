import axios from 'axios'
import type { StocksResponse, Stock, RecommendationsResponse, SyncResponse } from '@/types/stock'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:3000'

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const stocksApi = {
  // Get all stocks with pagination
  getStocks: async (limit = 50, offset = 0): Promise<StocksResponse> => {
    const response = await api.get<StocksResponse>('/api/stocks', {
      params: { limit, offset }
    })
    return response.data
  },

  // Get stock by ID
  getStockById: async (id: string): Promise<Stock> => {
    const response = await api.get<Stock>(`/api/stocks/${id}`)
    return response.data
  },

  // Search stocks
  searchStocks: async (query: string): Promise<StocksResponse> => {
    const response = await api.get<StocksResponse>('/api/stocks/search', {
      params: { q: query }
    })
    return response.data
  },

  // Sync stocks from external API
  syncStocks: async (pages?: number): Promise<SyncResponse> => {
    const response = await api.post<SyncResponse>('/api/sync', {
      pages: pages || 0
    })
    return response.data
  },

  // Get recommendations
  getRecommendations: async (limit = 10): Promise<RecommendationsResponse> => {
    const response = await api.get<RecommendationsResponse>('/api/recommendations', {
      params: { limit }
    })
    return response.data
  },

  // Health check
  healthCheck: async (): Promise<{ status: string }> => {
    const response = await api.get<{ status: string }>('/health')
    return response.data
  }
}
