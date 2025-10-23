import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Stock, StockRecommendation } from '@/types/stock'
import { stocksApi } from '@/services/api'

export const useStockStore = defineStore('stock', () => {
  // State
  const stocks = ref<Stock[]>([])
  const recommendations = ref<StockRecommendation[]>([])
  const currentStock = ref<Stock | null>(null)
  const total = ref(0)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const syncing = ref(false)

  // Getters
  const hasStocks = computed(() => stocks.value.length > 0)
  const hasRecommendations = computed(() => recommendations.value.length > 0)

  // Actions
  const fetchStocks = async (limit = 50, offset = 0) => {
    loading.value = true
    error.value = null
    try {
      const response = await stocksApi.getStocks(limit, offset)
      stocks.value = response.data
      total.value = response.total
    } catch (err: any) {
      error.value = err.message || 'Error fetching stocks'
      console.error('Error fetching stocks:', err)
    } finally {
      loading.value = false
    }
  }

  const searchStocks = async (query: string) => {
    if (!query.trim()) {
      await fetchStocks()
      return
    }

    loading.value = true
    error.value = null
    try {
      const response = await stocksApi.searchStocks(query)
      stocks.value = response.data
      total.value = response.data.length
    } catch (err: any) {
      error.value = err.message || 'Error searching stocks'
      console.error('Error searching stocks:', err)
    } finally {
      loading.value = false
    }
  }

  const fetchStockById = async (id: string) => {
    loading.value = true
    error.value = null
    try {
      currentStock.value = await stocksApi.getStockById(id)
    } catch (err: any) {
      error.value = err.message || 'Error fetching stock details'
      console.error('Error fetching stock:', err)
    } finally {
      loading.value = false
    }
  }

  const fetchRecommendations = async (limit = 10) => {
    loading.value = true
    error.value = null
    try {
      const response = await stocksApi.getRecommendations(limit)
      recommendations.value = response.data
    } catch (err: any) {
      error.value = err.message || 'Error fetching recommendations'
      console.error('Error fetching recommendations:', err)
    } finally {
      loading.value = false
    }
  }

  const syncStocks = async (pages?: number) => {
    syncing.value = true
    error.value = null
    try {
      await stocksApi.syncStocks(pages)
      // Wait a bit and then refresh the stocks list
      setTimeout(() => {
        fetchStocks()
      }, 5000)
    } catch (err: any) {
      error.value = err.message || 'Error syncing stocks'
      console.error('Error syncing stocks:', err)
    } finally {
      syncing.value = false
    }
  }

  return {
    // State
    stocks,
    recommendations,
    currentStock,
    total,
    loading,
    error,
    syncing,
    // Getters
    hasStocks,
    hasRecommendations,
    // Actions
    fetchStocks,
    searchStocks,
    fetchStockById,
    fetchRecommendations,
    syncStocks
  }
})
