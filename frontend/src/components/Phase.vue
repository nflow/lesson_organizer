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
            :labels="element.labels"
            v-model:ideas="element.ideas"
          />
        </div>
      </template>
    </Draggable>
    <card-button @click="onAddMethod" class="mr-4 ml-4" />
    <q-dialog full-width full-height v-model="showMethodsDialog"
      ><q-card>
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">Add Method</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>
        <q-card-section>
          <q-table
            :rows="allMethods"
            :rows-per-page-options="[0]"
            :columns="methodColumns"
            hide-pagination
            grid
          >
            <template v-slot:item="props">
              <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4">
                <Method
                  @click="onMethodSelect(props.row)"
                  :title="props.row.title"
                  :description="props.row.description"
                  :labels="props.row.labels"
                />
              </div>
            </template>
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
import { BoardMethodDto, MethodDto, resolveLabelName } from "@/types/method";
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
    const onMethodSelect = (row: MethodDto): void => {
      console.log(row);
      refMethods.value.push({
        ...row,
        ideas: [],
      });
      showMethodsDialog.value = false;
    };

    const methodColumns = [
      { name: "title", label: "Title", field: "title", sortable: true },
      {
        name: "description",
        label: "Description",
        field: "description",
        sortable: true,
      },
      {
        name: "category",
        label: "Category",
        field: "category",
        sortable: true,
      },
      { name: "labels", label: "Labels", field: "labels", sortable: true },
    ];

    return {
      newEntryInput,
      refMethods,

      onAddMethod,
      showMethodsDialog,
      allMethods,
      onMethodSelect,
      resolveLabelName,
      methodColumns,
    };
  },
});
</script>
