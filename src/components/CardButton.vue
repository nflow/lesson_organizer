<template>
  <button
    class="
      w-4/12
      flex flex-col
      justify-center
      items-center
      text-3xl
      font-bold
      border-2
      rounded-2xl
      border-dashed
      cursor-pointer
      transition
      duration-200
      hover:shadow-2xl
    "
  >
    <div>+</div>
    <div>Add New Card</div>
  </button>
</template>

<script lang="ts">
import { Vue, Options } from "vue-class-component";

@Options({
  props: ["title", "content", "color"],
})
export default class HelloWorld extends Vue {}
</script>
