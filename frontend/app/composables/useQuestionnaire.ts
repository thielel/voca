import { questions, type Question, type Answer, type PersonalityResult } from '~/data/questions'

export type { Question, Answer, PersonalityResult }

export const useQuestionnaire = () => {
  const config = useRuntimeConfig()
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

  const calculateResults = async (language?: string): Promise<PersonalityResult & { id: string }> => {
    isLoading.value = true

    try {
      // Generate a session ID for this submission
      const sessionId = crypto.randomUUID()

      // Convert answers to backend format (question_id instead of questionId)
      const backendAnswers = answers.value.map(a => ({
        question_id: a.questionId,
        value: a.value,
      }))

      // Submit answers to backend - backend calculates scores and starts AI generation in background
      const response = await $fetch<{
        result: {
          id: string
          extraversion: number
          agreeableness: number
          conscientiousness: number
          emotional_stability: number
          openness: number
        }
      }>(`${config.public.apiUrl}/api/results`, {
        method: 'POST',
        body: {
          session_id: sessionId,
          answers: backendAnswers,
          language: language || 'de', // Pass language for background AI generation
        },
      })

      // Convert backend response to frontend format
      const personalityResult: PersonalityResult & { id: string } = {
        id: response.result.id,
        extraversion: response.result.extraversion,
        agreeableness: response.result.agreeableness,
        conscientiousness: response.result.conscientiousness,
        emotionalStability: response.result.emotional_stability,
        openness: response.result.openness,
      }

      result.value = personalityResult
      return personalityResult
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
