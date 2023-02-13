<script setup>
import { Card } from "../../components/dashboard";
const cards = [
  { name: "Total", count: 12 },
  { name: "Pending", count: 5 },
  { name: "On hold", count: 2 },
  { name: "Delivered", count: 5 },
];
</script>

<template>
  <section class="pt-10 pb-5">
    <h3 class="text-2xl font-bold">Dashboard</h3>
    <div class="grid grid-cols-1 gap-5 mt-5 lg:grid lg:grid-cols-4">
      <card v-for="card of cards" :heading="card.name" :count="card.count" />
    </div>
  </section>
</template>
