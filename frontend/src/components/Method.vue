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
      <div
        class="
          rounded-t-md
          pl-4
          pr-4
          pt-4
          text-gray-800 text-3xl
          font-semibold
          w-full
        "
      >
        {{ title }}
      </div>
      <div
        v-if="labels && labels.length > 0"
        class="relative text-gray-600 pl-2 pr-2 w-full overflow-hidden"
      >
        <q-chip v-for="(v, i) in labels" :key="i">
          {{ resolveLabelName(v) }}
        </q-chip>
      </div>
      <div
        v-if="description"
        class="relative max-h-24 p-4 mt-2 text-gray-600 w-full overflow-hidden"
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
    <div v-if="refIdeas" class="p-0 m-0 w-full">
      <list class="p-4" v-model="refIdeas" />
    </div>
  </div>
</template>

<script lang="ts">
import { ContentDto } from "@/types/content";
import List from "../components/List.vue";
import {
  computed,
  ComputedRef,
  defineComponent,
  PropType,
} from "@vue/runtime-core";
import { Label, resolveLabelName } from "@/types/method";

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
    labels: {
      type: Object as PropType<Label[]>,
      required: true,
    },
    ideas: {
      type: Object as PropType<Array<ContentDto>>,
      required: false,
      default: undefined,
    },
  },
  setup(props, { emit }) {
    const refIdeas = computed({
      get: () => {
        return props.ideas;
      },
      set: (value) => {
        emit("update:ideas", value);
      },
    });

    return { refIdeas, resolveLabelName };
  },
});
</script>
