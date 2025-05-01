<script setup lang="ts">
import NewCard from "./NewCard.vue";
import type { CardTypes } from "../types/CardTypes";
const cts: CardTypes[] = [
  {
    headline: "chat about tutorial hell",
    headlinetwo: "by bikeshedding a ton",
    iconKey: "fire",
  },
  {
    headline: "quickly waste time with",
    headlinetwo: "memes, jokes, and lols",
    iconKey: "lightning",
  },
  {
    headline: "don't build just chat",
    headlinetwo: "No time to build",
    iconKey: "tools",
  },
  {
    headline: "talk about stuff",
    headlinetwo: "doomscroll and have fun",
    iconKey: "graduatehat",
  },
  {
    headline: "chat constantly online",
    headlinetwo: "interrupting your builds",
    iconKey: "cloud",
  },
  {
    headline: "Free because its a project",
    headlinetwo: "Is that true chat?",
    iconKey: "money",
  },
];
</script>

<template>
  <div class="grid gap-20 pl-[2%] sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-3">
    <NewCard class="textlg" v-for="card in cts" :card="card" />
  </div>
</template>
