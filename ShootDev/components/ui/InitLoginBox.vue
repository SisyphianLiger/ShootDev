<script setup lang="ts">
import type { FormType } from "../types/FormTypes";

const props = defineProps<{
  formState: FormType;
}>();

function onSetLogIn(newForm: FormType) {
  emits("setFormState", newForm);
}

const emits = defineEmits<{
  setFormState: [formState: FormType];
}>();
</script>

<template>
  <div class="fixed inset-0 bg-black/70 z-30"></div>
  <div class="fixed inset-0 flex items-center justify-center z-40">
    <div
      class="h-[400px] border border-yellow-500 shadow-md shadow-yellow-800 relative bg-[url(assets/pictures/mainpage/thirdbackground.png)] rounded-lg p-8 shadow-lg w-full max-w-md"
    >
      <button
        class="absolute top-2 right-4 text-[32px] text-gray-500 item-right hover:text-gray-800"
        @click="onSetLogIn('quit')"
      >
        X
      </button>
      <h2 class="text-3xl pb-10 pt-5 font-marko mb-4 text-white text-center">
        Sign in to Shoot.dev
      </h2>
      <div class="flex flex-col gap-8">
        <div
          class="mx-auto w-95 h-16 flex gap-5 justify-center py-2 px-4 border text-white font-dmsans rounded-full hover:bg-gray-400"
        >
          <div
            @click="onSetLogIn('signUp')"
            class="w-84 flex items-center gap-5 justify-start"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="lucide lucide-mail-icon lucide-mail"
            >
              <rect width="20" height="16" x="2" y="4" rx="2" />
              <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
            </svg>
            <span class="font-dmsans text-[24px] pl-10"
              >Sign up with Email</span
            >
          </div>
        </div>

        <div
          class="mx-auto w-95 flex gap-5 h-16 justify-center py-2 px-4 border text-white font-dmsans rounded-full hover:bg-gray-400"
        >
          <div
            @click="onSetLogIn('logIn')"
            class="w-84 flex items-center gap-5 justify-start"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 2e"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="lucide lucide-log-in-icon lucide-log-in"
            >
              <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4" />
              <polyline points="10 17 15 12 10 7" />
              <line x1="15" x2="3" y1="12" y2="12" />
            </svg>
            <span class="font-dmsans text-[24px] pl-20">Log In</span>
          </div>
        </div>
      </div>

      <div class="mt-4 flex justify-center text-[20px] pt-10 text-gray-600">
        <label class="inline-flex items-center">
          <input type="checkbox" class="form-checkbox mr-2" checked />
          I’ve read and agree to the &nbsp;<a
            href="/terms"
            class="text-blue-600 underline"
          >
            terms</a
          >.
        </label>
      </div>
    </div>
  </div>
</template>
