<script setup>
import { FooterPartial, NavbarPartial } from "../components/partials/home";

console.log("Rendered");
</script>

<template>
  <navbar-partial />
  <slot />
  <footer-partial />
</template>
