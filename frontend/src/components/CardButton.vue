<template>
  <button
    @click="$emit('click', $event)"
    class="
      tw-bg-gray-400
      tw-cursor-pointer
      tw-duration-200
      tw-flex
      tw-flex-col
      tw-font-bold
      tw-hover:opacity-100
      tw-items-center
      tw-justify-center
      tw-opacity-40
      tw-pl-1
      tw-pr-1
      tw-rounded
      tw-text-xl
      tw-text-white
      tw-transition
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
