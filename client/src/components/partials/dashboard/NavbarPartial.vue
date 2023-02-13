<script setup>
import { RouterLink } from "vue-router";
</script>

<template>
  <nav class="py-5 border-b shadow-lg w-full">
    <div class="container w-full flex justify-end items-center">
      <button
        class="flex items-center text-black/75 py-3 px-3 opacity-75 hover:opacity-100 rounded hover:underline duration-200 active:scale-95"
      >
        <img
          class="mr-4 w-6 h-6"
          src="../../../assets/svgs/logout.svg"
          alt=""
        />
        Logout
      </button>
    </div>
  </nav>
</template>
