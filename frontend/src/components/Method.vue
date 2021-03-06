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
        {{ method.title }}
      </div>
      <div
        v-if="method.category && method.category.length > 0"
        class="
          tw-rounded-t-md
          tw-pl-4
          tw-pr-4
          tw-text-gray-800
          tw-text-sm
          tw-w-full
          tw-italic
        "
      >
        {{ method.category }}
      </div>
      <div
        v-if="method.labels && method.labels.length > 0"
        class="
          tw-relative
          tw-text-gray-600
          tw-pl-2
          tw-pr-2
          tw-w-full
          tw-overflow-hidden
        "
      >
        <q-chip v-for="(v, i) in method.labels" :key="i">
          {{ resolveLabelName(v) }}
        </q-chip>
      </div>
      <div
        v-if="method.description && method.description.length > 0"
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
        {{ method.description }}
      </div>
    </div>
    <div v-if="methodId" class="tw-p-0 tw-m-0 tw-w-full">
      <content-list class="tw-p-4" :methodId="methodId" />
    </div>
  </div>
</template>

<script lang="ts">
import ContentList from "./ContentList.vue";
import { defineComponent, PropType } from "@vue/runtime-core";
import { MethodDto, resolveLabelName } from "@/types/method";

export default defineComponent({
  name: "Method",
  components: {
    ContentList,
  },
  props: {
    phaseId: {
      type: String,
      required: false,
      default: undefined,
    },
    methodId: {
      type: String,
      required: false,
      default: undefined,
    },
    method: {
      type: Object as PropType<MethodDto>,
      required: true,
    },
  },
  setup() {
    return { resolveLabelName };
  },
});
</script>
