<script setup lang="ts">
import type { FormType } from "../types/FormTypes";

const props = defineProps<{
  ButtonName: string;
}>();

const emits = defineEmits<{
  setFormState: [formState: FormType];
}>();
</script>

<template>
  <div class="font-marko">
    <button @click="emits('setFormState', 'init')" class="text-shadow-lg/30">
      {{ props.ButtonName }}
    </button>
  </div>
</template>

<style scoped>
button {
  background: linear-gradient(to bottom, #8f703c, #7f673e);
  color: #c0c4c7;
  border: 1px solid #342715;
  padding: 10px 20px;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
  box-shadow:
    inset 0 1px 0 0 rgba(255, 255, 255, 0.3),
    0 2px 0 0 #342715,
    0 4px 0 0 #342715,
    0 6px 0 0 #130f08,
    0 8px 0 0 #130f08,
    0 8px 16px 0 rgba(52, 39, 21, 0.5);
  overflow: hidden;
  transition:
    transform 0.4s ease,
    box-shadow 0.4s ease;
}

button::before {
  content: "";
  position: absolute;
  inset: 0;
  background: linear-gradient(
    to right,
    transparent,
    rgba(255, 255, 255, 0.1),
    transparent
  );
  transform: translateX(-100%);
  animation: shimmer 2s infinite;
  box-shadow: 0 0 5px #130f08;
}

button:hover {
  transform: translateY(5px);
  background: linear-gradient(to bottom, #b8860b, #daa520);
  box-shadow:
    inset 0 1px 0 0 rgba(255, 255, 255, 0.3),
    0 1px 0 1px #342715,
    0 1px 0 1px #342715,
    0 1px 0 1px #130f08,
    0 1px 0 1px #130f08,
    0 1px 1px 0 rgba(52, 39, 21, 0.5);
}

button::hover i {
  animation: bounce 2s infinite;
}

@keyframes bounce {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-25%);
  }
}
</style>
