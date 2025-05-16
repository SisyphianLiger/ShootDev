<script setup lang="ts">
import LandingPage from "~/components/LandingPage.vue";

import type { FormType } from "~/components/types/FormTypes";

import InitLoginBox from "~/components/ui/InitLoginBox.vue";
import LoginBox from "~/components/ui/LoginBox.vue";
import EmailSignUpLoginBox from "~/components/ui/EmailSignUpLoginBox.vue";

const formState = ref<"init" | "logIn" | "signUp" | "forgotPass" | "quit">(
  "quit",
);

function setFormState(newState: FormType) {
  formState.value = newState;
}
</script>
<template>
  <div>
    <LandingPage :formState="formState" @setFormState="setFormState" />
    <div>
      <template v-if="formState === 'init'">
        <InitLoginBox :formState="formState" @setFormState="setFormState" />
      </template>
      <template v-if="formState === 'logIn'">
        <LoginBox :formState="formState" @setFormState="setFormState" />
      </template>
      <template v-if="formState === 'signUp'">
        <EmailSignUpLoginBox
          :formState="formState"
          @setFormState="setFormState"
        />
      </template>
    </div>
  </div>
</template>

<style scoped></style>
