<script setup>
import { RouterView } from "vue-router";
import { FooterPartial, NavbarPartial } from "../components/partials/dashboard";

console.log("Rendered");
</script>

<template>
  <navbar-partial />
  <router-view />
  <footer-partial />
</template>
