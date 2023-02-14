<script setup>
const { title, text, img } = defineProps(["title", "text", "img"]);
</script>

<template>
  <div
    class="img-card mx-auto transition-all bg-white md:hover:scale-105 duration-200 max-w-md border shadow cursor-pointer rounded-xl relative"
  >
    <figure class="px-10 pt-10">
      <img :src="img" alt="..." class="rounded-xl duration-200" />
    </figure>
    <div class="card-body items-center text-center px-10">
      <h5 class="card-title text-black/75">{{ title }}</h5>
      <p class="text-black/50 hidden md:block">{{ text }}</p>
    </div>
    <div
      class="flex justify-center text-center items-center mb-5 sticky top-full"
    >
      <div class="card-actions w-full px-10">
        <button class="btn btn-warning text-white/90 w-full hover:opacity-70">
          Learn More
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.img-card:hover img {
  scale: 1.1;
}
</style>
