<script setup>
import { RouterLink } from "vue-router";
</script>

<template>
  <nav class="py-5 border-b shadow-lg">
    <div class="container flex justify-between items-center">
      <router-link class="text-xl p-0 m-0 h-100" to="/"
        ><h1 class="text-2xl hover:underline">Go to site</h1>
      </router-link>
      <button
        class="bg-red-500 text-white py-3 px-3 rounded hover:scale-105 duration-200 active:scale-95"
      >
        Logout
      </button>
    </div>
  </nav>
</template>
