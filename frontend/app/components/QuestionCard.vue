<script setup lang="ts">
  import type { Question } from '~/data/questions'

  const { t } = useI18n()

  const props = defineProps<{
    question: Question
    currentAnswer?: number
  }>()

  const emit = defineEmits<{
    answer: [value: number]
  }>()

  const options = computed(() => [
    { value: 1, label: t('likert.1') },
    { value: 2, label: t('likert.2') },
    { value: 3, label: t('likert.3') },
    { value: 4, label: t('likert.4') },
    { value: 5, label: t('likert.5') }
  ])

  const questionText = computed(() => t(`questions.${props.question.id}`))

  const selectAnswer = (value: number) => {
    emit('answer', value)
  }
</script>

<template>
  <UCard class="w-full max-w-2xl">
    <template #header>
      <h2 class="text-lg sm:text-xl font-semibold text-gray-900 dark:text-white">
        {{ questionText }}
      </h2>
    </template>

    <div class="space-y-2 sm:space-y-3">
      <button
        v-for="option in options"
        :key="option.value"
        class="w-full p-3 sm:p-4 text-left rounded-lg border-2 transition-all duration-200"
        :class="[
          currentAnswer === option.value
            ? 'border-primary-500 bg-primary-50 dark:bg-primary-900/20'
            : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600',
        ]"
        @click="selectAnswer(option.value)"
      >
        <div class="flex items-center gap-2 sm:gap-3">
          <div
            class="w-5 h-5 sm:w-6 sm:h-6 rounded-full border-2 flex items-center justify-center shrink-0"
            :class="[
              currentAnswer === option.value
                ? 'border-primary-500 bg-primary-500'
                : 'border-gray-300 dark:border-gray-600',
            ]"
          >
            <UIcon
              v-if="currentAnswer === option.value"
              name="i-lucide-check"
              class="w-3 h-3 sm:w-4 sm:h-4 text-white"
            />
          </div>
          <span
            class="text-sm sm:text-base text-gray-700 dark:text-gray-300"
            :class="{ 'font-medium': currentAnswer === option.value }"
          >
            {{ option.label }}
          </span>
        </div>
      </button>
    </div>

    <template #footer>
      <p class="text-xs sm:text-sm text-gray-500 dark:text-gray-400 text-center">
        {{ t('questionnaire.selectOption') }}
      </p>
    </template>
  </UCard>
</template>
