<script setup>
import { ref } from "vue";
import { RouterLink } from "vue-router";

const links = [
  { name: "index", text: "Home" },
  { name: "services", text: "Services" },
  { name: "about", text: "About" },
  { name: "contact", text: "Contact Us" },
];

const showMenu = ref(false);

function toggleMenu() {
  showMenu.value = !showMenu.value;
}
</script>

<template>
  <nav class="bg-white shadow border-b fixed h-[88px] top-0 left-0 right-0">
    <div
      class="container flex flex-col lg:flex-row items-center justify-between"
    >
      <div
        class="w-full text-white flex justify-between items-center py-5 lg:w-20"
      >
        <div class="flex items-center justify-between w-full lg:w-auto">
          <router-link
            to="/"
            class="text-2xl text-white/90 p-2 rounded-full bg-red-500"
            >B<span class="text-yellow-500">T</span>L</router-link
          >
          <div class="cursor-pointer lg:hidden" @click="toggleMenu">
            <svg
              v-if="showMenu"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              stroke-width="1.5"
              stroke="currentColor"
              aria-hidden="true"
              class="h-7 w-7 text-black hover:scale-110"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M6 18L18 6M6 6l12 12"
              ></path>
            </svg>
            <svg
              v-else
              class="h-6 w-6 text-black hover:scale-110"
              viewBox="0 0 16 16"
              stroke-width="1.5"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="m 1 2 h 14 v 2 h -14 z m 0 0" />
              <path d="m 1 7 h 14 v 2 h -14 z m 0 0" />
              <path d="m 1 12 h 14 v 2 h -14 z m 0 0" />
            </svg>
          </div>
        </div>
      </div>
      <div
        class="lg:block w-full duration-200 transition-all relative"
        :class="showMenu ? 'block' : 'hidden'"
      >
        <ul
          class="absolute w-full lg:static bg-blue-500 md:bg-white rounded-lg md:rounded-none border lg:border-none shadow md:shadow-lg lg:shadow-none flex flex-col md:flex-row justify-center md:justify-around mt-5 lg:mt-0 lg:justify-end items-center py-5 md:py-3 gap-5 md:gap-8"
        >
          <li class="" v-for="link of links" :key="link.name">
            <router-link
              class="relative uppercase text-white md:text-black font-normal nav-link after:bg-white md:after:bg-blue-500"
              exact-active-class="active"
              :to="{ name: link.name }"
              >{{ link.text }}</router-link
            >
          </li>
          <li class="flex">
            <router-link class="w-full md:w-100" :to="{ name: 'tracking' }">
              <button
                class="hover:shadow bg-white md:bg-blue-500 py-2 md:py-3 px-5 text-blue-500 md:text-white rounded-lg uppercase hover:scale-105 active:scale-95 duration-200"
              >
                Track
              </button>
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.nav-link::after {
  transition: all 0.15s ease-in;
}
.nav-link::after,
.active::after {
  content: "";
  width: 0%;
  height: 3px;
  position: absolute;
  bottom: -8px;
  left: 0;
}

.active::after,
.nav-link:hover::after {
  width: 100%;
}
</style>
