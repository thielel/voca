<script setup lang="ts">
import type { PersonalityResult } from '~/data/questions'

const { t, locale } = useI18n()
const localePath = useLocalePath()
const router = useRouter()
const { resetQuestionnaire } = useQuestionnaire()

// API response type with snake_case fields
interface ApiResultResponse {
  id: string
  session_id: string
  extraversion: number
  agreeableness: number
  conscientiousness: number
  emotional_stability: number
  openness: number
  created_at: string
  interpretations?: Record<string, string>
}

// Trait key mapping from API to frontend
const traitKeyMap: Record<string, string> = {
  extraversion: 'extraversion',
  agreeableness: 'agreeableness',
  conscientiousness: 'conscientiousness',
  emotional_stability: 'emotionalStability',
  openness: 'openness',
}

const route = useRoute()

const result = ref<PersonalityResult | null>(null)
const interpretations = ref<Record<string, string>>({})
const isLoading = ref(true)
const isRegenerating = ref(false)
const error = ref<string | null>(null)

const hasInterpretations = computed(() => Object.keys(interpretations.value).length > 0)

const fetchResult = async () => {
  isLoading.value = true
  error.value = null

  try {
    const id = route.params.id as string
    const data = await $fetch<ApiResultResponse>(`http://localhost:8080/api/results/${id}`)

    // Convert snake_case to camelCase for PersonalityChart compatibility
    result.value = {
      id: data.id,
      extraversion: data.extraversion,
      agreeableness: data.agreeableness,
      conscientiousness: data.conscientiousness,
      emotionalStability: data.emotional_stability,
      openness: data.openness,
      createdAt: data.created_at,
    }

    // Map interpretations from API (snake_case) to frontend (camelCase)
    if (data.interpretations) {
      const mapped: Record<string, string> = {}
      for (const [key, value] of Object.entries(data.interpretations)) {
        const mappedKey = traitKeyMap[key] || key
        mapped[mappedKey] = value
      }
      interpretations.value = mapped
    } else {
      interpretations.value = {}
    }
  } catch (e) {
    error.value = t('results.failedToLoad')
    console.error('Failed to fetch result:', e)
  } finally {
    isLoading.value = false
  }
}

const regenerateInterpretations = async () => {
  if (!result.value?.id) return

  isRegenerating.value = true
  error.value = null

  try {
    const data = await $fetch<ApiResultResponse>(
      `http://localhost:8080/api/results/${result.value.id}/regenerate`,
      {
        method: 'POST',
        body: { language: locale.value }
      }
    )

    // Map interpretations from API (snake_case) to frontend (camelCase)
    if (data.interpretations) {
      const mapped: Record<string, string> = {}
      for (const [key, value] of Object.entries(data.interpretations)) {
        const mappedKey = traitKeyMap[key] || key
        mapped[mappedKey] = value
      }
      interpretations.value = mapped
    }
  } catch (e) {
    error.value = t('results.failedToGenerateInterpretations')
    console.error('Failed to regenerate interpretations:', e)
  } finally {
    isRegenerating.value = false
  }
}

onMounted(() => {
  fetchResult()
})

const startOver = () => {
  resetQuestionnaire()
  router.push(localePath('/'))
}

const getLevel = (value: number): 'high' | 'medium' | 'low' => {
  if (value >= 70) return 'high'
  if (value >= 40) return 'medium'
  return 'low'
}

const getLevelLabel = (value: number) => {
  const level = getLevel(value)
  if (level === 'high') return t('results.levelHigh')
  if (level === 'medium') return t('results.levelMedium')
  return t('results.levelLow')
}

const getLevelColor = (value: number) => {
  const level = getLevel(value)
  if (level === 'high') return 'text-green-600 dark:text-green-400'
  if (level === 'medium') return 'text-yellow-600 dark:text-yellow-400'
  return 'text-red-600 dark:text-red-400'
}

// Trait metadata (icon only, names come from i18n)
const traitIcons = {
  extraversion: 'i-lucide-smile',
  agreeableness: 'i-lucide-heart-handshake',
  conscientiousness: 'i-lucide-check-circle',
  emotionalStability: 'i-lucide-shield',
  openness: 'i-lucide-lightbulb',
}

