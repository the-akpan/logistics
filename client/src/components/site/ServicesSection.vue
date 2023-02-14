<script setup>
import Card from "./Card.vue";

import Air from "../../assets/imgs/slide2.jpg";
import Land from "../../assets/imgs/slide2.jpg";
import Customs from "../../assets/imgs/slide2.jpg";
import Warehouse from "../../assets/imgs/slide2.jpg";
import Train from "../../assets/imgs/slide2.jpg";
import Sea from "../../assets/imgs/slide2.jpg";

const services = [
  {
    title: "Air Freight",
    text: "Bright Time Logistics is a highly dependable airfreight forwarders that",
    img: Air,
  },
  {
    title: "Sea Freight",
    text: "Bright Time Logistics Shipping around the World First Rate Delivery is your ideal partner offering you comprehensive",
    img: Sea,
  },
  { title: "Train Freight", text: "", img: Train },
  { title: "Land Freight", text: "", img: Land },
  {
    title: "Customs Clearance",
    text: "Bright Time Logistics with its years of experience and team of global cargo agents makes sure that custom clearance for",
    img: Customs,
  },
  {
    title: "Warehousing",
    text: "We are one of the leading companies providing reliable distribution services to our clients that cover almost all",
    img: Warehouse,
  },
];
</script>

<template>
  <section class="bg-blue-50">
    <div class="container py-20">
      <h2 class="text-black/80 text-4xl text-center">Services</h2>
      <p class="text-center text-xl text-black/50">
        We offer a the best services
      </p>
      <div class="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mt-10">
        <card
          v-for="service of services"
          :title="service.title"
          :text="service.text"
          :img="service.img"
          :key="service.title"
        />
      </div>
    </div>
  </section>
</template>
