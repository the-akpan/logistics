<script setup>
import { ref } from "vue";
import Slide1 from "../../assets/imgs/slide1.jpg";

const tracker = ref("");
</script>

<template>
  <section>
    <div class="w-full">
      <div class="relative w-full">
        <!-- <img :src="Slide1" class="w-full" /> -->
        <div
          class="transform top-0 left-0 right-0 bottom-0 flex items-center justify-center"
          :style="{ backgroundImage: 'url(' + Slide1 + ')' }"
        >
          <div class="bg-black/75 w-full h-full">
            <div
              class="flex flex-col justify-center items-center max-w-md md:max-w-3xl text-center gap-6 mx-auto py-20"
            >
              <h2 class="text-center text-white text-3xl md:text-5xl">
                SERVICE COMES
                <span class="inline-block font-semibold text-yellow-500/75"
                  >FIRST
                </span>
                WITH US
              </h2>
              <p class="text-white/80">
                Take your business to the next level with our expertise.
                <span class="hidden md:inline-block">
                  With direct access to our global network and wide portfolio of
                  shipping services, connect with people and businesses around
                  the world.
                </span>
              </p>
              <form class="flex mt-4">
                <input
                  v-model="tracker"
                  placeholder="Enter Your Trackind ID"
                  class="px-4 md:px-6 rounded-l bg-white outline-none uppercase text-xl md:text-2xl text-black/60"
                  maxlength="14"
                  type="text"
                />
                <button
                  class="block bg-red-600 text-white py-3 px-4 md:px-6 md:py-5 rounded-r"
                >
                  TRACK
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
