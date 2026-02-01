<script setup lang="ts">
  const { t, locale, locales } = useI18n()

  const currentLocale = computed(() => {
    const current = locales.value.find((l) => typeof l !== 'string' && l.code === locale.value)
    return typeof current !== 'string' ? current : undefined
  })

  useHead({
    htmlAttrs: {
      lang: locale,
      dir: currentLocale.value?.dir as 'ltr' | 'rtl' | undefined
    }
  })

  useSeoMeta({
    title: () => `Voca - ${t('app.title')}`,
    description: () => t('home.description')
  })
</script>

<template>
  <UApp class="min-h-screen bg-[#F8FAFC]">
    <!-- Fixed colorful background - covers entire viewport -->
    <div class="fixed inset-0 pointer-events-none overflow-hidden -z-10">
      <div class="absolute -top-40 -right-40 w-[500px] h-[500px] bg-gradient-to-br from-emerald-400/30 to-teal-300/20 rounded-full blur-3xl" />
      <div class="absolute -bottom-40 -left-40 w-[500px] h-[500px] bg-gradient-to-br from-violet-400/30 to-purple-300/20 rounded-full blur-3xl" />
      <div class="absolute top-1/2 right-0 w-[300px] h-[300px] bg-gradient-to-br from-rose-400/20 to-pink-300/15 rounded-full blur-3xl" />
      <div class="absolute top-1/4 left-1/4 w-[250px] h-[250px] bg-gradient-to-br from-amber-400/20 to-orange-300/15 rounded-full blur-3xl" />
    </div>
    
    <div class="fixed top-3 right-3 sm:top-4 sm:right-4 z-50">
      <LanguageSwitcher />
    </div>
    <UContainer>
      <NuxtPage />
    </UContainer>
  </UApp>
</template>
