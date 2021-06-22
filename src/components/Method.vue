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
      <div class="rounded-t-md p-4 text-gray-800 text-3xl font-semibold w-full">
        {{ title }}
      </div>
      <div v-if="description" class="p-4 mt-2 text-gray-600 w-full">
        {{ description }}
      </div>
    </div>
    <div class="p-0 m-0 w-full">
      <list v-model="refIdeas" />
    </div>
  </div>
</template>

<script lang="ts">
import { IdeaDto } from "@/types/IdeaDto";
import List from "../components/List.vue";
import { defineComponent, PropType, ref } from "@vue/runtime-core";

export default defineComponent({
  name: "Method",
  components: {
    List,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    description: {
      type: String,
      required: false,
      default: undefined,
    },
    ideas: {
      type: Object as PropType<IdeaDto>,
      required: true,
    },
  },
  setup(props): Record<string, unknown> {
    const refIdeas = ref(props.ideas);

    return { refIdeas };
  },
});
</script>
