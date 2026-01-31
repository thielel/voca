// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ['@nuxt/eslint', '@nuxt/ui', '@nuxtjs/i18n'],

  devtools: {
    enabled: true
  },

  css: ['~/assets/css/main.css'],

  routeRules: {
    '/': { prerender: true }
  },

  compatibilityDate: '2025-01-15',

  i18n: {
    locales: [
      { code: 'de', name: 'Deutsch', file: 'de.json', dir: 'ltr' },
      { code: 'en', name: 'English', file: 'en.json', dir: 'ltr' },
      { code: 'tr', name: 'Türkçe', file: 'tr.json', dir: 'ltr' },
      { code: 'ar', name: 'العربية', file: 'ar.json', dir: 'rtl' },
      { code: 'ru', name: 'Русский', file: 'ru.json', dir: 'ltr' },
      { code: 'pl', name: 'Polski', file: 'pl.json', dir: 'ltr' },
      { code: 'ro', name: 'Română', file: 'ro.json', dir: 'ltr' },
      { code: 'it', name: 'Italiano', file: 'it.json', dir: 'ltr' },
      { code: 'uk', name: 'Українська', file: 'uk.json', dir: 'ltr' },
      { code: 'bg', name: 'Български', file: 'bg.json', dir: 'ltr' }
    ],
    defaultLocale: 'en',
    bundle: {
      fullInstall: false
    },
    strategy: 'prefix_except_default',
    detectBrowserLanguage: {
      useCookie: true,
      cookieKey: 'i18n_redirected',
      redirectOn: 'root'
    }
  },

  eslint: {
    config: {
      stylistic: {
        commaDangle: 'never',
        braceStyle: '1tbs'
      }
    }
  }
})