<template>
  <button
    @click="$emit('click', $event)"
    class="
      bg-gradient-to-br
      border-2
      cursor-pointer
      duration-200
      flex flex-col
      font-bold
      from-gray-400
      hover:opacity-100
      items-center
      justify-center
      opacity-40
      pl-2
      pr-2
      rounded
      text-3xl text-white
      to-gray-600
      transition
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
