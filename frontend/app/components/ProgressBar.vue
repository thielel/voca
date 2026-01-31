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
  <div class="w-full space-y-2">
    <div class="flex justify-between text-sm text-gray-600 dark:text-gray-400">
      <span>{{ t('questionnaire.question', { current: currentQuestion, total: totalQuestions }) }}</span>
      <span>{{ t('questionnaire.completed', { percent: progressValue }) }}</span>
    </div>
    <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2.5 overflow-hidden">
      <div
        class="bg-primary-500 h-full rounded-full transition-all duration-300 ease-out"
        :style="{ width: `${progressValue}%` }"
      />
    </div>
  </div>
</template>
