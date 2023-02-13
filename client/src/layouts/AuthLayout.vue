<script setup>
import { RouterView } from "vue-router";
</script>

<template>
  <section class="h-screen w-full flex items-center justify-center bg-blue-50">
    <div class="bg-white shadow rounded p-10">
      <router-view />
    </div>
  </section>
</template>
