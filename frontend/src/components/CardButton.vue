<template>
  <button
    @click="$emit('click', $event)"
    class="
      bg-gray-400
      cursor-pointer
      duration-200
      flex flex-col
      font-bold
      hover:opacity-100
      items-center
      justify-center
      opacity-40
      pl-1
      pr-1
      rounded
      text-xl text-white
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
