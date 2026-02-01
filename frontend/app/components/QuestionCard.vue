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
  <div class="w-full max-w-2xl">
    <!-- Card with gradient border effect -->
    <div class="relative p-[3px] rounded-3xl bg-gradient-to-r from-violet-500 via-purple-500 to-pink-500 shadow-xl shadow-purple-500/20">
      <div class="bg-white rounded-[21px] p-6 sm:p-8">
        <!-- Question text -->
        <h2 class="text-lg sm:text-xl font-bold text-slate-800 leading-relaxed mb-6">
          {{ questionText }}
        </h2>

        <!-- Answer options -->
        <div class="space-y-3">
          <button
            v-for="(option, index) in options"
            :key="option.value"
            class="w-full p-4 text-left rounded-xl border-2 transition-all duration-200 hover:-translate-y-0.5 hover:shadow-md"
            :class="[
              currentAnswer === option.value
                ? 'border-transparent bg-gradient-to-r from-violet-500 to-purple-500 text-white shadow-lg shadow-purple-500/30'
                : 'border-slate-200 bg-white hover:border-purple-300 hover:bg-purple-50',
            ]"
            @click="selectAnswer(option.value)"
          >
            <div class="flex items-center gap-4">
              <div
                class="w-8 h-8 rounded-full flex items-center justify-center shrink-0 font-bold text-sm transition-all"
                :class="[
                  currentAnswer === option.value
                    ? 'bg-white/20 text-white'
                    : 'bg-gradient-to-br from-slate-100 to-slate-200 text-slate-500',
                ]"
              >
                {{ index + 1 }}
              </div>
              <span
                class="text-base"
                :class="currentAnswer === option.value ? 'font-semibold text-white' : 'text-slate-700'"
              >
                {{ option.label }}
              </span>
              <UIcon
                v-if="currentAnswer === option.value"
                name="i-lucide-check-circle"
                class="w-5 h-5 text-white ml-auto"
              />
            </div>
          </button>
        </div>

        <!-- Footer hint -->
        <p class="text-sm text-slate-400 text-center mt-6 flex items-center justify-center gap-2">
          <UIcon name="i-lucide-mouse-pointer-click" class="w-4 h-4" />
          {{ t('questionnaire.selectOption') }}
        </p>
      </div>
    </div>
  </div>
</template>
