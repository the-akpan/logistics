<script setup>
import { RouterView } from "vue-router";
import { FooterPartial, NavbarPartial } from "../components/partials/home";

console.log("Rendered");
</script>

<template>
  <navbar-partial />
  <main class="mt-[88px]">
    <router-view />
  </main>
  <footer-partial />
</template>
