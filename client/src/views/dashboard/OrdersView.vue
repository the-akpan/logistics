<script setup>
import { TableFilters } from "../../components/dashboard";
const orders = [
  {
    tracker: "123456",
    recipient: "Random person",
    to: "China",
    from: "Russia",
    createdAt: new Date(),
    status: "On Hold",
  },
  {
    tracker: "123456",
    recipient: "Random person",
    to: "China",
    from: "Russia",
    createdAt: new Date(),
    status: "On Hold",
  },
  {
    tracker: "123456",
    recipient: "Random person",
    to: "China",
    from: "Russia",
    createdAt: new Date(),
    status: "On Hold",
  },
  {
    tracker: "123456",
    recipient: "Random person",
    to: "China",
    from: "Russia",
    createdAt: new Date(),
    status: "On Hold",
  },
];
</script>

<template>
  <section class="pt-10 pb-5">
    <h3 class="text-2xl font-bold">Orders</h3>
    <table-filters />
    <table class="table-auto w-full text-left mt-5">
      <thead class="font-bold uppercase bg-gray-100">
        <tr>
          <th scope="col" class="px-6 py-3">Tracker</th>
          <th scope="col" class="px-6 py-3">Recipient</th>
          <th scope="col" class="px-6 py-3">From</th>
          <th scope="col" class="px-6 py-3">To</th>
          <th scope="col" class="px-6 py-3">Status</th>
          <th scope="col" class="px-6 py-3">Created At</th>
          <th scope="col" class="px-6 py-3">Actions</th>
        </tr>
      </thead>
      <tbody class="">
        <tr
          class="&:nth-child(odd):bg-gray-50 border-b border-gray-300"
          v-for="order of orders"
          :key="order.tracker"
        >
          <th scope="row" class="px-6 py-4">{{ order.tracker }}</th>
          <td class="px-6 py-4">{{ order.recipient }}</td>
          <td class="px-6 py-4">{{ order.from }}</td>
          <td class="px-6 py-4">{{ order.to }}</td>
          <td class="px-6 py-4">{{ order.status }}</td>
          <td class="px-6 py-4">{{ order.createdAt.toDateString() }}</td>
          <td class="px-6 py-4">
            <div class="flex justify-between gap-3">
              <button class="bg-yellow-400 p-2 rounded text-white">Edit</button>
              <button class="bg-red-500 p-2 rounded text-white">Delete</button>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </section>
</template>
