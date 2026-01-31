<script setup lang="ts">
  const router = useRouter()
  const { t } = useI18n()
  const localePath = useLocalePath()

  const startTest = () => {
    router.push(localePath('/questionnaire'))
  }

  const speedRun = async () => {
    const { getQuestions, resetQuestionnaire, answers, calculateResults } = useQuestionnaire()
    
    // Reset any existing progress
    resetQuestionnaire()
    
    // Fill in random answers for all questions
    const allQuestions = getQuestions()
    for (const question of allQuestions) {
      const randomValue = Math.floor(Math.random() * 5) + 1 // 1-5
      answers.value.push({ questionId: question.id, value: randomValue })
    }
    
    // Calculate results and navigate to the result page
    const result = await calculateResults()
    router.push(localePath(`/results/${result.id}`))
  }

  const traits = [
    { key: 'extraversion', icon: 'i-lucide-zap', color: 'text-yellow-500' },
    { key: 'agreeableness', icon: 'i-lucide-heart-handshake', color: 'text-green-500' },
    { key: 'conscientiousness', icon: 'i-lucide-target', color: 'text-blue-500' },
    { key: 'emotionalStability', icon: 'i-lucide-anchor', color: 'text-purple-500' },
    { key: 'openness', icon: 'i-lucide-lightbulb', color: 'text-pink-500' }
  ]
</script>

