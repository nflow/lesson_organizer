<template>
  <div
    class="
      tw-max-w-md
      tw-p-0
      tw-bg-white
      tw-shadow-lg
      tw-rounded-lg
      tw-transition
      tw-duration-200
      hover:tw-shadow-2xl
    "
  >
    <div>
      <div
        class="
          tw-rounded-t-md
          tw-pl-4
          tw-pr-4
          tw-pt-4
          tw-text-gray-800
          tw-text-3xl
          tw-font-semibold
          tw-w-full
        "
      >
        {{ title }}
      </div>
      <div
        v-if="labels && labels.length > 0"
        class="
          tw-relative
          tw-text-gray-600
          tw-pl-2
          tw-pr-2
          tw-w-full
          tw-overflow-hidden
        "
      >
        <q-chip v-for="(v, i) in labels" :key="i">
          {{ resolveLabelName(v) }}
        </q-chip>
      </div>
      <div
        v-if="description"
        class="
          tw-relative
          tw-max-h-24
          tw-p-4
          tw-mt-2
          tw-text-gray-600
          tw-w-full
          tw-overflow-hidden
        "
      >
        <div
          class="
            tw-absolute
            tw-top-0
            tw-left-0
            tw-w-full
            tw-h-full
            tw-bg-gradient-to-b
            tw-from-transparent
            tw-to-white
          "
        />
        {{ description }}
      </div>
    </div>
    <div v-if="refIdeas" class="tw-p-0 tw-m-0 tw-w-full">
      <list class="tw-p-4" v-model="refIdeas" />
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
