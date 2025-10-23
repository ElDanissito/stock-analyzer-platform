export interface Stock {
  id: string
  ticker: string
  company: string
  target_from: string
  target_to: string
  action: string
  brokerage: string
  rating_from: string
  rating_to: string
  time: string
  last_updated: string
  created_at: string
}

export interface StocksResponse {
  data: Stock[]
  total: number
  limit: number
  offset: number
}

export interface StockRecommendation {
  stock: Stock
  score: number
  reason: string
}

export interface RecommendationsResponse {
  data: StockRecommendation[]
}

export interface SyncResponse {
  message: string
}
