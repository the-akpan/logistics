<script setup>
import { RouterLink } from "vue-router";
</script>

<template>
  <form class="flex flex-col gap-6">
    <h2 class="text-3xl font-semibold">Login</h2>
    <input
      class="px-4 py-3 outline-none border-b border-b-gray-400 focus:border-b-blue-600"
      type="email"
      placeholder="Email"
      required
    />
    <input
      class="px-4 py-3 outline-none border-b border-b-gray-400 focus:border-b-blue-600"
      type="email"
      placeholder="Password"
      required
    />
    <button class="px-4 py-3 bg-blue-600 font-semibold text-white">
      Login
    </button>

    <router-link
      class="mx-auto hover:underline text-black/50 hover:text-black/75"
      to="/auth/reset-password"
      >Reset Password</router-link
    >
  </form>
</template>
