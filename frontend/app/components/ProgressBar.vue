<script setup lang="ts">
  const { t } = useI18n()

  const props = defineProps<{
    progress: number
    currentQuestion: number
    totalQuestions: number
  }>()

  const progressValue = computed(() => Math.round(props.progress))
</script>

<template>
  <div class="w-full space-y-3">
    <div class="flex justify-between items-center">
      <span class="text-sm font-medium text-slate-600 bg-white/80 backdrop-blur-sm px-3 py-1 rounded-full">
        {{ t('questionnaire.question', { current: currentQuestion, total: totalQuestions }) }}
      </span>
      <span class="text-sm font-bold bg-gradient-to-r from-violet-600 to-purple-600 text-transparent bg-clip-text">
        {{ t('questionnaire.completed', { percent: progressValue }) }}
      </span>
    </div>
    <div class="w-full bg-slate-200/80 backdrop-blur-sm rounded-full h-3 overflow-hidden shadow-inner">
      <div
        class="h-full rounded-full transition-all duration-500 ease-out bg-gradient-to-r from-violet-500 via-purple-500 to-pink-500 shadow-lg shadow-purple-500/30 relative"
        :style="{ width: `${progressValue}%` }"
      >
        <!-- Animated shine effect -->
        <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent animate-shimmer" />
      </div>
    </div>
  </div>
</template>
