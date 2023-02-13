<script setup>
import { ref } from "vue";
import { RouterLink } from "vue-router";

import dashboard from "../../../assets/svgs/dashboard.svg";
import orders from "../../../assets/svgs/orders.svg";
import profile from "../../../assets/svgs/profile.svg";

const navs = [
  { link: "/dashboard", name: "Dashboard", logo: dashboard },
  { link: "/dashboard/orders", name: "Orders", logo: orders },
  { link: "/dashboard/profile", name: "Profile", logo: profile },
];

const showMenu = ref(true);

function toggleMenu() {
  showMenu.value = !showMenu.value;
}

const lastIndex = navs.length - 1;
</script>

<template>
  <section
    class="fixed top-0 left-0 h-screen z-10 flex bg-black/50 transition-all duration-200"
    :class="showMenu ? 'w-screen' : 'w-0'"
  >
    <nav class="relative w-1/2 h-screen bg-blue-700 z-30 py-10">
      <ul
        class="relative transition duration-200 h-full overflow-scroll"
        :class="showMenu ? 'block' : 'hidden'"
      >
        <li class="mb-10">
          <div class="px-4 flex">
            <router-link to="/dashboard">
              <img class="h-7 rotate-45" :src="dashboard" alt="" />
            </router-link>
          </div>
        </li>
        <li v-for="(nav, index) of navs" :key="nav.link">
          <router-link
            class="flex items-center px-4 py-3 w-full text-white/75 hover:bg-blue-800"
            :class="index !== lastIndex && 'mb-3'"
            :to="nav.link"
            exact-active-class="bg-blue-800 text-white/100 border-r-2 border-blue-900"
            ><img class="mr-4 w-8 h-8" :src="nav.logo" alt="" />{{
              nav.name
            }}</router-link
          >
        </li>
        <li class="absolute bottom-0 w-full">
          <a
            class="flex items-center px-4 py-3 w-full text-white hover:bg-red-800"
            ><img
              class="mr-4 w-8 h-8"
              src="../../../assets/svgs/logout.svg"
              alt=""
            />Logout</a
          >
        </li>
      </ul>
      <div
        class="absolute top-0 mt-24 z-10 -right-10 p-1 cursor-pointer border rounded flex items-center justify-center bg-blue-600/50"
        @click="toggleMenu"
        :class="showMenu ? '-right-10' : 'right-10'"
      >
        <svg
          v-if="showMenu"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          aria-hidden="true"
          class="h-6 w-6 text-white hover:scale-110"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M6 18L18 6M6 6l12 12"
          ></path>
        </svg>
        <svg
          v-else
          class="h-6 w-6 hover:scale-110"
          viewBox="0 0 16 16"
          stroke-width="1.5"
          aria-hidden="true"
          xmlns="http://www.w3.org/2000/svg"
        >
          <g fill="#ffffff">
            <path d="m 1 2 h 14 v 2 h -14 z m 0 0" />
            <path d="m 1 7 h 14 v 2 h -14 z m 0 0" />
            <path d="m 1 12 h 14 v 2 h -14 z m 0 0" />
          </g>
        </svg>
      </div>
    </nav>
  </section>
</template>

<style scoped>
li:last-of-type {
  position: sticky;
  top: 100%;
}
</style>
