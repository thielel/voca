<script setup lang="ts">
  const router = useRouter()
  const { t } = useI18n()
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
  } = useQuestionnaire()

  watch(isComplete, async (complete) => {
    if (complete) {
      const result = await calculateResults()
      router.push(localePath(`/results/${result.id}`))
    }
  })
</script>

<template>
  <div class="min-h-screen flex flex-col items-center justify-center px-3 sm:px-4 py-6 sm:py-8 relative">
    <div class="relative w-full max-w-2xl space-y-6 sm:space-y-8">
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

        <UButton variant="ghost" size="sm" class="sm:size-md" :to="localePath('/')">
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
