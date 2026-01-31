import { questions, traits, type Question, type Answer, type PersonalityResult } from '~/data/questions'

export type { Question, Answer, PersonalityResult }

export const useQuestionnaire = () => {
  const currentQuestionIndex = useState<number>('currentQuestionIndex', () => 0)
  const answers = useState<Answer[]>('answers', () => [])
  const result = useState<PersonalityResult | null>('result', () => null)
  const isLoading = useState<boolean>('isLoading', () => false)

  const totalQuestions = computed(() => questions.length)
  const currentQuestion = computed(() => questions[currentQuestionIndex.value])
  const progress = computed(() => ((currentQuestionIndex.value) / totalQuestions.value) * 100)
  const isComplete = computed(() => currentQuestionIndex.value >= totalQuestions.value)

  const answerQuestion = (value: number) => {
    if (!currentQuestion.value) return

    const questionId = currentQuestion.value.id
    const existingIndex = answers.value.findIndex(a => a.questionId === questionId)

    if (existingIndex >= 0) {
      answers.value[existingIndex]!.value = value
    } else {
      answers.value.push({ questionId, value })
    }

    if (currentQuestionIndex.value < totalQuestions.value - 1) {
      currentQuestionIndex.value++
    } else {
      currentQuestionIndex.value = totalQuestions.value
    }
  }

  const previousQuestion = () => {
    if (currentQuestionIndex.value > 0) {
      currentQuestionIndex.value--
    }
  }

  const getCurrentAnswer = () => {
    const questionId = currentQuestion.value?.id
    const answer = answers.value.find(a => a.questionId === questionId)
    return answer?.value
  }

  const calculateResults = async (): Promise<PersonalityResult & { id: string }> => {
    isLoading.value = true

    try {
      const scores: Record<string, number> = {}

      for (const trait of traits) {
        const traitQuestions = questions.filter(q => q.trait === trait)
        let total = 0

        for (const question of traitQuestions) {
          const answer = answers.value.find(a => a.questionId === question.id)
          if (answer) {
            // Reverse score for reversed questions
            const score = question.reversed ? (6 - answer.value) : answer.value
            total += score
          }
        }

        // Calculate average and normalize to 0-100 scale
        const average = total / traitQuestions.length
        scores[trait] = Math.round(((average - 1) / 4) * 100)
      }

      const personalityResult: PersonalityResult = {
        extraversion: scores.extraversion ?? 0,
        agreeableness: scores.agreeableness ?? 0,
        conscientiousness: scores.conscientiousness ?? 0,
        emotionalStability: scores.emotionalStability ?? 0,
        openness: scores.openness ?? 0,
      }

      // Save to backend API
      const response = await $fetch<{ id: string }>('http://localhost:8080/api/results', {
        method: 'POST',
        body: {
          extraversion: personalityResult.extraversion,
          agreeableness: personalityResult.agreeableness,
          conscientiousness: personalityResult.conscientiousness,
          emotional_stability: personalityResult.emotionalStability,
          openness: personalityResult.openness,
        },
      })

      result.value = { ...personalityResult, id: response.id }
      return { ...personalityResult, id: response.id }
    } finally {
      isLoading.value = false
    }
  }

  const resetQuestionnaire = () => {
    currentQuestionIndex.value = 0
    answers.value = []
    result.value = null
  }

  const getQuestions = () => questions

  return {
    currentQuestionIndex,
    currentQuestion,
    totalQuestions,
    progress,
    isComplete,
    isLoading,
    answers,
    result,
    answerQuestion,
    previousQuestion,
    getCurrentAnswer,
    calculateResults,
    resetQuestionnaire,
    getQuestions,
  }
}
