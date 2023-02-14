<script setup>
import { RouterView } from "vue-router";
import { FooterPartial, NavbarPartial } from "../components/partials/home";

console.log("Rendered");
</script>

<template>
  <navbar-partial />
  <div class="mt-[88px]">
    <router-view />
  </div>
  <footer-partial />
</template>
