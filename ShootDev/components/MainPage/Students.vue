<script setup lang="ts">
    import DemoCourseButton from '../ui/DemoCourseButton.vue';
    import BottomBorder from '../ui/BottomBorder.vue';
    import borderImg from '../../assets/pictures/mainpage/wooddivider.webp';
    import type { Border } from '../types/BorderTypes';

    const borderInfo: Border = {
	AssetName: "WoodenDivider",
	AssetURL: borderImg,
    }

    import type { FormType } from "../types/FormTypes";
    const props = defineProps<{
      formState: FormType;
    }>();

    function onSetLogIn(newForm: FormType) {
      emits("setFormState", newForm);
    }

    const emits = defineEmits<{
      setFormState: [setFormState: FormType];
    }>();

</script>
<template>
    <div class="pb-20 h-auto min-h-[110vh] bg-cover bg-no-repeat bg-[url(assets/pictures/mainpage/thirdbackground.png)]">

	<h1 class="mx-auto max-w-xs leading-none lg:max-w-2xl text-gray-100 lg:text-4xl text-3xl text-white text-center font-marko pt-20"> Join 12 people who would like to chat here instead of discord </h1>
	<h2 class="text-center font-marko mx-auto pt-10 max-w-xs text-xl text-gray-300 md:max-w-lg lg:text-xl"> Connect with an online community of people vibin' and learning about the freshes memes </h2>

	<div class="mx-auto mb-16 mt-8 flex w-full max-w-xl columns-1 flex-row justify-evenly text-center lg:columns-3"> 
	    <div><h3 class="text-4xl text-gray-400 font-bold">12</h3><p class="text-xs text-gray-400">ackshual hoomans</p></div>
	    <div><h3 class="text-4xl text-gray-400 font-bold">69.9</h3><p class="text-xs text-gray-400">hours chatted</p></div>
	    <div><h3 class="text-4xl text-gray-400 font-bold">9000</h3><p class="text-xs text-gray-400">animes mentioned</p></div>
	</div>
    <div class="text-gray-300 font-lora mb-8 grid w-full max-w-4xl grid-cols-1 gap-4 px-6 lg:grid-cols-3 mx-auto">
      <div class="flex flex-col gap-4 p-6 border border-yellow-500 shadow-md shadow-yellow-800 md:text-2xl lg:col-span-3">
	  <h4 class="font-dmsans text-[20px]">"Ryan told me he wanted to work on his Capstone Project and kept pestering me with types and monads, so I said, you seem to like to chat a lot, why don't you make a place where you can?"</h4>
	  <div class="flex flex-row gap-2 justify-center">
	      <img src="../../assets/pictures/mainpage/mppeople/dan.webp" width="48" height="48" class="m-1 box-content rounded-full border border-transparent hover:border-yellow-500" alt="Dan" </img>
	  <div class="flex flex-col">
		<p class="text-sm">Dan the Man</p>
		<p class="text-xs text-gray-400">Dancity&trade;</p>
		<p class="text-xs text-gray-400">Nortown, Normerica</p>
	  </div>
	  </div>
      </div>

      <div class="flex flex-col gap-4 p-6 border border-yellow-500 shadow-md shadow-yellow-800">

	  <h4 class="font-dmsans text-[20px]">"I logged on here by accident one time and haven't been able to make it out ever since"</h4>
	  <div class="flex flex-row gap-2 justify-center">
	      <img src="../../assets/pictures/mainpage/mppeople/dan.webp" width="48" height="48" class="m-1 box-content rounded-full border border-transparent hover:border-yellow-500" alt="Dan" 
	      </ img>
	  <div class="flex flex-col">
		<p class="text-sm">Still Dan</p>
		<p class="text-xs text-gray-400">StillDancity&trade;</p>
		<p class="text-xs text-gray-400">Nortown, Normerica</p>
	  </div>
      </div>
      </div>
      <div class="flex flex-col gap-4 p-6 border border-yellow-500 shadow-md shadow-yellow-800 lg:col-span-2">

	  <h4 class="font-dmsans text-[20px]" >"Ryan promised me a wealth of educated discussions about functional prorgams but after using the chat, I feel less certain about monads, and realized that a computer is just a state machine"</h4>
	  <div class="flex flex-row gap-2 justify-center">
	      <img src="../../assets/pictures/mainpage/mppeople/waseem.webp" width="48" height="48" class="m-1 box-content rounded-full border border-transparent hover:border-yellow-500" alt="waseem" </img>
	  <div class="flex flex-col">
		<p class="text-sm">Waseem</p>
		<p class="text-xs text-gray-400">Waseem Inc.</p>
		<p class="text-xs text-gray-400">Watown, Wajersey</p>
	  </div>
      </div>
      </div>
      <div class="bg-image-blue flex flex-col gap-4 p-6 border border-yellow-500 shadow-md shadow-yellow-800 lg:col-span-2">
	<h4 class="font-dmsans text-[20px]">"Ryan Convinced me that I people chat about Elixir in here all the time but its literally just 12 people talking about how great functional programming is...but I am not complaining"</h4>

	  <div class="flex flex-row gap-2 justify-center">
	      <img src="../../assets/pictures/mainpage/mppeople/waseem.webp" width="48" height="48" class="m-1 box-content rounded-full border border-transparent hover:border-yellow-500" alt="waseeem" </img>
	  <div class="flex flex-col">
		<p class="text-sm">Waseem</p>
		<p class="text-xs text-gray-400">Waseem Inc.</p>
		<p class="text-xs text-gray-400">Watown, Wajersey</p>
	  </div>
      </div>
      </div>
      <div class="bg-image-blue flex flex-col gap-4 p-6 border border-yellow-500 shadow-md shadow-yellow-800">
      <h4 class="font-dmsans text-[20px]">
	    "Someone told me there would be a Primeagen gift certificate"
      </h4>
	  <div class="flex flex-row gap-2 justify-center">
	      <img src="../../assets/pictures/mainpage/mppeople/ryan.webp" width="48" height="48" class="m-1 box-content rounded-full border border-transparent hover:border-yellow-500" alt="ryan" </img>
	  <div class="flex flex-col">
		<p class="text-sm">Ryan</p>
		<p class="text-xs text-gray-400">Ryan ApS</p>
		<p class="text-xs text-gray-400">Ryan, Iowa</p>
	  </div>
      </div>
    </div>
    </div>
    <div class="mx-auto text-center">
	<DemoCourseButton ButtonName="Start the Chat!" :logIn="props.formState" @setFormState="onSetLogIn"/>
	<p class="pt-5 text-gray-500" > (Click here to Chat!) </p>
	<p class="pt-20 text-gray-500"> Start Bikeshedding Now! </p>
    </div>
</div>
    <BottomBorder :BorderInfo="borderInfo" class="h-10"/>
</template>


<style scoped>

</style>
