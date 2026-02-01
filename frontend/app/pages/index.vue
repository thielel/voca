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

  // Vibrant trait colors - more bold and engaging
  const traits = [
    { key: 'extraversion', icon: 'i-lucide-zap', color: 'text-rose-500', bgColor: 'bg-gradient-to-br from-rose-100 to-pink-200', borderColor: 'hover:border-rose-300' },
    { key: 'agreeableness', icon: 'i-lucide-heart-handshake', color: 'text-emerald-500', bgColor: 'bg-gradient-to-br from-emerald-100 to-teal-200', borderColor: 'hover:border-emerald-300' },
    { key: 'conscientiousness', icon: 'i-lucide-target', color: 'text-amber-500', bgColor: 'bg-gradient-to-br from-amber-100 to-orange-200', borderColor: 'hover:border-amber-300' },
    { key: 'emotionalStability', icon: 'i-lucide-anchor', color: 'text-sky-500', bgColor: 'bg-gradient-to-br from-sky-100 to-blue-200', borderColor: 'hover:border-sky-300' },
    { key: 'openness', icon: 'i-lucide-lightbulb', color: 'text-violet-500', bgColor: 'bg-gradient-to-br from-violet-100 to-purple-200', borderColor: 'hover:border-violet-300' }
  ]
</script>

