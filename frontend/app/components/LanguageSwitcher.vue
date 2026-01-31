<script setup lang="ts">
  type LocaleCode = 'de' | 'en' | 'tr' | 'ar' | 'ru' | 'pl' | 'ro' | 'it' | 'uk' | 'bg'

  const { locale, locales, setLocale } = useI18n()

  const availableLocales = computed(() => {
    return locales.value.filter((l): l is { code: string, name: string, file: string, dir: string } => typeof l !== 'string')
  })

  const currentLocale = computed(() => {
    const current = availableLocales.value.find((l) => l.code === locale.value)
    return current || { code: 'en', name: 'English' }
  })

  const items = computed(() => {
    return availableLocales.value.map((loc) => ({
      label: loc.name,
      icon: loc.code === locale.value ? 'i-lucide-check' : undefined,
      onSelect: () => setLocale(loc.code as LocaleCode)
    }))
  })
</script>

<template>
  <UDropdownMenu :items="items">
    <UButton variant="ghost" size="sm" class="gap-2">
      <UIcon name="i-lucide-languages" class="w-4 h-4" />
      <span class="hidden sm:inline">{{ currentLocale.name }}</span>
    </UButton>
  </UDropdownMenu>
</template>