const getTraitName = (key: string) => {
  return t(`traits.${key}`)
}

const traits = computed(() => {
  if (!result.value) return []
  return [
    { key: 'extraversion' as const, value: result.value.extraversion },
    { key: 'agreeableness' as const, value: result.value.agreeableness },
    { key: 'conscientiousness' as const, value: result.value.conscientiousness },
    { key: 'emotionalStability' as const, value: result.value.emotionalStability },
    { key: 'openness' as const, value: result.value.openness },
  ]
})

const getInterpretation = (traitKey: string) => {
  return interpretations.value[traitKey] || null
}

// Format AI interpretation text (simple markdown-like parsing)
const formatInterpretation = (text: string | null): string => {
  if (!text) return ''

  // Convert markdown headings to HTML
  let formatted = text
    // ## Heading -> <h4>
    .replace(/^## (.+)$/gm, '<h4 class="font-semibold text-gray-900 dark:text-white mt-4 mb-2 first:mt-0">$1</h4>')
    // Bold text **text**
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    // Line breaks
    .replace(/\n\n/g, '</p><p class="mb-2">')
    .replace(/\n/g, '<br>')

  // Wrap in paragraph if not starting with heading
  if (!formatted.startsWith('<h4')) {
    formatted = `<p class="mb-2">${formatted}</p>`
  } else {
    // Ensure content after headings is wrapped in paragraphs
    formatted = formatted.replace(/<\/h4>(?!<h4|<p|$)/g, '</h4><p class="mb-2">')
  }

  return formatted
}
</script>

<template>
  <div class="min-h-screen py-8 px-4">
    <div class="max-w-4xl mx-auto">
      <!-- Header -->
      <div class="mb-6 sm:mb-8 text-center">
        <UIcon
          name="i-lucide-sparkles"
          class="w-10 h-10 sm:w-12 sm:h-12 mx-auto text-primary-500 mb-4"
        />
        <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 dark:text-white">
          {{ t('results.title') }}
        </h1>
        <p class="mt-2 text-sm sm:text-base text-gray-600 dark:text-gray-400">
          {{ t('results.subtitle') }}
        </p>
      </div>

      <!-- Error State -->
      <UCard v-if="error && !isLoading" class="mb-6 border-red-200 dark:border-red-800">
        <div class="flex items-center gap-3 text-red-600 dark:text-red-400">
          <UIcon name="i-lucide-alert-circle" class="w-5 h-5" />
          <span>{{ error }}</span>
        </div>
        <div class="flex gap-3 mt-4">
          <UButton variant="outline" @click="fetchResult">
            {{ t('results.retry') }}
          </UButton>
          <UButton :to="localePath('/')">
            {{ t('results.home') }}
          </UButton>
        </div>
      </UCard>

      <!-- Loading State -->
      <div v-else-if="isLoading" class="flex justify-center py-12">
        <div class="text-center">
          <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-500 mb-4" />
          <p class="text-gray-600 dark:text-gray-400">{{ t('results.loading') }}</p>
        </div>
      </div>

      <!-- Result Content -->
      <template v-else-if="result">
        <!-- Personality Chart -->
        <UCard class="mb-6">
          <template #header>
            <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('results.traitValues') }}
            </h2>
          </template>

          <PersonalityChart :result="result" />
        </UCard>

        <!-- Detailed Interpretations -->
        <UCard>
          <template #header>
            <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
              <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
                {{ t('results.detailedAnalysis') }}
              </h2>
              <UButton
                :loading="isRegenerating"
                :disabled="isRegenerating"
                icon="i-lucide-sparkles"
                variant="outline"
                size="sm"
                class="self-start sm:self-auto"
                @click="regenerateInterpretations"
              >
                <span class="hidden sm:inline">{{ hasInterpretations ? t('results.regenerate') : t('results.generateAiInterpretation') }}</span>
                <span class="sm:hidden">{{ hasInterpretations ? t('results.regenerate') : t('results.generate') }}</span>
              </UButton>
            </div>
          </template>

          <!-- Regenerating Indicator -->
          <div v-if="isRegenerating" class="py-12">
            <div class="flex flex-col items-center justify-center gap-4">
              <UIcon name="i-lucide-loader-2" class="w-10 h-10 animate-spin text-primary-500" />
              <div class="text-center">
                <p class="font-medium text-gray-900 dark:text-white">{{ t('results.generatingInterpretations') }}</p>
                <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">
                  {{ t('results.generatingTime') }}
                </p>
              </div>
              <div class="w-64 bg-gray-200 dark:bg-gray-700 rounded-full h-2 overflow-hidden">
                <div class="bg-primary-500 h-2 rounded-full animate-pulse w-full"></div>
              </div>
            </div>
          </div>

          <!-- No Interpretations State -->
          <div v-else-if="!hasInterpretations" class="py-8">
            <div class="flex flex-col items-center justify-center gap-4 text-center">
              <UIcon name="i-lucide-file-question" class="w-12 h-12 text-gray-400" />
              <div>
                <p class="font-medium text-gray-900 dark:text-white">{{ t('results.noInterpretations') }}</p>
                <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">
                  {{ t('results.noInterpretationsDescription') }}
                </p>
              </div>
              <UButton
                icon="i-lucide-sparkles"
                @click="regenerateInterpretations"
              >
                {{ t('results.generateNow') }}
              </UButton>
            </div>
          </div>

          <!-- Interpretations Content -->
          <div v-else class="space-y-6 sm:space-y-8">
            <div
              v-for="trait in traits"
              :key="trait.key"
              class="border-b border-gray-200 dark:border-gray-700 pb-6 last:border-0 last:pb-0"
            >
              <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2 sm:gap-3 mb-3">
                <div class="flex items-center gap-2 sm:gap-3">
                  <UIcon
                    :name="traitIcons[trait.key]"
                    class="w-5 h-5 sm:w-6 sm:h-6 text-gray-600 dark:text-gray-400"
                  />
                  <h3 class="text-base sm:text-lg font-medium text-gray-900 dark:text-white">
                    {{ getTraitName(trait.key) }}
                  </h3>
                </div>
                <div class="flex items-center gap-2 ml-7 sm:ml-0">
                  <span
                    class="text-xs sm:text-sm font-medium px-2 py-1 rounded"
                    :class="getLevelColor(trait.value)"
                  >
                    {{ getLevelLabel(trait.value) }}
                  </span>
                  <span class="text-lg sm:text-xl font-bold text-gray-900 dark:text-white">
                    {{ Math.round(trait.value) }}%
                  </span>
                </div>
              </div>

              <!-- AI-generated interpretation with markdown-like formatting -->
              <div class="text-gray-600 dark:text-gray-300 leading-relaxed prose prose-sm dark:prose-invert max-w-none">
                <div v-html="formatInterpretation(getInterpretation(trait.key))"></div>
              </div>
            </div>
          </div>
        </UCard>

        <!-- Disclaimer -->
        <div class="mt-6 p-4 bg-amber-50 dark:bg-amber-900/20 rounded-lg">
          <div class="flex gap-3">
            <UIcon name="i-lucide-info" class="w-5 h-5 text-amber-600 dark:text-amber-400 flex-shrink-0 mt-0.5" />
            <p class="text-sm text-amber-700 dark:text-amber-300">
              <strong>{{ t('results.remember') }}</strong> {{ t('results.disclaimer') }}
            </p>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex flex-col sm:flex-row gap-4 justify-center mt-8">
          <UButton variant="outline" @click="startOver">
            <UIcon name="i-lucide-refresh-cw" class="w-4 h-4 mr-2" />
            {{ t('results.retake') }}
          </UButton>
          <UButton :to="localePath('/')">
            <UIcon name="i-lucide-home" class="w-4 h-4 mr-2" />
            {{ t('results.home') }}
          </UButton>
        </div>
      </template>
    </div>
  </div>
</template>