<template>
  <div class="home-page min-h-screen">
    <!-- Hero Section -->
    <section class="hero-section relative overflow-hidden pt-20 pb-24 md:pt-32 md:pb-32">
      <!-- Background decoration -->
      <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <div class="absolute -top-40 -right-40 w-96 h-96 bg-gradient-to-br from-green-400/20 to-teal-500/10 rounded-full blur-3xl" />
        <div class="absolute -bottom-20 -left-20 w-72 h-72 bg-gradient-to-tr from-violet-400/15 to-blue-500/10 rounded-full blur-3xl" />
      </div>

      <div class="relative max-w-5xl mx-auto px-4 sm:px-6">
        <!-- Main heading -->
        <div class="text-center space-y-6">
          <h1 class="hero-title text-5xl md:text-7xl font-bold tracking-tight text-gray-900 dark:text-white">
            {{ t('home.title') }}
          </h1>
          <p class="hero-subtitle text-xl md:text-2xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto leading-relaxed">
            {{ t('home.description') }}
          </p>
        </div>

        <!-- CTA -->
        <div class="mt-12 flex flex-col items-center gap-4">
          <UButton
            size="xl"
            class="cta-button px-8 py-4 text-lg font-semibold shadow-lg shadow-green-500/25 hover:shadow-xl hover:shadow-green-500/30 transition-all duration-300"
            @click="startTest"
          >
            <UIcon name="i-lucide-play" class="w-5 h-5 mr-2" />
            {{ t('home.startButton') }}
          </UButton>
          <p class="text-sm text-gray-500 dark:text-gray-400 flex items-center gap-2">
            <UIcon name="i-lucide-clock" class="w-4 h-4" />
            {{ t('app.duration') }}
          </p>
          <button
            class="speed-run-btn mt-2 text-xs text-gray-300 dark:text-gray-600 hover:text-gray-400 dark:hover:text-gray-500 transition-colors cursor-pointer"
            @click="speedRun"
          >
            ⚡
          </button>
        </div>
      </div>
    </section>

    <!-- Big Five Traits Section -->
    <section class="traits-section py-16 md:py-24 bg-gray-50/50 dark:bg-gray-900/50">
      <div class="max-w-5xl mx-auto px-4 sm:px-6">
        <h2 class="text-2xl md:text-3xl font-bold text-center text-gray-900 dark:text-white mb-4">
          {{ t('home.traitsTitle') }}
        </h2>
        <p class="text-center text-gray-600 dark:text-gray-400 mb-12 max-w-2xl mx-auto">
          {{ t('home.traitsSubtitle') }}
        </p>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-4">
          <div
            v-for="(trait, index) in traits"
            :key="trait.key"
            class="trait-card group relative p-6 bg-white dark:bg-gray-800 rounded-2xl border border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600 transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
            :style="{ animationDelay: `${index * 100}ms` }"
          >
            <div class="flex flex-col items-center text-center">
              <div
                class="w-12 h-12 rounded-xl flex items-center justify-center mb-4 bg-gray-100 dark:bg-gray-700/50 group-hover:scale-110 transition-transform duration-300"
              >
                <UIcon :name="trait.icon" :class="['w-6 h-6', trait.color]" />
              </div>
              <h3 class="font-semibold text-gray-900 dark:text-white mb-2">
                {{ t(`traits.${trait.key}`) }}
              </h3>
              <p class="text-sm text-gray-500 dark:text-gray-400 leading-relaxed">
                {{ t(`traitDescriptions.${trait.key}`) }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- IPIP Foundation Section -->
    <section class="ipip-section py-16 md:py-24">
      <div class="max-w-4xl mx-auto px-4 sm:px-6">
        <div class="ipip-card relative overflow-hidden rounded-2xl sm:rounded-3xl bg-gradient-to-br from-gray-900 to-gray-800 dark:from-gray-800 dark:to-gray-900 p-6 sm:p-8 md:p-12 text-white">
          <!-- Decorative elements -->
          <div class="absolute top-0 right-0 w-64 h-64 bg-gradient-to-br from-green-500/20 to-transparent rounded-full blur-2xl" />
          <div class="absolute bottom-0 left-0 w-48 h-48 bg-gradient-to-tr from-blue-500/10 to-transparent rounded-full blur-2xl" />

          <div class="relative">
            <div class="flex items-center gap-3 mb-6">
              <div class="flex items-center justify-center w-10 h-10 rounded-lg bg-green-500/20">
                <UIcon name="i-lucide-flask-conical" class="w-5 h-5 text-green-400" />
              </div>
              <span class="text-sm font-medium text-green-400 uppercase tracking-wider">
                {{ t('home.scientificFoundation') }}
              </span>
            </div>

            <h2 class="text-2xl sm:text-3xl md:text-4xl font-bold mb-4 sm:mb-6">
              {{ t('home.ipipTitle') }}
            </h2>

            <div class="space-y-3 sm:space-y-4 text-sm sm:text-base text-gray-300 leading-relaxed">
              <p>{{ t('home.ipipDescription1') }}</p>
              <p>{{ t('home.ipipDescription2') }}</p>
            </div>

            <div class="mt-6 sm:mt-8 flex flex-wrap gap-3 sm:gap-4">
              <div class="flex items-center gap-2 text-sm text-gray-400">
                <UIcon name="i-lucide-check-circle-2" class="w-4 h-4 text-green-400" />
                {{ t('home.ipipFeature1') }}
              </div>
              <div class="flex items-center gap-2 text-sm text-gray-400">
                <UIcon name="i-lucide-check-circle-2" class="w-4 h-4 text-green-400" />
                {{ t('home.ipipFeature2') }}
              </div>
              <div class="flex items-center gap-2 text-sm text-gray-400">
                <UIcon name="i-lucide-check-circle-2" class="w-4 h-4 text-green-400" />
                {{ t('home.ipipFeature3') }}
              </div>
            </div>

            <div class="mt-8 pt-8 border-t border-gray-700">
              <a
                href="https://ipip.ori.org/"
                target="_blank"
                rel="noopener noreferrer"
                class="inline-flex items-center gap-2 text-sm text-green-400 hover:text-green-300 transition-colors"
              >
                {{ t('home.learnMoreIpip') }}
                <UIcon name="i-lucide-external-link" class="w-4 h-4" />
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="py-8 border-t border-gray-200 dark:border-gray-800">
      <div class="max-w-5xl mx-auto px-4 sm:px-6 text-center">
        <p class="text-sm text-gray-500 dark:text-gray-400">
          {{ t('home.footer') }}
        </p>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.hero-title {
  background: linear-gradient(135deg, var(--color-gray-900) 0%, var(--color-gray-700) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

:root.dark .hero-title {
  background: linear-gradient(135deg, #fff 0%, #a3a3a3 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.trait-card {
  animation: fadeInUp 0.5s ease-out both;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.ipip-badge {
  animation: fadeIn 0.6s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.cta-button {
  transition: transform 0.2s ease-out;
}

.cta-button:hover {
  transform: translateY(-2px);
}

.cta-button:active {
  transform: translateY(0);
}

.speed-run-btn {
  opacity: 0.3;
  font-size: 10px;
  padding: 4px 8px;
  border-radius: 4px;
  background: transparent;
  border: none;
}

.speed-run-btn:hover {
  opacity: 0.6;
  background: rgba(128, 128, 128, 0.1);
}
</style>
