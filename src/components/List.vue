<template>
  <div class="w-full">
    <div class="bg-white my-6">
      <table
        class="min-w-max w-full table-auto"
        @drop="onDragDrop($event)"
        @dragover.prevent
        @dragenter.prevent
      >
        <tbody class="text-gray-600 text-sm font-light">
          <tr
            v-for="(entry, index) in refModelValue.list"
            :key="index"
            class="border-b border-gray-200 hover:bg-gray-100"
            draggable="true"
            @dragstart="onDragStart($event, entry.id)"
          >
            <td class="py-3 px-6 text-left">
              <div class="relative">
                <span>{{ entry.value }}</span>
                <div class="absolute inset-y-0 right-0 flex items-center">
                  <span
                    class="cursor-pointer hover:text-red-500"
                    @click="remove(entry.id)"
                  >
                    <svg
                      class="w-6 h-6"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
                      ></path>
                    </svg>
                  </span>
                </div>
              </div>
            </td>
          </tr>
          <tr
            class="
              border-b
              py-3
              px-6
              text-left
              border-gray-200
              hover:bg-gray-100
            "
          >
            <td>
              <div class="relative">
                <div
                  class="
                    absolute
                    inset-y-0
                    left-0
                    pl-2
                    flex
                    items-center
                    pointer-events-none
                  "
                >
                  <svg
                    class="w-6 h-6 text-gray-300"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
                    ></path>
                  </svg>
                </div>
                <input
                  id="new-entry"
                  v-model="newEntryInput"
                  type="text"
                  name="new-entry"
                  class="
                    inset-0
                    py-3
                    px-6
                    block
                    w-full
                    pl-9
                    pr-12
                    border-gray-200
                    hover:bg-gray-100
                  "
                  placeholder="Eine tolle Idee ..."
                  @keydown.enter="addNew"
                />
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, {
  AppContext,
  computed,
  PropType,
  Ref,
  ref,
  SetupContext,
  toRef,
} from "vue";
import { MutationTypes, useStore } from "@/store";
import { List, ListEntry } from "@/types/List";
import { v4 as uuidv4 } from "uuid";

export default {
  name: "List",
  props: {
    modelValue: { required: true },
  },
  setup(props: any, { emit }: SetupContext): Record<string, unknown> {
    const store = useStore();

    const newEntryInput = ref("");
    const refModelValue: Ref<List> = toRef(props, "modelValue");

    const addNew = (): void => {
      if (!newEntryInput.value) {
        return;
      }

      const newEntry: ListEntry = { id: uuidv4(), value: newEntryInput.value };
      refModelValue.value.list.push(newEntry);

      newEntryInput.value = "";
    };

    const remove = (id: string): ListEntry | undefined => {
      const list = refModelValue.value.list;

      const index: number | undefined = list.findIndex(
        (element) => element.id == id
      );

      if (index < 0) {
        return undefined;
      }
      const foo = refModelValue.value.list.splice(index, 1);
      console.log(foo);
      return foo[0];
    };

    const onDragStart = (event: DragEvent, id: string): void => {
      if (event.dataTransfer) {
        store.commit(MutationTypes.SET_DRAG_CALLBACK, {
          id: id,
          callback: remove,
        });
        event.dataTransfer.dropEffect = "move";
        event.dataTransfer.effectAllowed = "move";
      }
    };

    const onDragDrop = (event: DragEvent): void => {
      const dragedEntry = store.state.dragCallback;
      if (dragedEntry) {
        const entry: ListEntry = dragedEntry.callback(dragedEntry.id);
        console.log(entry);
        refModelValue.value.list.push(entry);
      }
    };

    return {
      newEntryInput,
      addNew,
      remove,
      refModelValue,
      onDragStart,
      onDragDrop,
    };
  },
};
</script>
