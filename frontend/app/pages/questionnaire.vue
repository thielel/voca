<script setup lang="ts">
  const router = useRouter()
  const { t, locale } = useI18n()
  const localePath = useLocalePath()
  const {
    currentQuestion,
    currentQuestionIndex,
    totalQuestions,
    progress,
    isComplete,
    previousQuestion,
    answerQuestion,
    getCurrentAnswer,
    calculateResults,
    resetQuestionnaire,
  } = useQuestionnaire()

  const isSaving = ref(false)
  const saveError = ref<string | null>(null)

  const handleCancel = () => {
    resetQuestionnaire()
    router.push(localePath('/'))
  }

  watch(isComplete, async (complete) => {
    if (complete) {
      isSaving.value = true
      saveError.value = null
      try {
        // Pass locale for AI interpretation generation in the background
        const result = await calculateResults(locale.value)
        router.push(localePath(`/results/${result.id}`))
      } catch (e) {
        console.error('Failed to save results:', e)
        saveError.value = t('questionnaire.saveError')
        isSaving.value = false
      }
    }
  })
</script>

<template>
  <div class="min-h-screen flex flex-col items-center justify-center px-3 sm:px-4 py-6 sm:py-8 relative">
    <!-- Saving State -->
    <div v-if="isSaving" class="relative w-full max-w-2xl">
      <div class="flex flex-col items-center justify-center py-16">
        <div class="w-20 h-20 sm:w-24 sm:h-24 mb-6 rounded-3xl bg-gradient-to-br from-violet-400 via-purple-500 to-indigo-500 flex items-center justify-center shadow-xl shadow-purple-500/30">
          <UIcon name="i-lucide-loader-2" class="w-10 h-10 sm:w-12 sm:h-12 text-white animate-spin" />
        </div>
        <h2 class="text-2xl sm:text-3xl font-bold text-gray-900 dark:text-white mb-2">
          {{ t('questionnaire.saving') }}
        </h2>
        <p class="text-gray-600 dark:text-gray-400">
          {{ t('questionnaire.savingDescription') }}
        </p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="saveError" class="relative w-full max-w-2xl">
      <UCard class="border-red-200 dark:border-red-800">
        <div class="flex flex-col items-center text-center py-8">
          <div class="w-16 h-16 mb-4 rounded-2xl bg-red-100 dark:bg-red-900/30 flex items-center justify-center">
            <UIcon name="i-lucide-alert-circle" class="w-8 h-8 text-red-500" />
          </div>
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-2">
            {{ t('questionnaire.saveErrorTitle') }}
          </h2>
          <p class="text-gray-600 dark:text-gray-400 mb-6">
            {{ saveError }}
          </p>
          <div class="flex gap-3">
            <UButton variant="outline" @click="handleCancel">
              {{ t('questionnaire.cancel') }}
            </UButton>
            <UButton @click="calculateResults().then(r => router.push(localePath(`/results/${r.id}`)))">
              {{ t('questionnaire.retry') }}
            </UButton>
          </div>
        </div>
      </UCard>
    </div>

    <!-- Questionnaire -->
    <div v-else class="relative w-full max-w-2xl space-y-6 sm:space-y-8">
      <ProgressBar
        :progress="progress"
        :current-question="currentQuestionIndex + 1"
        :total-questions="totalQuestions"
      />

      <Transition name="slide" mode="out-in">
        <QuestionCard
          v-if="currentQuestion"
          :key="currentQuestion.id"
          :question="currentQuestion"
          :current-answer="getCurrentAnswer()"
          @answer="answerQuestion"
        />
      </Transition>

      <div class="flex justify-between gap-2">
        <UButton
          variant="ghost"
          size="sm"
          class="sm:size-md"
          :disabled="currentQuestionIndex === 0"
          @click="previousQuestion"
        >
          <UIcon name="i-lucide-arrow-left" class="w-4 h-4 sm:mr-2" />
          <span class="hidden sm:inline">{{ t('questionnaire.back') }}</span>
        </UButton>

        <UButton variant="ghost" size="sm" class="sm:size-md" @click="handleCancel">
          <UIcon name="i-lucide-x" class="w-4 h-4 sm:mr-2" />
          <span class="hidden sm:inline">{{ t('questionnaire.cancel') }}</span>
        </UButton>
      </div>
    </div>
  </div>
</template>

<style scoped>
  .slide-enter-active,
  .slide-leave-active {
    transition: all 0.3s ease;
  }

  .slide-enter-from {
    opacity: 0;
    transform: translateX(30px);
  }

  .slide-leave-to {
    opacity: 0;
    transform: translateX(-30px);
  }
</style>
