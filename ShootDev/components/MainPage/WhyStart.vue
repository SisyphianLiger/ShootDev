<script setup lang="ts">
import Grid from "../ui/Grid.vue";
import BottomBorder from "../ui/BottomBorder.vue";
import borderImg from "../../assets/pictures/mainpage/wooddivider.webp";
import type { Border } from "../types/BorderTypes";
const borderInfo: Border = {
  AssetName: "WoodenDivider",
  AssetURL: borderImg,
};
</script>
<template>
  <div
    class="-mt-10 pb-20 h-auto min-h-[60vh] bg-cover bg-no-repeat bg-[url(assets/pictures/mainpage/2_main_pg_bg.png)]"
  >
    <div class="flex flex-col items-center pt-24">
      <h1
        class="-mt-6 text-white font-marko text-[50px] text-center leading-tight"
      >
        Why Stop Chatting?
      </h1>
      <h2
        class="text-stone-500 font-dmsans text-[30px] text-center mx-auto pt-15 pb-20 px-10 w-full"
      >
        The only way to bikeshed about your lang is to chat with others
      </h2>
    </div>
    <div class="flex justify-center">
      <Grid />
    </div>
  </div>
  <BottomBorder :BorderInfo="borderInfo" class="h-10" />
</template>

<style scoped></style>
