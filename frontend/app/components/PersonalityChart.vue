<script setup lang="ts">
  import type { PersonalityResult } from '~/composables/useQuestionnaire'

  const { t } = useI18n()

  const props = defineProps<{
    result: PersonalityResult
  }>()

  type ChartType = 'bar' | 'radar'
  const chartType = ref<ChartType>('bar')

  // Vibrant trait colors - bold gradients for each personality trait
  const traits = computed(() => [
    {
      key: 'extraversion',
      name: t('traits.extraversion'),
      value: props.result.extraversion,
      icon: 'i-lucide-zap',
      color: 'bg-gradient-to-r from-rose-500 to-pink-500',
      iconBg: 'bg-gradient-to-br from-rose-100 to-pink-200',
      iconColor: 'text-rose-500',
      strokeColor: '#f43f5e',
      description: t('traitDescriptions.extraversion')
    },
    {
      key: 'agreeableness',
      name: t('traits.agreeableness'),
      value: props.result.agreeableness,
      icon: 'i-lucide-heart-handshake',
      color: 'bg-gradient-to-r from-emerald-500 to-teal-500',
      iconBg: 'bg-gradient-to-br from-emerald-100 to-teal-200',
      iconColor: 'text-emerald-500',
      strokeColor: '#10b981',
      description: t('traitDescriptions.agreeableness')
    },
    {
      key: 'conscientiousness',
      name: t('traits.conscientiousness'),
      value: props.result.conscientiousness,
      icon: 'i-lucide-target',
      color: 'bg-gradient-to-r from-amber-500 to-orange-500',
      iconBg: 'bg-gradient-to-br from-amber-100 to-orange-200',
      iconColor: 'text-amber-500',
      strokeColor: '#f59e0b',
      description: t('traitDescriptions.conscientiousness')
    },
    {
      key: 'emotionalStability',
      name: t('traits.emotionalStability'),
      value: props.result.emotionalStability,
      icon: 'i-lucide-anchor',
      color: 'bg-gradient-to-r from-sky-500 to-blue-500',
      iconBg: 'bg-gradient-to-br from-sky-100 to-blue-200',
      iconColor: 'text-sky-500',
      strokeColor: '#0ea5e9',
      description: t('traitDescriptions.emotionalStability')
    },
    {
      key: 'openness',
      name: t('traits.openness'),
      value: props.result.openness,
      icon: 'i-lucide-lightbulb',
      color: 'bg-gradient-to-r from-violet-500 to-purple-500',
      iconBg: 'bg-gradient-to-br from-violet-100 to-purple-200',
      iconColor: 'text-violet-500',
      strokeColor: '#8b5cf6',
      description: t('traitDescriptions.openness')
    }
  ])

  const getLevel = (value: number) => {
    if (value >= 75) return t('levels.high')
    if (value >= 50) return t('levels.medium')
    if (value >= 25) return t('levels.low')
    return t('levels.veryLow')
  }

  // Radar chart calculations - responsive sizing
  const isSmallScreen = ref(false)
  
  onMounted(() => {
    const checkScreenSize = () => {
      isSmallScreen.value = window.innerWidth < 640
    }
    checkScreenSize()
    window.addEventListener('resize', checkScreenSize)
    onUnmounted(() => window.removeEventListener('resize', checkScreenSize))
  })
  
  const radarSize = computed(() => isSmallScreen.value ? 320 : 420)
  const radarCenter = computed(() => radarSize.value / 2)
  const radarRadius = computed(() => isSmallScreen.value ? 110 : 150)
  const levels = [25, 50, 75, 100]

  // Calculate point position on the radar for a given index and value (0-100)
  const getRadarPoint = (index: number, value: number) => {
    const angleOffset = -Math.PI / 2 // Start from top
    const angle = angleOffset + (index * 2 * Math.PI) / 5
    const r = (value / 100) * radarRadius.value
    return {
      x: radarCenter.value + r * Math.cos(angle),
      y: radarCenter.value + r * Math.sin(angle)
    }
  }

  // Generate pentagon points for grid levels
  const getLevelPolygon = (level: number) => {
    return Array.from({ length: 5 }, (_, i) => {
      const point = getRadarPoint(i, level)
      return `${point.x},${point.y}`
    }).join(' ')
  }

  // Generate axis lines from center to each vertex
  const getAxisLine = (index: number) => {
    const point = getRadarPoint(index, 100)
    return `M${radarCenter.value},${radarCenter.value} L${point.x},${point.y}`
  }

  // Generate the data polygon path
  const dataPolygon = computed(() => {
    return traits.value
      .map((trait, i) => {
        const point = getRadarPoint(i, trait.value)
        return `${point.x},${point.y}`
      })
      .join(' ')
  })

  // Label positions (slightly outside the radar)
  const getLabelPosition = (index: number) => {
    const labelOffset = isSmallScreen.value ? 125 : 130
    const point = getRadarPoint(index, labelOffset)
    return { x: point.x, y: point.y }
  }

  // Tooltip state for hover interactions
  const hoveredTrait = ref<string | null>(null)
  const tooltipPosition = ref({ x: 0, y: 0 })
  const svgRef = ref<SVGSVGElement | null>(null)

  const showTooltip = (trait: typeof traits.value[0], index: number, event: MouseEvent) => {
    hoveredTrait.value = trait.key
    if (svgRef.value) {
      const svgRect = svgRef.value.getBoundingClientRect()
      const labelPos = getLabelPosition(index)
      tooltipPosition.value = {
        x: svgRect.left + labelPos.x,
        y: svgRect.top + labelPos.y - 10
      }
    }
  }

  const hideTooltip = () => {
    hoveredTrait.value = null
  }

  const getHoveredTrait = computed(() => {
    return traits.value.find(t => t.key === hoveredTrait.value)
  })
