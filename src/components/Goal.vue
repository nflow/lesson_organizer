<template>
  <div
    :style="{ backgroundColor: bgColor }"
    class="relative block align-middle text-center text-white rounded"
  >
    <div :style="{ color: bgColor }" class="fixed h-full pl-2 pt-2 opacity-80">
      <div class="bg-white flex w-6 h-6 rounded-full items-center">
        <span class="w-full">
          {{ order_id }}
        </span>
      </div>
    </div>
    <textarea
      class="
        text-shadow
        resize-y
        max-w-xs
        break-words
        bg-transparent
        align-middle
        p-2
        pl-10
        text-md
        block
        rounded-lg
        w-full
        h-full
      "
      v-model="refText"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef } from "@vue/runtime-core";

export default defineComponent({
  name: "Goal",
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
  setup(props) {
    const refText = toRef(props, "text");

    return { refText };
  },
});
</script>
