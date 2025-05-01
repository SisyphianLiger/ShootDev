<script lang="ts" setup>
import type { CardTypes } from "../types/CardTypes";
import { LandingCardSVG } from "../types/CardTypes";

const props = defineProps<{
  card: CardTypes;
}>();

const iconSvg = LandingCardSVG[props.card.iconKey] || "Error Icon Not Found";
</script>

<template>
  <div>
    <div class="text-white" v-html="iconSvg"></div>
    <p class="font-marko text-[25px] text-bold text-white">
      {{ props.card.headline }}
    </p>
    <p class="font-dmsans text-[20px] text-stone-500">
      {{ props.card.headlinetwo }}
    </p>
  </div>
</template>
