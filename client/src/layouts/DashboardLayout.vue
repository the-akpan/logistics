<script setup>
import { RouterView } from "vue-router";
import {
  FooterPartial,
  NavbarPartial,
  SideNavPartial,
} from "../components/partials/dashboard";
</script>

<template>
  <navbar-partial />
  <side-nav-partial />
  <main class="py-5">
    <div class="container">
      <router-view />
    </div>
  </main>
  <footer-partial />
</template>
