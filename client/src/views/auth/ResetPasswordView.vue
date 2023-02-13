<template>
  <form class="flex flex-col gap-6">
    <h2 class="text-3xl font-semibold">Reset Password</h2>
    <input
      class="px-4 py-3 outline-none border-b border-b-gray-400 focus:border-b-blue-600"
      type="email"
      placeholder="Email"
      required
    />
    <button class="px-4 py-3 bg-blue-600 font-semibold text-white">
      Reset Password
    </button>

    <router-link
      class="mx-auto hover:underline text-black/50 hover:text-black/75"
      to="/auth/login"
      >Login</router-link
    >
  </form>
</template>
