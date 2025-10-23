<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2">🎯 Top Stock Recommendations</h1>
      <p class="text-gray-600">AI-powered analysis of the best investment opportunities</p>
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
    <div v-else-if="!stockStore.hasRecommendations" class="text-center py-12">
      <span class="text-6xl mb-4 block">🤔</span>
      <h3 class="text-xl font-semibold text-gray-900 mb-2">No recommendations available</h3>
      <p class="text-gray-600 mb-4">Sync stock data first to get recommendations</p>
      <router-link
        to="/"
        class="inline-block px-6 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700"
      >
        Go to Dashboard
      </router-link>
    </div>

    <!-- Recommendations Grid -->
    <div v-else>
      <!-- Top 3 Recommendations -->
      <div v-if="topThree.length > 0" class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">🏆 Top 3 Picks</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <RecommendationCard
            v-for="(rec, index) in topThree"
            :key="rec.stock.id"
            :recommendation="rec"
          />
        </div>
      </div>

      <!-- Other Recommendations -->
      <div v-if="otherRecommendations.length > 0">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">📊 Other Strong Candidates</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="rec in otherRecommendations"
            :key="rec.stock.id"
            class="bg-white rounded-lg shadow p-6 hover:shadow-lg transition-shadow duration-200"
          >
            <div class="flex justify-between items-start mb-4">
              <div>
                <h3 class="text-lg font-bold text-gray-900">{{ rec.stock.ticker }}</h3>
                <p class="text-sm text-gray-600">{{ rec.stock.company }}</p>
              </div>
              <span class="bg-primary-100 text-primary-800 px-3 py-1 rounded-full text-sm font-semibold">
                {{ rec.score.toFixed(1) }}
              </span>
            </div>
            <p class="text-sm text-gray-700 mb-3">{{ rec.reason }}</p>
            <div class="flex justify-between text-xs text-gray-500">
              <span>{{ rec.stock.rating_to }}</span>
              <span>{{ formatDate(rec.stock.time) }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Refresh Button -->
      <div class="flex justify-center mt-8">
        <button
          @click="refreshRecommendations"
          :disabled="stockStore.loading"
          class="px-6 py-3 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:bg-gray-400 disabled:cursor-not-allowed transition-colors duration-200"
        >
          <span v-if="stockStore.loading">🔄 Loading...</span>
          <span v-else>🔄 Refresh Recommendations</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useStockStore } from '@/stores/stockStore'
import RecommendationCard from '@/components/RecommendationCard.vue'

const stockStore = useStockStore()

const topThree = computed(() => stockStore.recommendations.slice(0, 3))
const otherRecommendations = computed(() => stockStore.recommendations.slice(3))

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

const refreshRecommendations = () => {
  stockStore.fetchRecommendations(10)
}

onMounted(() => {
  stockStore.fetchRecommendations(10)
})
</script>
