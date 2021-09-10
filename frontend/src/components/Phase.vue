<template>
  <div class="tw-flex tw-flex-col hover:tw-bg-gray-300">
    <draggable
      :id="boardPhase.id"
      group="method"
      v-model="phaseMethods"
      @end="onUpdateMethod"
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
          {{ boardPhase.phase.title }}
        </div>
      </template>
      <template #item="{ element }">
        <div :id="element.id" class="tw-flex tw-flex-col tw-p-2">
          <method
            :boardId="boardId"
            :phaseId="boardPhase.id"
            :methodId="element.id"
            :method="element.method"
            :contents="element.contents"
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
  </div>
</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  getCurrentInstance,
  PropType,
  ref,
  toRef,
  toRefs,
} from "vue";
import Method from "../components/Method.vue";
import CardButton from "../components/CardButton.vue";
import Draggable from "vuedraggable";
import { MethodDto, resolveLabelName } from "@/types/method";
import { getMethods } from "@/api";
import { postMethodAssociation, putMethodOrder } from "@/api/board";
import { BoardPhaseDto } from "@/types/phase";
import { useQueryClient } from "vue-query";
import { BoardDto } from "@/types/board";
import { Mutex } from "async-mutex";

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
    boardPhase: {
      type: Object as PropType<BoardPhaseDto>,
      required: true,
    },
  },
  setup(props) {
    const queryClient = useQueryClient();
    const allMethods = getMethods();
    const updateMethod = putMethodOrder(props.boardId, props.boardPhase.id);
    const mutex: Mutex =
      getCurrentInstance()?.appContext.config.globalProperties
        .boardCacheUpdateMutex;
    const phaseMethods = computed({
      get: () => {
        return props.boardPhase.methods;
      },
      set: async (newMethodOrder) => {
        const currentBoard = queryClient.getQueryData<BoardDto>("board");
        console.log(currentBoard);
        if (!currentBoard) {
          return;
        }
        const phase = currentBoard.phases.find(
          (e) => e.id == props.boardPhase.id
        );

        if (!phase) {
          return;
        }
        phase.methods = newMethodOrder;
        queryClient.setQueryData("board", currentBoard);
        return;
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
        const board = queryClient.getQueryData<BoardDto>("board");
        if (board) {
          const phase = board.phases.find((e) => e.id == evt.to.id);
          if (phase) {
            afterId = phase.methods[evt.newDraggableIndex - 1].id;
          }
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
            queryClient.invalidateQueries("board");
          },
        }
      );
    };

    const associateMethod = postMethodAssociation(
      ref(props.boardId),
      props.boardPhase.id
    );
    let showMethodsDialog = ref(false);
    const onAddMethod = (): void => {
      showMethodsDialog.value = true;
    };
    const onMethodSelect = (row: MethodDto): void => {
      associateMethod.mutate(
        { id: row.id },
        {
          onSuccess: () => {
            queryClient.invalidateQueries("board");
            showMethodsDialog.value = false;
          },
        }
      );
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
    };
  },
});
</script>
