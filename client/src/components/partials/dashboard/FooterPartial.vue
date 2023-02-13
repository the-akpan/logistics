<script setup></script>
<template>
  <footer class="py-5 bg-blue-600 text-white">
    <div
      class="container flex flex-col gap-5 md:flex-row justify-between items-center"
    >
      <p class="opacity-75">
        Copyright &#169; {{ new Date().getFullYear() }} Bright Time
      </p>
      <ul class="flex gap-5">
        <li>
          <a
            class="hover:underline opacity-75 hover:opacity-100"
            href="https://twitter.com"
            >Twitter</a
          >
        </li>
        <li>
          <a
            class="hover:underline opacity-75 hover:opacity-100"
            href="https://instagram.com"
            >Instagram</a
          >
        </li>
        <li>
          <a
            class="hover:underline opacity-75 hover:opacity-100"
            href="https://facebook.com"
            >Facebook</a
          >
        </li>
      </ul>
    </div>
  </footer>
</template>
