<template>
  <div class="bg-gradient-to-br from-primary-500 to-primary-700 rounded-lg shadow-lg p-6 text-white">
    <div class="flex justify-between items-start mb-4">
      <div class="flex-1">
        <div class="flex items-center space-x-2 mb-2">
          <h3 class="text-2xl font-bold">{{ recommendation.stock.ticker }}</h3>
          <span class="bg-white bg-opacity-20 px-2 py-1 rounded text-sm font-semibold">
            Score: {{ recommendation.score.toFixed(1) }}
          </span>
        </div>
        <p class="text-sm opacity-90">{{ recommendation.stock.company }}</p>
      </div>
      <div class="text-right">
        <div class="text-3xl font-bold">
          {{ getScoreEmoji(recommendation.score) }}
        </div>
      </div>
    </div>

    <div class="bg-white bg-opacity-10 rounded p-3 mb-4">
      <p class="text-xs opacity-75 mb-1">Reason</p>
      <p class="text-sm font-medium">{{ recommendation.reason }}</p>
    </div>

    <div class="grid grid-cols-2 gap-3 mb-3">
      <div class="bg-white bg-opacity-10 rounded p-2">
        <p class="text-xs opacity-75">Target</p>
        <p class="text-sm font-semibold">
          {{ recommendation.stock.target_from }} → {{ recommendation.stock.target_to }}
        </p>
      </div>
      <div class="bg-white bg-opacity-10 rounded p-2">
        <p class="text-xs opacity-75">Rating</p>
        <p class="text-sm font-semibold">{{ recommendation.stock.rating_to }}</p>
      </div>
    </div>

    <div class="flex items-center justify-between text-xs opacity-75">
      <span>{{ recommendation.stock.brokerage || 'N/A' }}</span>
      <span>{{ formatDate(recommendation.stock.time) }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { StockRecommendation } from '@/types/stock'

defineProps<{
  recommendation: StockRecommendation
}>()

const getScoreEmoji = (score: number) => {
  if (score >= 80) return '🌟'
  if (score >= 60) return '⭐'
  if (score >= 40) return '✨'
  return '💫'
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}
</script>