<template>
  <div class="home-page min-h-screen relative">
    <!-- Hero Section -->
    <section class="hero-section relative pt-20 pb-24 md:pt-32 md:pb-32">
      <div class="relative max-w-5xl mx-auto px-4 sm:px-6">
        <!-- Main heading -->
        <div class="text-center space-y-6">
          <h1 class="hero-title text-5xl md:text-7xl font-bold tracking-tight">
            {{ t('home.title') }}
          </h1>
          <p class="hero-subtitle text-xl md:text-2xl text-slate-600 max-w-2xl mx-auto leading-relaxed">
            {{ t('home.description') }}
          </p>
        </div>

        <!-- CTA -->
        <div class="mt-12 flex flex-col items-center gap-4">
          <button
            class="cta-button group relative px-10 py-5 text-lg font-bold text-white rounded-2xl bg-gradient-to-r from-sky-500 via-teal-500 to-emerald-500 shadow-xl shadow-teal-500/30 hover:shadow-2xl hover:shadow-teal-500/40 transition-all duration-300 hover:-translate-y-1"
            @click="startTest"
          >
            <span class="flex items-center gap-3">
              <UIcon name="i-lucide-sparkles" class="w-6 h-6 group-hover:rotate-12 transition-transform" />
              {{ t('home.startButton') }}
              <UIcon name="i-lucide-arrow-right" class="w-5 h-5 group-hover:translate-x-1 transition-transform" />
            </span>
          </button>
          <p class="text-sm text-slate-500 flex items-center gap-2 bg-white/60 backdrop-blur-sm px-4 py-2 rounded-full">
            <UIcon name="i-lucide-clock" class="w-4 h-4 text-teal-500" />
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
    <section class="traits-section py-16 md:py-24 relative">
      <div class="relative max-w-5xl mx-auto px-4 sm:px-6">
        <div class="text-center mb-12">
          <span class="inline-block px-4 py-2 bg-gradient-to-r from-violet-100 to-pink-100 text-violet-600 text-sm font-semibold rounded-full mb-4">
            {{ t('home.traitsTitle') }}
          </span>
          <h2 class="text-2xl md:text-4xl font-bold text-slate-800 mb-4">
            The Big Five Personality Traits
          </h2>
          <p class="text-slate-600 max-w-2xl mx-auto">
            {{ t('home.traitsSubtitle') }}
          </p>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-5">
          <div
            v-for="(trait, index) in traits"
            :key="trait.key"
            class="trait-card group relative p-6 bg-white rounded-2xl border-2 border-transparent shadow-lg shadow-slate-200/50 transition-all duration-300 hover:-translate-y-2 hover:shadow-xl"
            :class="trait.borderColor"
            :style="{ animationDelay: `${index * 100}ms` }"
          >
            <div class="flex flex-col items-center text-center">
              <div
                class="w-14 h-14 rounded-2xl flex items-center justify-center mb-4 group-hover:scale-110 group-hover:rotate-3 transition-all duration-300 shadow-md"
                :class="trait.bgColor"
              >
                <UIcon :name="trait.icon" :class="['w-7 h-7', trait.color]" />
              </div>
              <h3 class="font-bold text-slate-800 mb-2">
                {{ t(`traits.${trait.key}`) }}
              </h3>
              <p class="text-sm text-slate-500 leading-relaxed">
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
        <div class="ipip-card relative overflow-hidden rounded-3xl bg-gradient-to-br from-sky-400 via-teal-400 to-emerald-400 p-6 sm:p-8 md:p-12 text-white shadow-2xl shadow-teal-500/30">
          <!-- Decorative elements - light accent blobs -->
          <div class="absolute top-0 right-0 w-72 h-72 bg-gradient-to-br from-white/30 to-cyan-200/20 rounded-full blur-3xl" />
          <div class="absolute bottom-0 left-0 w-56 h-56 bg-gradient-to-tr from-emerald-300/30 to-teal-200/20 rounded-full blur-3xl" />
          <div class="absolute top-1/2 right-1/4 w-40 h-40 bg-gradient-to-br from-sky-300/25 to-blue-200/15 rounded-full blur-2xl" />

          <div class="relative">
            <div class="flex items-center gap-3 mb-6">
              <div class="flex items-center justify-center w-12 h-12 rounded-xl bg-white/25 backdrop-blur-sm shadow-lg">
                <UIcon name="i-lucide-flask-conical" class="w-6 h-6 text-white" />
              </div>
              <span class="text-sm font-bold text-white/90 uppercase tracking-wider">
                {{ t('home.scientificFoundation') }}
              </span>
            </div>

            <h2 class="text-2xl sm:text-3xl md:text-4xl font-bold mb-4 sm:mb-6">
              {{ t('home.ipipTitle') }}
            </h2>

            <div class="space-y-3 sm:space-y-4 text-sm sm:text-base text-white/90 leading-relaxed">
              <p>{{ t('home.ipipDescription1') }}</p>
              <p>{{ t('home.ipipDescription2') }}</p>
            </div>

            <div class="mt-6 sm:mt-8 flex flex-wrap gap-3 sm:gap-4">
              <div class="flex items-center gap-2 text-sm bg-white/25 backdrop-blur-sm px-3 py-1.5 rounded-full">
                <UIcon name="i-lucide-check-circle-2" class="w-4 h-4 text-white" />
                <span class="text-white font-medium">{{ t('home.ipipFeature1') }}</span>
              </div>
              <div class="flex items-center gap-2 text-sm bg-white/25 backdrop-blur-sm px-3 py-1.5 rounded-full">
                <UIcon name="i-lucide-check-circle-2" class="w-4 h-4 text-white" />
                <span class="text-white font-medium">{{ t('home.ipipFeature2') }}</span>
              </div>
              <div class="flex items-center gap-2 text-sm bg-white/25 backdrop-blur-sm px-3 py-1.5 rounded-full">
                <UIcon name="i-lucide-check-circle-2" class="w-4 h-4 text-white" />
                <span class="text-white font-medium">{{ t('home.ipipFeature3') }}</span>
              </div>
            </div>

            <div class="mt-8 pt-8 border-t border-white/25">
              <a
                href="https://ipip.ori.org/"
                target="_blank"
                rel="noopener noreferrer"
                class="inline-flex items-center gap-2 text-sm font-semibold text-teal-700 bg-white hover:bg-white/90 px-5 py-2.5 rounded-xl shadow-lg transition-all hover:shadow-xl"
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
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 25%, #d946ef 50%, #f43f5e 75%, #f97316 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  background-size: 200% 200%;
  animation: gradientShift 8s ease infinite;
}

@keyframes gradientShift {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
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
