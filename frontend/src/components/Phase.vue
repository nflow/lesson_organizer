<template>
  <div class="tw-flex tw-flex-col hover:tw-bg-gray-300">
    <draggable
      group="method"
      v-model="refMethods"
      fallbackOnBody="true"
      swapThreshold="0.65"
      animation="150"
      item-key="id"
      style="min-width: 20rem"
      class="tw-flex tw-flex-col"
      delay="100"
      delayOnTouchOnly="true"
    >
      <template #header>
        <div
          class="
            tw-flex-initial tw-font-extrabold
            tw-hover:bg-orange
            tw-text-white
            tw-text-sm
            tw-bg-yellow-orange
            tw-rounded
            tw-p-1
            tw-m-1
            tw-text-center
          "
        >
          {{ title }}
        </div>
      </template>
      <template #item="{ element }">
        <div class="tw-flex tw-flex-col tw-p-2">
          <method
            :title="element.title"
            :category="element.category"
            :description="element.description"
            :labels="element.labels"
            v-model:ideas="element.ideas"
          />
        </div>
      </template>
    </draggable>
    <card-button @click="onAddMethod" class="tw-m-2" />
    <q-dialog full-width full-height v-model="showMethodsDialog"
      ><q-card>
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">Add Method</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>
        <q-card-section>
          <q-table
            :rows="allMethods.value"
            :rows-per-page-options="[0]"
            :columns="methodColumns"
            hide-pagination
            grid
          >
            <template v-slot:item="props">
              <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4">
                <method
                  @click="onMethodSelect(props.row)"
                  :title="props.row.title"
                  :category="props.row.category"
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
import { getMethods } from "@/api";

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
    const allMethods = getMethods();

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
      showMethodsDialog.value = true;
    };
    const onMethodSelect = (row: MethodDto): void => {
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
