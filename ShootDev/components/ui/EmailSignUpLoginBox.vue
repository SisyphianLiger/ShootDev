<script setup lang="ts">
import * as v from "valibot";
import type { FormSubmitEvent } from "@nuxt/ui";
import type { FormType } from "../types/FormTypes";
import DemoCourseButton from "./DemoCourseButton.vue";

const emailForm = v.object({
  email: v.pipe(v.string(), v.email("Invalid Email")),
  password: v.pipe(
    v.string(),
    v.minLength(8, "Must Be Longer than 8 Characters"),
  ),
});

type EmailForm = v.InferOutput<typeof emailForm>;

const emailState = reactive({
  email: "",
  password: "",
});

const toast = useToast();
function onSetLogIn(newForm: FormType) {
  emits("setFormState", newForm);
}

async function onSubmit(event: FormSubmitEvent<EmailForm>) {
  toast.add({
    title: "Success",
    description: "The form has been submitted.",
    color: "success",
  });
  console.log(event.data);
}

const emits = defineEmits<{
  setFormState: [formState: FormType];
}>();
</script>

<template>
  <div class="fixed inset-0 bg-black/70 z-30"></div>
  <div class="fixed inset-0 flex items-center justify-center z-40">
    <div
      class="min-h-[400px] border border-yellow-500 shadow-md shadow-yellow-800 relative bg-[url(assets/pictures/mainpage/thirdbackground.png)] rounded-lg p-8 shadow-lg w-full max-w-md"
    >
      <button
        class="absolute top-2 right-4 text-[32px] text-gray-500 item-right hover:text-gray-800"
        @click="onSetLogIn('quit')"
      >
        X
      </button>
      <h1 class="text-white flex justify-center text-[40px] pt-5 font-marko">
        Sign Up
      </h1>
      <div class="text-white">
        <UForm
          :schema="emailForm"
          :state="emailState"
          class="pt-2 "
          @submit="onSubmit"
        >
          <UFormField 
	   :ui="{ error: 'text-red-500 font-semibold text-sm pt-1' }"
	  class="w-full h-30 font-marko text-[20px]" label="Email" name="email" >
            <UInput
	    size="lg"
              class="w-full min-h-[45px] border border-yellow-500 shadow-md shadow-yellow-800 text-white placeholder-gray-400"
              v-model="emailState.email"
              type="email"
            />
          </UFormField>

          <UFormField 
	   :ui="{ error: 'text-red-500 font-semibold text-sm pt-1' }"
	  class="w-full h-30 font-marko text-[20px]" label="Password" name="password" >
            <UInput
	    size="lg"
              class="w-full min-h-[45px] border border-yellow-500 shadow-md shadow-yellow-800 text-white placeholder-gray-400"
              v-model="emailState.password"
              type="password"
            />
          </UFormField>
          <div class="pt-3 flex justify-center ">
	    <button class="border text-[#c0c4c7] rounded-full w-40 h-12 border-yellow-500 shadow-md shadow-yellow-800 bg-yellow-900 hover:bg-yellow-800 transition-colors duration-200">
	      Sign Up
	    </button>
          </div>
        </UForm>
      </div>
    </div>
  </div>
</template>