</script>

<template>
  <div class="space-y-6">
    <!-- Chart Type Toggle -->
    <div class="flex justify-center">
      <div class="inline-flex rounded-xl bg-slate-100 p-1.5 shadow-inner">
        <button
          class="flex items-center gap-2 px-5 py-2.5 rounded-lg text-sm font-semibold transition-all"
          :class="
            chartType === 'bar'
              ? 'bg-gradient-to-r from-violet-500 to-purple-500 text-white shadow-lg shadow-purple-500/30'
              : 'text-slate-500 hover:text-slate-700 hover:bg-white/50'
          "
          @click="chartType = 'bar'"
        >
          <UIcon name="i-lucide-bar-chart-3" class="w-4 h-4" />
          {{ t('chart.bar') }}
        </button>
        <button
          class="flex items-center gap-2 px-5 py-2.5 rounded-lg text-sm font-semibold transition-all"
          :class="
            chartType === 'radar'
              ? 'bg-gradient-to-r from-violet-500 to-purple-500 text-white shadow-lg shadow-purple-500/30'
              : 'text-slate-500 hover:text-slate-700 hover:bg-white/50'
          "
          @click="chartType = 'radar'"
        >
          <UIcon name="i-lucide-radar" class="w-4 h-4" />
          {{ t('chart.radar') }}
        </button>
      </div>
    </div>

    <!-- Bar Chart View -->
    <div v-if="chartType === 'bar'" class="space-y-5">
      <div v-for="trait in traits" :key="trait.key" class="p-4 bg-white rounded-2xl shadow-md shadow-slate-100 border border-slate-100">
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-xl flex items-center justify-center shadow-sm" :class="trait.iconBg">
              <UIcon :name="trait.icon" :class="['w-5 h-5', trait.iconColor]" />
            </div>
            <div>
              <span class="font-bold text-slate-800 block">
                {{ trait.name }}
              </span>
              <span class="text-xs text-slate-500">
                {{ getLevel(trait.value) }}
              </span>
            </div>
          </div>
          <div class="text-2xl font-bold text-slate-800">
            {{ trait.value }}%
          </div>
        </div>

        <div class="w-full bg-slate-100 rounded-full h-3 overflow-hidden">
          <div
            class="h-full rounded-full transition-all duration-1000 ease-out shadow-sm"
            :class="trait.color"
            :style="{ width: `${trait.value}%` }"
          />
        </div>

        <p class="text-sm text-slate-500 mt-3 leading-relaxed">
          {{ trait.description }}
        </p>
      </div>
    </div>

    <!-- Radar Chart View -->
    <div v-else class="flex flex-col items-center space-y-4 sm:space-y-6 relative">
      <svg ref="svgRef" :width="radarSize" :height="radarSize" class="overflow-visible max-w-full">
        <!-- Background grid levels -->
        <polygon
          v-for="level in levels"
          :key="level"
          :points="getLevelPolygon(level)"
          fill="none"
          class="stroke-gray-200 dark:stroke-gray-700"
          stroke-width="1"
        />

        <!-- Axis lines -->
        <path
          v-for="i in 5"
          :key="`axis-${i}`"
          :d="getAxisLine(i - 1)"
          class="stroke-gray-200 dark:stroke-gray-700"
          stroke-width="1"
        />

        <!-- Data polygon with gradient fill -->
        <defs>
          <linearGradient id="radarGradient" x1="0%" y1="0%" x2="100%" y2="100%">
            <stop offset="0%" style="stop-color: #6BB8E4; stop-opacity: 0.5" />
            <stop offset="50%" style="stop-color: #4ABFBF; stop-opacity: 0.5" />
            <stop offset="100%" style="stop-color: #A8E6CF; stop-opacity: 0.5" />
          </linearGradient>
        </defs>
        <polygon
          :points="dataPolygon"
          fill="url(#radarGradient)"
          class="stroke-sky-500"
          stroke-width="2"
          stroke-linejoin="round"
        />

        <!-- Data points -->
        <circle
          v-for="(trait, i) in traits"
          :key="`point-${trait.key}`"
          :cx="getRadarPoint(i, trait.value).x"
          :cy="getRadarPoint(i, trait.value).y"
          :r="isSmallScreen ? 6 : 7"
          :fill="trait.strokeColor"
          class="stroke-white dark:stroke-gray-900"
          stroke-width="2"
        />

        <!-- Labels with hover interaction -->
        <g
          v-for="(trait, i) in traits"
          :key="`label-${trait.key}`"
          class="cursor-pointer"
          @mouseenter="showTooltip(trait, i, $event)"
          @mouseleave="hideTooltip"
        >
          <!-- Invisible hit area for easier hovering -->
          <rect
            :x="getLabelPosition(i).x - 50"
            :y="getLabelPosition(i).y - 12"
            width="100"
            :height="isSmallScreen ? 32 : 40"
            fill="transparent"
          />
          <text
            :x="getLabelPosition(i).x"
            :y="getLabelPosition(i).y"
            text-anchor="middle"
            dominant-baseline="middle"
            :class="[
              isSmallScreen ? 'text-xs' : 'text-sm',
              hoveredTrait === trait.key ? 'fill-sky-600 dark:fill-sky-400' : 'fill-gray-700 dark:fill-gray-300'
            ]"
            class="font-medium transition-colors"
          >
            {{ trait.name }}
          </text>
          <text
            :x="getLabelPosition(i).x"
            :y="getLabelPosition(i).y + (isSmallScreen ? 14 : 18)"
            text-anchor="middle"
            dominant-baseline="middle"
            :class="[
              isSmallScreen ? 'text-xs' : 'text-sm',
              hoveredTrait === trait.key ? 'fill-sky-700 dark:fill-sky-300' : 'fill-gray-900 dark:fill-white'
            ]"
            class="font-semibold transition-colors"
          >
            {{ trait.value }}%
          </text>
        </g>
      </svg>

      <!-- Tooltip -->
      <Teleport to="body">
        <Transition
          enter-active-class="transition duration-150 ease-out"
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
          leave-active-class="transition duration-100 ease-in"
          leave-from-class="opacity-100 scale-100"
          leave-to-class="opacity-0 scale-95"
        >
          <div
            v-if="hoveredTrait && getHoveredTrait"
            class="fixed z-50 max-w-xs px-3 py-2 text-sm bg-white dark:bg-gray-800 rounded-lg shadow-lg border border-gray-200 dark:border-gray-700 pointer-events-none"
            :style="{
              left: `${tooltipPosition.x}px`,
              top: `${tooltipPosition.y}px`,
              transform: 'translate(-50%, -100%)'
            }"
          >
            <div class="flex items-center gap-2 mb-1">
              <div
                class="w-2 h-2 rounded-full shrink-0"
                :style="{ backgroundColor: getHoveredTrait.strokeColor }"
              />
              <span class="font-medium text-gray-900 dark:text-white">
                {{ getHoveredTrait.name }}
              </span>
              <span class="text-xs text-gray-500 dark:text-gray-400">
                {{ getLevel(getHoveredTrait.value) }}
              </span>
            </div>
            <p class="text-xs text-gray-600 dark:text-gray-300 leading-relaxed">
              {{ getHoveredTrait.description }}
            </p>
            <!-- Tooltip arrow -->
            <div class="absolute left-1/2 -bottom-1.5 -translate-x-1/2 w-3 h-3 rotate-45 bg-white dark:bg-gray-800 border-r border-b border-gray-200 dark:border-gray-700" />
          </div>
        </Transition>
      </Teleport>
    </div>
  </div>
</template>
