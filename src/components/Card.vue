<template>
  <div
    class="
      max-w-md
      p-0
      bg-white
      shadow-lg
      rounded-lg
      transition
      duration-200
      hover:shadow-2xl
    "
    :style="{ backgroundColor: color }"
  >
    <div>
      <input
        v-model="title"
        type="text"
        class="rounded-t-md p-4 text-gray-800 text-3xl font-semibold w-full"
        :style="{ backgroundColor: color }"
      />
      <div
        v-if="content"
        class="p-4 mt-2 text-gray-600 w-full"
        :style="{ backgroundColor: color }"
      >
        {{ content }}
      </div>
    </div>
    <div class="p-0 m-0 w-full">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Options } from "vue-class-component";

@Options({
  props: ["title", "content", "color"],
})
export default class HelloWorld extends Vue {}
</script>
<style scoped></style>
