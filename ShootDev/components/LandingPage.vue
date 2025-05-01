<script setup lang="ts">
import TopBanner from "./MainPage/TopBanner.vue";
import WhyStart from "./MainPage/WhyStart.vue";
import Students from "./MainPage/Students.vue";
import EndBanner from "./MainPage/EndBanner.vue";

import type { FormType } from "./types/FormTypes";

const props = defineProps<{
  formState: FormType;
}>();

const emits = defineEmits<{
  setFormState: [formState: FormType];
}>();

function onSetLogIn(formState: FormType) {
  emits("setFormState", formState);
}
</script>
<template>
  <TopBanner :formState="props.formState" @setFormState="onSetLogIn" />
  <WhyStart />
  <Students :formState="props.formState" @setFormState="onSetLogIn" />
  <EndBanner />
</template>
