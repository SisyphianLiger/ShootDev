<script setup lang="ts">
import DemoCourseButton from "~/components/ui/DemoCourseButton.vue";
import IntroCompanies from "../ui/IntroCompanies.vue";
import PeopleBanner from "../ui/PeopleBanner.vue";
import BottomBorder from "../ui/BottomBorder.vue";
import bottomBorderImg from "../../assets/pictures/mainpage/border.webp";
import type { Border } from "../types/BorderTypes";

const borderInfo: Border = {
  AssetName: "DiamondBorder",
  AssetURL: bottomBorderImg,
};

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
  <div class="bg-image-bootdev">
    <div class="flex gap-6 items-center px-5 pt-6">
      <div>
        <img
          src="../../assets/pictures/Shoot.dev_logo.png"
          alt="pic"
          class="w-64 pl-20"
        />
      </div>
      <div>
        <span
          @click="onSetLogIn('init')"
          class="text-white font-marko hover:underline hover:text-gray-400"
          >Sign In & Skip Demo →</span
        >
      </div>
    </div>
    <div class="flex flex-col justify-center space-y-6 py-20 items-center">
      <div class="text-white text-center">
        <h1 class="text-[48px] font-marko">
          Chat About Backend Development the Fun Way
        </h1>
        <h2 class="pt-10 font-muli text-[25px] text-gray-400">
          Welcome to the most captivating front-end for a chat-app portfolio
          project ever made...probably<br />
          A fully fledged chat app for the not so subtle bike-shedding
          programmer
        </h2>
      </div>
      <div class="pt-15">
        <DemoCourseButton
          ButtonName="Start the Chat!"
          @setFormState="onSetLogIn"
        />
      </div>
      <span class="text-gray-400 text-[20px] font-muli">
        It's pretty quick to log in though
      </span>
    </div>
    <div
      class="flex flex-col justify-center items-center relative hidden md:hidden lg:flex"
    >
      <PeopleBanner />
      <IntroCompanies />
    </div>
  </div>
  <BottomBorder :BorderInfo="borderInfo" />
</template>

<style scoped>
.bg-image-bootdev {
  background-blend-mode: hue, hard-light, normal, normal;
  background-image:
    linear-gradient(0deg, rgba(85, 98, 246, 0.2), rgba(85, 98, 246, 0.2)),
    url("../../assets/pictures/mainpage/boot_dev_main_bg.png");
  background-position: 50%;
  background-repeat: no-repeat;
  background-size: cover;
  min-height: 100vh;
}
</style>
