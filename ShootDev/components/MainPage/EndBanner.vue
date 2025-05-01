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
    class="pb-20 h-auto min-h-[60vh] bg-cover bg-no-repeat bg-[url(assets/pictures/mainpage/2_main_pg_bg.png)]"
  >
    <div class="flex flex-col items-center pt-24 p-8">
      <!-- Disclaimer Box -->
      <div
        class="bg-black/70 border border-yellow-500 shadow-md shadow-yellow-800 rounded-lg mx-auto max-w-4xl p-10 mb-10"
      >
        <h4 class="text-white font-marko text-left mx-auto leading-tight">
          DISCLAIMER: This is a capstone project that is meant to look like
          <a href="https://boot.dev" class="text-yellow-400 hover:underline"
            >boot.dev</a
          >
          an amazing website where I learned how to become a better backend dev.
          I would highly encourage you check out their website if you are
          interested in learning programming. All information on this site is
          therefore a parody and should not be taken seriously. It was made to
          give the appearance of a front-end for a chat-app feature.
        </h4>
      </div>
      <!-- Navigation Links -->
      <nav class="flex flex-wrap justify-center gap-6 pt-5">
        <router-link
          to="/courses"
          class="text-stone-400 hover:text-yellow-300 font-marko text-xl transition-colors"
        >
          Courses
        </router-link>
        <router-link
          to="/leaderboard"
          class="text-stone-400 hover:text-yellow-300 font-marko text-xl transition-colors"
        >
          Leaderboard
        </router-link>
        <router-link
          to="/chat"
          class="text-stone-400 hover:text-yellow-300 font-marko text-xl transition-colors"
        >
          Chat
        </router-link>
        <router-link
          to="/contact"
          class="text-stone-400 hover:text-yellow-300 font-marko text-xl transition-colors"
        >
          Contact
        </router-link>
      </nav>
      <!-- Social Media Links -->
      <div class="flex justify-center pt-10 gap-5">
        <a href="https://x.com/SGrenedier">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            x="0px"
            y="0px"
            width="40"
            height="40"
            viewBox="0,0,256,256"
          >
            <g
              fill="#ffffff"
              fill-rule="nonzero"
              stroke="none"
              stroke-width="1"
              stroke-linecap="butt"
              stroke-linejoin="miter"
              stroke-miterlimit="10"
              stroke-dasharray=""
              stroke-dashoffset="0"
              font-family="none"
              font-weight="none"
              font-size="none"
              text-anchor="none"
              style="mix-blend-mode: normal"
            >
              <g transform="scale(5.12,5.12)">
                <path
                  d="M11,4c-3.85433,0 -7,3.14567 -7,7v28c0,3.85433 3.14567,7 7,7h28c3.85433,0 7,-3.14567 7,-7v-28c0,-3.85433 -3.14567,-7 -7,-7zM11,6h28c2.77367,0 5,2.22633 5,5v28c0,2.77367 -2.22633,5 -5,5h-28c-2.77367,0 -5,-2.22633 -5,-5v-28c0,-2.77367 2.22633,-5 5,-5zM13.08594,13l9.22266,13.10352l-9.30859,10.89648h2.5l7.9375,-9.29297l6.53906,9.29297h7.9375l-10.125,-14.38672l8.21094,-9.61328h-2.5l-6.83984,8.00977l-5.63672,-8.00977zM16.91406,15h3.06445l14.10742,20h-3.06445z"
                ></path>
              </g>
            </g>
          </svg>
        </a>
        <a href="https://www.linkedin.com/in/ryanwilliamsdk/">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            x="0px"
            y="0px"
            width="40"
            height="40"
            viewBox="0,0,256,256"
          >
            <g
              fill="#ffffff"
              fill-rule="nonzero"
              stroke="none"
              stroke-width="1"
              stroke-linecap="butt"
              stroke-linejoin="miter"
              stroke-miterlimit="10"
              stroke-dasharray=""
              stroke-dashoffset="0"
              font-family="none"
              font-weight="none"
              font-size="none"
              text-anchor="none"
              style="mix-blend-mode: normal"
            >
              <g transform="scale(5.12,5.12)">
                <path
                  d="M9,4c-2.74952,0 -5,2.25048 -5,5v32c0,2.74952 2.25048,5 5,5h32c2.74952,0 5,-2.25048 5,-5v-32c0,-2.74952 -2.25048,-5 -5,-5zM9,6h32c1.66848,0 3,1.33152 3,3v32c0,1.66848 -1.33152,3 -3,3h-32c-1.66848,0 -3,-1.33152 -3,-3v-32c0,-1.66848 1.33152,-3 3,-3zM14,11.01172c-1.09522,0 -2.08078,0.32736 -2.81055,0.94141c-0.72977,0.61405 -1.17773,1.53139 -1.17773,2.51367c0,1.86718 1.61957,3.32281 3.67969,3.4668c0.0013,0.00065 0.0026,0.0013 0.00391,0.00195c0.09817,0.03346 0.20099,0.05126 0.30469,0.05273c2.27301,0 3.98828,-1.5922 3.98828,-3.52148c-0.00018,-0.01759 -0.00083,-0.03518 -0.00195,-0.05274c-0.10175,-1.90023 -1.79589,-3.40234 -3.98633,-3.40234zM14,12.98828c1.39223,0 1.94197,0.62176 2.00195,1.50391c-0.01215,0.85625 -0.54186,1.51953 -2.00195,1.51953c-1.38541,0 -2.01172,-0.70949 -2.01172,-1.54492c0,-0.41771 0.15242,-0.7325 0.47266,-1.00195c0.32023,-0.26945 0.83428,-0.47656 1.53906,-0.47656zM11,19c-0.55226,0.00006 -0.99994,0.44774 -1,1v19c0.00006,0.55226 0.44774,0.99994 1,1h6c0.55226,-0.00006 0.99994,-0.44774 1,-1v-5.86523v-13.13477c-0.00006,-0.55226 -0.44774,-0.99994 -1,-1zM20,19c-0.55226,0.00006 -0.99994,0.44774 -1,1v19c0.00006,0.55226 0.44774,0.99994 1,1h6c0.55226,-0.00006 0.99994,-0.44774 1,-1v-10c0,-0.82967 0.22639,-1.65497 0.625,-2.19531c0.39861,-0.54035 0.90147,-0.86463 1.85742,-0.84766c0.98574,0.01695 1.50758,0.35464 1.90234,0.88477c0.39477,0.53013 0.61523,1.32487 0.61523,2.1582v10c0.00006,0.55226 0.44774,0.99994 1,1h6c0.55226,-0.00006 0.99994,-0.44774 1,-1v-10.73828c0,-2.96154 -0.87721,-5.30739 -2.38086,-6.89453c-1.50365,-1.58714 -3.59497,-2.36719 -5.80664,-2.36719c-2.10202,0 -3.70165,0.70489 -4.8125,1.42383v-0.42383c-0.00006,-0.55226 -0.44774,-0.99994 -1,-1zM12,21h4v12.13477v4.86523h-4zM21,21h4v1.56055c0.00013,0.43 0.27511,0.81179 0.68291,0.94817c0.40781,0.13638 0.85714,-0.00319 1.11591,-0.34661c0,0 1.57037,-2.16211 5.01367,-2.16211c1.75333,0 3.25687,0.58258 4.35547,1.74219c1.0986,1.15961 1.83203,2.94607 1.83203,5.51953v9.73828h-4v-9c0,-1.16667 -0.27953,-2.37289 -1.00977,-3.35352c-0.73023,-0.98062 -1.9584,-1.66341 -3.47266,-1.68945c-1.52204,-0.02703 -2.77006,0.66996 -3.50195,1.66211c-0.73189,0.99215 -1.01562,2.21053 -1.01562,3.38086v9h-4z"
                ></path>
              </g>
            </g>
          </svg>
        </a>
        <a href="https://bsky.app/profile/sirpandathered.bsky.social">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            x="0px"
            y="0px"
            width="40"
            height="40"
            viewBox="0,0,256,256"
          >
            <defs>
              <linearGradient
                x1="24"
                y1="5.565"
                x2="24"
                y2="40.968"
                gradientUnits="userSpaceOnUse"
                id="color-1_3ovMFy5JDSWq_gr1"
              >
                <stop offset="0" stop-color="#ffffff"></stop>
                <stop offset="1" stop-color="#ffffff"></stop>
              </linearGradient>
            </defs>
            <g
              fill="url(#color-1_3ovMFy5JDSWq_gr1)"
              fill-rule="nonzero"
              stroke="none"
              stroke-width="1"
              stroke-linecap="butt"
              stroke-linejoin="miter"
              stroke-miterlimit="10"
              stroke-dasharray=""
              stroke-dashoffset="0"
              font-family="none"
              font-weight="none"
              font-size="none"
              text-anchor="none"
              style="mix-blend-mode: normal"
            >
              <g transform="scale(5.33333,5.33333)">
                <path
                  d="M12.104,7.999c4.815,3.615 9.995,10.945 11.896,14.878c1.902,-3.933 7.081,-11.263 11.896,-14.878c3.475,-2.608 9.104,-4.626 9.104,1.796c0,1.283 -0.735,10.774 -1.167,12.315c-1.499,5.357 -6.962,6.724 -11.821,5.897c8.494,1.445 10.655,6.234 5.988,11.023c-8.863,9.094 -12.738,-2.282 -13.732,-5.197c-0.182,-0.534 -0.267,-0.784 -0.268,-0.572c-0.001,-0.213 -0.086,0.037 -0.268,0.572c-0.994,2.915 -4.869,14.291 -13.732,5.197c-4.667,-4.789 -2.506,-9.577 5.988,-11.023c-4.86,0.827 -10.323,-0.539 -11.821,-5.897c-0.432,-1.541 -1.167,-11.033 -1.167,-12.315c0,-6.422 5.63,-4.404 9.104,-1.796z"
                ></path>
              </g>
            </g>
          </svg>
        </a>
      </div>
    </div>
    <div class="text-center text-gray-400 pt-20">
      <p>&copy;Shoot.dev 2025</p>
    </div>
  </div>
  <BottomBorder :BorderInfo="borderInfo" class="h-10" />
</template>

<style scoped></style>
