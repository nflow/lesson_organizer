<template>
  <draggable
    :id="phaseId"
    group="method"
    v-model="phaseMethods"
    @end="onUpdateMethod"
    fallbackOnBody="true"
    swapThreshold="0.65"
    animation="150"
    item-key="id"
    class="tw-w-96 tw-flex tw-flex-col hover:tw-bg-gray-300"
    delay="100"
    delayOnTouchOnly="true"
  >
    <template #header>
      <div
        class="
          tw-relative tw-flex-initial tw-font-extrabold
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
        {{ phase.title }}
        <span
          class="
            tw-absolute
            tw-m-auto
            tw-top-0
            tw-bottom-0
            tw-right-1
            tw-cursor-pointer
            hover:tw-text-red-400
          "
          @click="removePhase()"
        >
          <svg
            class="tw-w-4 tw-h-full"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
            ></path>
          </svg>
        </span>
      </div>
    </template>
    <template #item="{ element }">
      <div :id="element.id" class="tw-relative tw-flex tw-flex-col tw-p-2">
        <span
          class="
            tw-bg-white
            tw-shadow-md
            tw-p-1
            tw-rounded-md
            tw-absolute
            tw-top-1
            tw-right-1
            tw-cursor-pointer
            hover:tw-text-red-400
          "
          @click="removeMethod(element)"
        >
          <svg
            class="tw-w-4 tw-h-full"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
            ></path>
          </svg>
        </span>
        <method :methodId="element.id" :method="element.method" />
      </div>
    </template>
    <template #footer>
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
              :rows="allMethods.data.value"
              :rows-per-page-options="[0]"
              :columns="methodColumns"
              hide-pagination
              grid
            >
              <template v-slot:item="props">
                <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4">
                  <method
                    @click="onMethodSelect(props.row)"
                    :method="props.row"
                  />
                </div>
              </template>
            </q-table>
          </q-card-section>
        </q-card>
      </q-dialog>
    </template>
  </draggable>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref } from "vue";
import Method from "../components/Method.vue";
import CardButton from "../components/CardButton.vue";
import Draggable from "vuedraggable";
import { BoardMethodDto, MethodDto, resolveLabelName } from "@/types/method";
import { getMethods } from "@/api";
import {
  deleteMethod,
  deletePhase,
  getPhaseMethods,
  postMethodAssociation,
  putMethodOrder,
} from "@/api/board";
import { useQueryClient } from "vue-query";
import { PhaseDto } from "@/types/phase";

export default defineComponent({
  name: "Phase",
  components: {
    Draggable,
    Method,
    CardButton,
  },
  props: {
    boardId: {
      type: String,
      required: true,
    },
    phaseId: {
      type: String,
      required: true,
    },
    phase: {
      type: Object as PropType<PhaseDto>,
      required: true,
    },
  },
  setup(props) {
    const queryClient = useQueryClient();
    const allMethods = getMethods();
    const updateMethod = putMethodOrder(props.phaseId);
    const retrievePhaseMethods = getPhaseMethods(props.phaseId);
    const phaseMethods = computed({
      get: () => {
        return retrievePhaseMethods.data.value;
      },
      set: (methods) => {
        queryClient.setQueryData(["phase_methods", props.phaseId], methods);
      },
    });

    const onUpdateMethod = (evt: any) => {
      if (
        evt.newDraggableIndex == evt.oldDraggableIndex &&
        evt.to.id == evt.from.id
      ) {
        return;
      }

      let afterId = undefined;
      if (evt.newDraggableIndex > 0) {
        const methods = queryClient.getQueryData<BoardMethodDto[]>([
          "phase_methods",
          evt.to.id,
        ]);
        if (methods) {
          afterId = methods[evt.newDraggableIndex - 1].id;
        }
      }

      updateMethod.mutate(
        {
          methodId: evt.item.id,
          parentPhaseId: evt.from.id != evt.to.id ? evt.to.id : undefined,
          afterMethodId: afterId,
        },
        {
          onSuccess: () => {
            // TODO: Could return the new list of phases .. or maybe not.
            queryClient.invalidateQueries(["phase_methods", evt.to.id]);
            if (evt.from.id != evt.to.id) {
              queryClient.invalidateQueries(["phase_methods", evt.from.id]);
            }
          },
        }
      );
    };

    const associateMethod = postMethodAssociation(props.phaseId);
    let showMethodsDialog = ref(false);
    const onAddMethod = (): void => {
      showMethodsDialog.value = true;
    };
    const onMethodSelect = (row: MethodDto): void => {
      associateMethod.mutate(
        { id: row.id },
        {
          onSuccess: () => {
            queryClient.invalidateQueries(["phase_methods", props.phaseId]);
            showMethodsDialog.value = false;
          },
        }
      );
    };

    const removeMethodMutation = deleteMethod();
    const removeMethod = (element: MethodDto): void => {
      removeMethodMutation.mutate([props.boardId, props.phaseId, element.id], {
        onSuccess: () => {
          queryClient.invalidateQueries(["phase_methods", props.phaseId]);
        },
      });

      const newList = [...retrievePhaseMethods.data.value];
      newList.splice(newList.indexOf(element), 1);
      queryClient.setQueryData(["phase_methods", props.phaseId], newList);
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
      phaseMethods,
      onAddMethod,
      onUpdateMethod,
      showMethodsDialog,
      allMethods,
      onMethodSelect,
      resolveLabelName,
      methodColumns,
      removeMethod,
    };
  },
});
</script>
