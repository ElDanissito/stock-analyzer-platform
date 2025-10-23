<template>
  <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow duration-200 p-6">
    <div class="flex justify-between items-start mb-4">
      <div>
        <h3 class="text-xl font-bold text-gray-900">{{ stock.ticker }}</h3>
        <p class="text-sm text-gray-600 mt-1">{{ stock.company }}</p>
      </div>
      <span
        class="px-3 py-1 rounded-full text-xs font-semibold"
        :class="getRatingClass(stock.rating_to)"
      >
        {{ stock.rating_to }}
      </span>
    </div>

    <div class="grid grid-cols-2 gap-4 mb-4">
      <div>
        <p class="text-xs text-gray-500">Target From</p>
        <p class="text-lg font-semibold text-gray-900">{{ stock.target_from }}</p>
      </div>
      <div>
        <p class="text-xs text-gray-500">Target To</p>
        <p class="text-lg font-semibold" :class="getTargetChangeClass()">
          {{ stock.target_to }}
        </p>
      </div>
    </div>

    <div class="mb-4">
      <p class="text-xs text-gray-500">Action</p>
      <p class="text-sm font-medium text-gray-900">{{ stock.action }}</p>
    </div>

    <div class="flex items-center justify-between text-sm">
      <span class="text-gray-600">{{ stock.brokerage || 'N/A' }}</span>
      <span class="text-gray-400">{{ formatDate(stock.time) }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Stock } from '@/types/stock'
import { computed } from 'vue'

const props = defineProps<{
  stock: Stock
}>()

const getRatingClass = (rating: string) => {
  const ratingLower = rating.toLowerCase()
  if (ratingLower.includes('buy') || ratingLower.includes('outperform')) {
    return 'bg-green-100 text-green-800'
  } else if (ratingLower.includes('hold') || ratingLower.includes('neutral')) {
    return 'bg-yellow-100 text-yellow-800'
  } else if (ratingLower.includes('sell') || ratingLower.includes('underperform')) {
    return 'bg-red-100 text-red-800'
  }
  return 'bg-gray-100 text-gray-800'
}

const getTargetChangeClass = () => {
  const from = parseFloat(props.stock.target_from.replace('$', ''))
  const to = parseFloat(props.stock.target_to.replace('$', ''))
  if (to > from) return 'text-green-600'
  if (to < from) return 'text-red-600'
  return 'text-gray-900'
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}
</script>
