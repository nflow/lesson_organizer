<template>
  <Draggable
    group="method"
    v-model="refMethods"
    fallbackOnBody="true"
    swapThreshold="0.65"
    animation="150"
    item-key="id"
    style="min-width: 20rem"
    class="flex flex-col hover:bg-gray-300"
  >
    <template #header>
      <div
        class="
          flex-initial
          font-extrabold
          hover:bg-green-500
          bg-green-400
          p-4
          text-center
        "
      >
        {{ title }}
      </div>
    </template>
    <template #item="{ element }">
      <div class="flex flex-col p-4 space-y-10">
        <Method
          :title="element.title"
          :description="element.description"
          v-model:ideas="element.ideas"
        />
      </div>
    </template>
  </Draggable>
</template>

<script lang="ts">
import { computed, ComputedRef, defineComponent, PropType, ref } from "vue";
import { MethodDto } from "@/types/MethodDto";
import Method from "../components/Method.vue";
import Draggable from "vuedraggable";

export default defineComponent({
  name: "Phase",
  components: {
    Draggable,
    Method,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    methods: {
      type: Object as PropType<Array<MethodDto>>,
      required: true,
    },
  },
  setup(props, { emit }) {
    const newEntryInput = ref("");
    const refMethods: ComputedRef<Array<MethodDto>> = computed({
      get: () => {
        return props.methods;
      },
      set: (value) => {
        emit("update:methods", value);
      },
    });

    return {
      newEntryInput,
      refMethods,
    };
  },
});
</script>
