<template>
  <div class="flex flex-col hover:bg-gray-300">
    <Draggable
      group="method"
      v-model="refMethods"
      fallbackOnBody="true"
      swapThreshold="0.65"
      animation="150"
      item-key="id"
      style="min-width: 20rem"
      class="flex flex-col"
      delay="60"
      delayOnTouchOnly="true"
    >
      <template #header>
        <div
          class="
            flex-initial
            font-extrabold
            hover:bg-orange
            text-white text-s
            bg-yellow-orange
            rounded
            p-2
            m-2
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
    <card-button @click="onAddMethod" class="mr-4 ml-4" />
    <q-dialog full-width full-height v-model="showMethodsDialog"
      ><q-card>
        <q-card-section>
          <q-table @row-click="onMethodSelect" :rows="allMethods" grid>
          </q-table>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { computed, ComputedRef, defineComponent, PropType, ref } from "vue";
import Method from "../components/Method.vue";
import CardButton from "../components/CardButton.vue";
import Draggable from "vuedraggable";
import { BoardMethodDto, MethodDto } from "@/types/method";
import { useStore } from "vuex";
import { ApiActionTypes } from "@/store/modules/api/action-types";

export default defineComponent({
  name: "Phase",
  components: {
    Draggable,
    Method,
    CardButton,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    methods: {
      type: Object as PropType<Array<BoardMethodDto>>,
      required: true,
    },
  },
  setup(props, { emit }) {
    const store = useStore();
    const newEntryInput = ref("");
    const refMethods: ComputedRef<Array<BoardMethodDto>> = computed({
      get: () => {
        return props.methods;
      },
      set: (value) => {
        emit("update:methods", value);
      },
    });

    let showMethodsDialog = ref(false);
    const onAddMethod = (): void => {
      store.dispatch(ApiActionTypes.FETCH_METHODS);
      showMethodsDialog.value = true;
    };
    const allMethods = computed(() => store.getters.allMethods);
    const onMethodSelect = (a: any, row: MethodDto, b: any): void => {
      refMethods.value.push(row as BoardMethodDto);
      showMethodsDialog.value = false;
    };

    return {
      newEntryInput,
      refMethods,

      onAddMethod,
      showMethodsDialog,
      allMethods,
      onMethodSelect,
    };
  },
});
</script>
