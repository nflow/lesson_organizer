<template>
  <div
    style="min-width: 20rem"
    class="flex flex-col hover:bg-gray-300"
    @drop="onDragDrop"
    @dragover.prevent
    @dragenter.prevent
  >
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
    <div class="flex-auto">
      <Draggable
        group="method"
        v-model="refMethods"
        fallbackOnBody="true"
        swapThreshold="0.65"
        animation="150"
        item-key="id"
        class="flex flex-col p-4 space-y-10 items-start self-center"
      >
        <template #item="{ element }">
          <Method
            :title="element.title"
            :description="element.description"
            :ideas="element.ideas"
          />
        </template>
      </Draggable>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, Ref, ref, toRef } from "vue";
import { MutationTypes, useStore } from "@/store";
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
  setup(props) {
    const store = useStore();

    const newEntryInput = ref("");
    const refMethods: Ref<Array<MethodDto>> = toRef(props, "methods");
    return {
      newEntryInput,
      refMethods,
    };
  },
});
</script>
