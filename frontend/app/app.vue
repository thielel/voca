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
  <UApp>
    <div class="fixed top-3 right-3 sm:top-4 sm:right-4 z-50">
      <LanguageSwitcher />
    </div>
    <UContainer>
      <NuxtPage />
    </UContainer>
  </UApp>
</template>
