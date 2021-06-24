<template>
  <div
    style="min-width: 20rem"
    class="flex-auto flex flex-col hover:bg-gray-300"
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
    <div class="flex-1 flex flex-col p-4 space-y-10 items-start self-center">
      <Method
        v-for="method in refMethods"
        :key="method.id"
        :title="method.title"
        :description="method.description"
        :ideas="method.ideas"
        draggable="true"
        @dragstart="onDragStart($event, method.id)"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, Ref, ref, toRef } from "vue";
import { MutationTypes, useStore } from "@/store";
import { MethodDto } from "@/types/MethodDto";
import Method from "../components/Method.vue";

export default defineComponent({
  name: "Phase",
  components: {
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

    const remove = (id: string): MethodDto | undefined => {
      const list = refMethods.value;

      const index: number | undefined = list.findIndex(
        (element) => element.id == id
      );

      if (index < 0) {
        return undefined;
      }

      return refMethods.value.splice(index, 1)[0];
    };

    const onDragStart = (event: DragEvent, id: string): void => {
      if (event.dataTransfer) {
        store.commit(MutationTypes.SET_DRAG_CALLBACK_METHOD, {
          id: id,
          callback: remove,
        });
        event.dataTransfer.dropEffect = "move";
        event.dataTransfer.effectAllowed = "move";
      }
    };

    const onDragDrop = (event: DragEvent): void => {
      const dragedEntry = store.state.dragCallbackMethod;
      if (dragedEntry) {
        const entry: MethodDto = dragedEntry.callback(dragedEntry.id);
        store.commit(MutationTypes.SET_DRAG_CALLBACK_METHOD, undefined);
        refMethods.value.push(entry);
        event.preventDefault();
      }
    };

    return {
      onDragStart,
      onDragDrop,
      newEntryInput,
      remove,
      refMethods,
    };
  },
});
</script>
