<script setup>
const statuses = [
  "on-hold",
  "created",
  "pending",
  "canceled",
  "delivered",
  "in-transit",
];
</script>

<template>
  <div class="flex items-center justify-between mt-5">
    <div>
      <label for="status">Status:</label>
      <select name="status" class="ml-2 p-2 rounded">
        <option
          v-for="status of statuses"
          class="uppercase"
          :value="status"
          :key="status"
        >
          {{ status.toLocaleUpperCase() }}
        </option>
      </select>
    </div>
    <button class="bg-blue-600 px-5 py-2 text-white rounded">Apply</button>
  </div>
</template>
