// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,

  modules: ["@nuxt/eslint", "@nuxt/ui"],

  devtools: {
    enabled: false,
  },

  css: ["~/assets/css/main.css"],

  devServer: {
    host: "0.0.0.0",
    port: 9245,
  },

  nitro: {
    output: {
      publicDir: "dist",
    },
  },

  future: {
    compatibilityVersion: 4,
  },

  compatibilityDate: "2025-01-15",

  eslint: {
    config: {
      stylistic: {
        commaDangle: "never",

        braceStyle: "1tbs",
      },
    },
  },
});
