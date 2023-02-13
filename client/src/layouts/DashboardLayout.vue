<script setup>
import { RouterView } from "vue-router";
import {
  FooterPartial,
  NavbarPartial,
  SideNavPartial,
} from "../components/partials/dashboard";
</script>

<template>
  <main class="flex relative">
    <side-nav-partial />

    <div class="w-full">
      <navbar-partial />
      <div class="container">
        <router-view />
      </div>
    </div>
  </main>
</template>
