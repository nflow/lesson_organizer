<template>
  <button
    @click="$emit('click', $event)"
    class="
      flex flex-col
      justify-center
      items-center
      bg-gradient-to-br
      from-gray-400
      to-gray-600
      text-3xl
      font-bold
      opacity-40
      border-2
      text-white
      rounded
      cursor-pointer
      transition
      duration-200
      hover:opacity-100
      colo
      p-2
    "
  >
    <div>+</div>
  </button>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "CardButton",
});
</script>
