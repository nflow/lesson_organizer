<template>
  <div class="w-full">
    <div class="bg-white">
      <Draggable
        v-model="refIdeas"
        tag="table"
        animation="150"
        group="ideas"
        item-key="id"
        class="min-w-max w-full table-auto text-gray-600 text-sm font-light"
      >
        <template #item="{ element }">
          <tr class="border-b border-gray-200 hover:bg-gray-100">
            <td class="py-3 px-6 text-left">
              <div class="relative">
                <span>{{ element.value }}</span>
                <div class="absolute inset-y-0 right-0 flex items-center">
                  <span
                    class="cursor-pointer hover:text-red-500"
                    @click="remove(element.id)"
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
                      />
                    </svg>
                  </span>
                </div>
              </div>
            </td>
          </tr>
        </template>
        <template #footer>
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
                    />
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
        </template>
      </Draggable>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, ComputedRef, defineComponent, PropType, ref } from "vue";
import { IdeaDto } from "@/types/IdeaDto";
import { v4 as uuidv4 } from "uuid";
import Draggable from "vuedraggable";

export default defineComponent({
  name: "List",
  components: {
    Draggable,
  },
  props: {
    modelValue: {
      type: Object as PropType<Array<IdeaDto>>,
      required: true,
    },
  },
  setup(props, { emit }) {
    const newEntryInput = ref("");
    const refIdeas: ComputedRef<Array<IdeaDto>> = computed({
      get: () => {
        return props.modelValue;
      },
      set: (value) => {
        emit("update:modelValue", value);
      },
    });

    const addNew = (): void => {
      if (!newEntryInput.value) {
        return;
      }

      const newEntry: IdeaDto = { id: uuidv4(), value: newEntryInput.value };
      refIdeas.value.push(newEntry);

      newEntryInput.value = "";
    };

    const remove = (id: string): IdeaDto | undefined => {
      const index: number = refIdeas.value.findIndex(
        (element: IdeaDto) => element.id == id
      );

      if (index < 0) {
        return undefined;
      }

      return refIdeas.value.splice(index, 1)[0];
    };

    return {
      newEntryInput,
      addNew,
      remove,
      refIdeas,
    };
  },
});
</script>
