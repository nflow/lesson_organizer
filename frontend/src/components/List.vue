<template>
  <div class="tw-w-full">
    <div class="tw-bg-white">
      <Draggable
        :list="conents"
        tag="table"
        animation="150"
        group="ideas"
        item-key="id"
        class="
          tw-min-w-max
          tw-w-full
          tw-table-auto
          tw-text-gray-600
          tw-text-sm
          tw-font-light
        "
      >
        <template #item="{ element }">
          <tr class="tw-border-b tw-border-gray-200 hover:tw-bg-gray-100">
            <td class="tw-py-3 tw-px-6 tw-text-left">
              <div class="tw-relative">
                <span>{{ element.value }}</span>
                <div
                  class="
                    tw-absolute tw-inset-y-0 tw-right-0 tw-flex tw-items-center
                  "
                >
                  <span
                    class="tw-cursor-pointer hover:tw-text-red-500"
                    @click="remove(element.id)"
                  >
                    <svg
                      class="tw-w-6 tw-h-6"
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
              tw-border-b tw-py-3 tw-px-6 tw-text-left tw-border-gray-200
              hover:tw-bg-gray-100
            "
          >
            <td>
              <div class="tw-relative">
                <div
                  class="
                    tw-absolute
                    tw-inset-y-0
                    tw-left-0
                    tw-pl-2
                    tw-flex
                    tw-items-center
                    tw-pointer-events-none
                  "
                >
                  <svg
                    class="tw-w-6 tw-h-6 tw-text-gray-300"
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
                    tw-inset-0
                    tw-py-3
                    tw-px-6
                    tw-block
                    tw-w-full
                    tw-pl-9
                    tw-pr-12
                    tw-border-gray-200
                    tw-hover:bg-gray-100
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
import { defineComponent, PropType, ref } from "vue";
import { ContentDto } from "@/types/content";
import { v4 as uuidv4 } from "uuid";
import Draggable from "vuedraggable";

export default defineComponent({
  name: "List",
  components: {
    Draggable,
  },
  props: {
    conents: {
      type: Object as PropType<Array<ContentDto>>,
      required: true,
    },
  },
  setup() {
    const newEntryInput = ref("");

    const addNew = (): void => {
      if (!newEntryInput.value) {
        return;
      }

      const newEntry: ContentDto = { id: uuidv4(), value: newEntryInput.value };
      //contents.value.push(newEntry);

      newEntryInput.value = "";
    };

    const remove = (id: string): ContentDto | undefined => {
      // const index: number = refIdeas.value.findIndex(
      //   (element: ContentDto) => element.id == id
      // );
      // if (index < 0) {
      //   return undefined;
      // }
      // return refIdeas.value.splice(index, 1)[0];
      return undefined;
    };

    return {
      newEntryInput,
      addNew,
      remove,
    };
  },
});
</script>
