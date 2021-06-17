<template>
  <div class="w-full lg:w-5/6">
    <div class="bg-white shadow-md rounded my-6">
      <table class="min-w-max w-full table-auto">
        <tbody class="text-gray-600 text-sm font-light">
          <tr
            v-for="(entry, index) in refModelValue"
            :key="index"
            class="border-b border-gray-200 hover:bg-gray-100"
          >
            <td class="py-3 px-6 text-left">
              <div class="relative">
                <span>{{ entry }}</span>
                <div class="absolute inset-y-0 right-0 flex items-center">
                  <span
                    class="cursor-pointer hover:text-red-500"
                    @click="remove(index)"
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
                  @keydown.enter="addNew"
                  v-model="newEntryInput"
                  type="text"
                  name="new-entry"
                  id="new-entry"
                  class="
                    inset-0
                    py-3
                    px-6
                    block
                    w-full
                    pl-9
                    pr-12
                    sm:text-sm
                    border-gray-200
                    hover:bg-gray-100
                  "
                  placeholder="Eine tolle Idee ..."
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
import Vue, { computed, PropType, ref, toRef } from "vue";

export default {
  name: "List",
  props: {
    modelValue: { type: Array, required: true },
  },
  setup(props: any, { emit }: any): Record<string, unknown> {
    const newEntryInput = ref("");
    const refModelValue = toRef(props, "modelValue");

    const addNew = (): void => {
      if (!newEntryInput.value) {
        return;
      }

      refModelValue.value.push(newEntryInput.value);

      newEntryInput.value = "";
    };

    const remove = (index: number): void => {
      const copyList = Array.from(props.modelValue);
      refModelValue.value.splice(index, 1);
      //emit("update:modelValue");
    };

    return { newEntryInput, addNew, remove, refModelValue };
  },
};
</script>
