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
  >
    <div>
      <div class="rounded-t-md p-4 text-gray-800 text-3xl font-semibold w-full">
        {{ title }}
      </div>
      <div
        v-if="description"
        class="relative h-24 p-4 mt-2 text-gray-600 w-full overflow-hidden"
      >
        <div
          class="
            absolute
            top-0
            left-0
            w-full
            h-full
            bg-gradient-to-b
            from-transparent
            to-white
          "
        />
        {{ description }}
      </div>
    </div>
    <div class="p-0 m-0 w-full">
      <list class="p-4" v-model="refIdeas" />
    </div>
  </div>
</template>

<script lang="ts">
import { ContentDto } from "@/types/ContentDto";
import List from "../components/List.vue";
import {
  computed,
  ComputedRef,
  defineComponent,
  PropType,
} from "@vue/runtime-core";

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
      type: Object as PropType<Array<ContentDto>>,
      required: true,
    },
  },
  setup(props, { emit }) {
    const refIdeas: ComputedRef<Array<ContentDto>> = computed({
      get: () => {
        return props.ideas;
      },
      set: (value) => {
        emit("update:ideas", value);
      },
    });

    return { refIdeas };
  },
});
</script>
