<script setup lang="ts">
import type { Border } from "../types/BorderTypes";
const props = defineProps<{
  BorderInfo: Border;
}>();
</script>

<template>
  <div class="-mt-10 relative w-full overflow-hidden flex justify-center">
    <img
      :src="props.BorderInfo.AssetURL"
      :alt="props.BorderInfo.AssetName"
      class="w-full object-cover"
      style="height: 70px"
    />
  </div>
</template>
