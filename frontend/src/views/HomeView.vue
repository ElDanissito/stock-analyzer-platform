<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-4">Stock Analysis Dashboard</h1>
      <div class="flex flex-col sm:flex-row gap-4">
        <!-- Search Bar -->
        <div class="flex-1">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search by ticker or company name..."
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
            @input="handleSearch"
          />
        </div>
        
        <!-- Sync Controls -->
        <div class="flex gap-2">
          <div class="flex items-center">
            <label for="syncPages" class="text-sm font-medium text-gray-700 mr-2 whitespace-nowrap">
              Pages:
            </label>
            <input
              id="syncPages"
              v-model.number="syncPages"
              type="number"
              min="1"
              max="100"
              placeholder="20"
              class="w-20 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
            />
          </div>
          
          <!-- Sync Button -->
          <button
            @click="handleSync"
            :disabled="stockStore.syncing"
            class="px-6 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:bg-gray-400 disabled:cursor-not-allowed transition-colors duration-200 flex items-center justify-center space-x-2"
          >
            <span v-if="stockStore.syncing">🔄 Syncing...</span>
            <span v-else>🔄 Sync Data</span>
          </button>
        </div>
      </div>
      
      <!-- Sync Info -->
      <div class="mt-2 text-sm text-gray-600">
        <p>
          💡 Each page contains ~10 stocks. Default: 20 pages (~200 stocks). Max: 100 pages (~1000 stocks).
        </p>
      </div>
    </div>

    <!-- Stats -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-primary-100 rounded-md p-3">
            <span class="text-2xl">📊</span>
          </div>
          <div class="ml-5">
            <p class="text-sm font-medium text-gray-500">Total Stocks</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stockStore.total }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-green-100 rounded-md p-3">
            <span class="text-2xl">📈</span>
          </div>
          <div class="ml-5">
            <p class="text-sm font-medium text-gray-500">Displayed</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stockStore.stocks.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-yellow-100 rounded-md p-3">
            <span class="text-2xl">⭐</span>
          </div>
          <div class="ml-5">
            <p class="text-sm font-medium text-gray-500">Status</p>
            <p class="text-sm font-semibold" :class="stockStore.loading ? 'text-yellow-600' : 'text-green-600'">
              {{ stockStore.loading ? 'Loading...' : 'Ready' }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Error Message -->
    <div v-if="stockStore.error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-6">
      {{ stockStore.error }}
    </div>

    <!-- Loading Spinner -->
    <div v-if="stockStore.loading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600"></div>
    </div>

    <!-- Empty State -->
    <div v-else-if="!stockStore.hasStocks" class="text-center py-12">
      <span class="text-6xl mb-4 block">📭</span>
      <h3 class="text-xl font-semibold text-gray-900 mb-2">No stocks found</h3>
      <p class="text-gray-600 mb-4">Click the "Sync Data" button to fetch stocks from the API</p>
    </div>

    <!-- Stocks Grid -->
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <StockCard
        v-for="stock in stockStore.stocks"
        :key="stock.id"
        :stock="stock"
      />
    </div>

    <!-- Pagination -->
    <div v-if="stockStore.hasStocks && stockStore.total > limit" class="flex justify-center items-center space-x-4 mt-8">
      <button
        @click="previousPage"
        :disabled="offset === 0"
        class="px-4 py-2 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        Previous
      </button>
      <span class="text-gray-700">
        Page {{ currentPage }} of {{ totalPages }}
      </span>
      <button
        @click="nextPage"
        :disabled="offset + limit >= stockStore.total"
        class="px-4 py-2 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useStockStore } from '@/stores/stockStore'
import StockCard from '@/components/StockCard.vue'

const stockStore = useStockStore()
const searchQuery = ref('')
const limit = ref(50)
const offset = ref(0)
const syncPages = ref(20) // Default: 20 pages (~200 stocks)

const currentPage = computed(() => Math.floor(offset.value / limit.value) + 1)
const totalPages = computed(() => Math.ceil(stockStore.total / limit.value))

let searchTimeout: ReturnType<typeof setTimeout>

const handleSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    offset.value = 0
    if (searchQuery.value.trim()) {
      stockStore.searchStocks(searchQuery.value)
    } else {
      stockStore.fetchStocks(limit.value, offset.value)
    }
  }, 500)
}

const handleSync = async () => {
  // Validate and clamp the value
  const pages = Math.max(1, Math.min(100, syncPages.value || 20))
  await stockStore.syncStocks(pages)
}

const nextPage = () => {
  offset.value += limit.value
  stockStore.fetchStocks(limit.value, offset.value)
}

const previousPage = () => {
  offset.value = Math.max(0, offset.value - limit.value)
  stockStore.fetchStocks(limit.value, offset.value)
}

onMounted(() => {
  stockStore.fetchStocks(limit.value, offset.value)
})
</script>
