// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  typescript: {
    typeCheck: true,
  },

  css: ["~/assets/main.css"],

  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },

  vite: {
    plugins: [tailwindcss()],
  },

  modules: ["@nuxt/ui", "@nuxt/fonts"],
  fonts: {
    priority: ["bunny", "google"],
    families: [
      { name: "Marko One", provider: "bunny" },
      { name: "Cutive Mono", provider: "bunny" },
      { name: "Lora", provider: "google" },
    ],
  },
});
