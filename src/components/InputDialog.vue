<template>
  <div
    :style="{ backgroundColor: bgColor }"
    class="flex-initial flex flex-row text-center p-4 text-white rounded"
  >
    <div class="flex-initial text-4xl align-middle pr-4">{{ order_id }}</div>
    <div class="flex items-center">{{ text }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "InputDialog",
  props: {
    order_id: {
      type: Number,
      require: true,
    },
    text: {
      type: String,
      require: true,
    },
    bgColor: {
      type: String,
      require: true,
    },
  },
});
</script>
