<script setup lang="ts">
const { t } = useI18n()
const router = useRouter()
const { results, isLoading, error, fetchResults, formatDate, truncateId } = useAdmin()

onMounted(() => {
  fetchResults()
})

const traitColor = (value: number) => {
  if (value >= 70) return 'text-green-600 dark:text-green-400'
  if (value >= 40) return 'text-yellow-600 dark:text-yellow-400'
  return 'text-red-600 dark:text-red-400'
}

const viewResult = (id: string) => {
  router.push(`/results/${id}`)
}
</script>

<template>
  <div class="min-h-screen py-8 px-4">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 dark:text-white">
              {{ t('admin.dashboard') }}
            </h1>
            <p class="mt-2 text-gray-600 dark:text-gray-400">
              {{ t('admin.overview') }}
            </p>
          </div>
          <div class="flex gap-2 sm:gap-3">
            <UButton
              icon="i-lucide-refresh-cw"
              variant="outline"
              :loading="isLoading"
              size="sm"
              class="sm:size-md"
              @click="fetchResults"
            >
              <span class="hidden sm:inline">{{ t('admin.refresh') }}</span>
            </UButton>
            <UButton
              to="/"
              icon="i-lucide-home"
              variant="ghost"
              size="sm"
              class="sm:size-md"
            >
              <span class="hidden sm:inline">{{ t('admin.home') }}</span>
            </UButton>
          </div>
        </div>
      </div>

      <!-- Error State -->
      <UCard v-if="error" class="mb-6 border-red-200 dark:border-red-800">
        <div class="flex items-center gap-3 text-red-600 dark:text-red-400">
          <UIcon name="i-lucide-alert-circle" class="w-5 h-5" />
          <span>{{ error }}</span>
        </div>
      </UCard>

      <!-- Loading State -->
      <div v-if="isLoading && results.length === 0" class="flex justify-center py-12">
        <div class="text-center">
          <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-500 mb-4" />
          <p class="text-gray-600 dark:text-gray-400">{{ t('admin.loadingResults') }}</p>
        </div>
      </div>

      <!-- Empty State -->
      <UCard v-else-if="!isLoading && results.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-inbox" class="w-12 h-12 text-gray-400 mx-auto mb-4" />
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">
          {{ t('admin.noResults') }}
        </h3>
        <p class="text-gray-600 dark:text-gray-400 mb-4">
          {{ t('admin.noResultsDescription') }}
        </p>
        <p class="text-sm text-gray-500 dark:text-gray-500">
          {{ t('admin.seedHint', { command: 'make seed' }) }}
        </p>
      </UCard>

      <!-- Results Table -->
      <UCard v-else>
        <template #header>
          <div class="flex items-center justify-between">
            <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('admin.testResults') }}
            </h2>
            <span class="text-sm text-gray-500 dark:text-gray-400">
              {{ t('admin.resultCount', { count: results.length }, results.length) }}
            </span>
          </div>
        </template>

        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b border-gray-200 dark:border-gray-700">
                <th class="text-left py-3 px-4 font-medium text-gray-600 dark:text-gray-400">ID</th>
                <th class="text-left py-3 px-4 font-medium text-gray-600 dark:text-gray-400">Session</th>
                <th class="text-center py-3 px-4 font-medium text-gray-600 dark:text-gray-400">E</th>
                <th class="text-center py-3 px-4 font-medium text-gray-600 dark:text-gray-400">A</th>
                <th class="text-center py-3 px-4 font-medium text-gray-600 dark:text-gray-400">C</th>
                <th class="text-center py-3 px-4 font-medium text-gray-600 dark:text-gray-400">ES</th>
                <th class="text-center py-3 px-4 font-medium text-gray-600 dark:text-gray-400">O</th>
                <th class="text-left py-3 px-4 font-medium text-gray-600 dark:text-gray-400">{{ t('admin.date') }}</th>
                <th class="w-10"></th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="result in results"
                :key="result.id"
                class="border-b border-gray-100 dark:border-gray-800 hover:bg-gray-50 dark:hover:bg-gray-800/50 transition-colors cursor-pointer"
                @click="viewResult(result.id)"
              >
                <td class="py-3 px-4 font-mono text-xs text-gray-600 dark:text-gray-400">
                  {{ truncateId(result.id) }}
                </td>
                <td class="py-3 px-4 font-mono text-xs text-gray-600 dark:text-gray-400">
                  {{ truncateId(result.session_id) }}
                </td>
                <td class="py-3 px-4 text-center font-medium" :class="traitColor(result.extraversion)">
                  {{ Math.round(result.extraversion) }}%
                </td>
                <td class="py-3 px-4 text-center font-medium" :class="traitColor(result.agreeableness)">
                  {{ Math.round(result.agreeableness) }}%
                </td>
                <td class="py-3 px-4 text-center font-medium" :class="traitColor(result.conscientiousness)">
                  {{ Math.round(result.conscientiousness) }}%
                </td>
                <td class="py-3 px-4 text-center font-medium" :class="traitColor(result.emotional_stability)">
                  {{ Math.round(result.emotional_stability) }}%
                </td>
                <td class="py-3 px-4 text-center font-medium" :class="traitColor(result.openness)">
                  {{ Math.round(result.openness) }}%
                </td>
                <td class="py-3 px-4 text-gray-600 dark:text-gray-400 whitespace-nowrap">
                  {{ formatDate(result.created_at) }}
                </td>
                <td class="py-3 px-4">
                  <UIcon name="i-lucide-chevron-right" class="w-4 h-4 text-gray-400" />
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <template #footer>
          <div class="flex flex-wrap items-center gap-3 sm:gap-6 text-xs text-gray-500 dark:text-gray-400">
            <div class="flex items-center gap-2">
              <span class="font-medium">{{ t('admin.legend') }}:</span>
            </div>
            <div class="flex items-center gap-1">
              <span class="w-3 h-3 rounded-full bg-green-500 shrink-0" />
              <span>{{ t('admin.high') }}</span>
            </div>
            <div class="flex items-center gap-1">
              <span class="w-3 h-3 rounded-full bg-yellow-500 shrink-0" />
              <span>{{ t('admin.medium') }}</span>
            </div>
            <div class="flex items-center gap-1">
              <span class="w-3 h-3 rounded-full bg-red-500 shrink-0" />
              <span>{{ t('admin.low') }}</span>
            </div>
          </div>
        </template>
      </UCard>

      <!-- Trait Legend -->
      <div class="mt-6 text-center text-xs sm:text-sm text-gray-500 dark:text-gray-400">
        <span class="font-medium block sm:inline mb-1 sm:mb-0">{{ t('admin.personalityDimensions') }}:</span>
        <span class="block sm:inline">{{ t('admin.dimensionsLegend') }}</span>
      </div>
    </div>
  </div>
</template>
