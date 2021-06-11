<template>
  <div
    class="
      max-w-md
      py-4
      px-8
      bg-white
      shadow-lg
      rounded-lg
      transition
      duration-200
      hover:shadow-2xl
    "
    v-bind:style="{ backgroundColor: color }"
  >
    <div>
      <input
        type="text"
        class="text-gray-800 text-3xl font-semibold w-full"
        v-model="title"
        v-bind:style="{ backgroundColor: color }"
      />
      <textarea
        v-if="content"
        class="mt-2 text-gray-600 w-full"
        v-model="content"
        v-bind:style="{ backgroundColor: color }"
      />
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
