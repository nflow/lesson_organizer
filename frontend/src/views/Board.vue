<template>
  <div class="tw-text-5xl" v-if="board.isLoading.value">Loading board ...</div>
  <div v-else class="tw-flex tw-flex-col tw-flex-nowrap tw-h-full">
    <div class="tw-flex-initial tw-flex tw-flex-col tw-bg-gray-200 tw-m-0">
      <div
        class="
          tw-flex-initial
          tw-bg-gradient-to-br
          tw-from-blue
          tw-to-turquoise
          tw-p-4
          tw-shadow-mds
          tw-flex
          sm:tw-flex-row
          tw-flex-col tw-gap-2
        "
      >
        <div
          class="
            tw-flex-initial
            tw-items-center
            tw-p-2
            tw-rounded
            tw-text-white
            tw-shadow-xl
            tw-bg-gradient-to-br
            tw-from-red-300
            tw-to-orange
            tw-text-xl
            tw-text-center
            tw-font-logo
          "
        >
          Lesson Organizer
        </div>
        <draggable
          class="tw-grid tw-grid-flow-row sm:tw-grid-flow-col tw-gap-2"
          v-model="board.goals"
          item-key="goal-id"
          animation="150"
          group="goals"
          delay="100"
          delayOnTouchOnly="true"
        >
          <template #item="{ element }">
            <goal
              :order_id="element.order_id"
              :text="element.text"
              :bgColor="element.color"
            />
          </template>
        </draggable>
        <card-button @click="console.log('foo')" class="flex-initial" />
      </div>
    </div>
    <div
      class="
        tw-relative
        tw-flex-grow
        tw-flex
        tw-overflow-x-auto
        tw-overflow-y-auto
        tw-bg-gray-200
      "
    >
      <card-button @click="onAddPhase" class="tw-sticky tw-top-0 tw-m-1" />
      <q-dialog
        full-width
        full-height
        v-model="showPhasesDialog"
        class="tw-flex-grow"
        ><q-card>
          <q-card-section class="row items-center q-pb-none">
            <div class="text-h6">Add Phase</div>
            <q-space />
            <q-btn icon="close" flat round dense v-close-popup />
          </q-card-section>
          <q-card-section>
            <q-table
              :rows="allPhases.data.value"
              :rows-per-page-options="[0]"
              @row-click="(a, row, e) => onPhaseSelect(row)"
              hide-pagination
              grid
            />
          </q-card-section>
        </q-card>
      </q-dialog>
      <draggable
        class="tw-grid tw-grid-flow-col"
        v-model="boardPhases"
        @end="onUpdatePhase"
        item-key="phase-id"
        animation="150"
        group="phases"
        delay="100"
        delayOnTouchOnly="true"
      >
        <template #item="{ element }">
          <phase
            :phaseId="element.id"
            :boardId="board.data.value.id"
            :phase="element.phase"
          />
        </template>
      </draggable>
    </div>
    <div>
      <content-list class="tw-self-start" :boardId="board.data.value.id" />
    </div>
  </div>
</template>

<script lang="ts">
import ContentList from "../components/ContentList.vue";
import Phase from "../components/Phase.vue";
import CardButton from "../components/CardButton.vue";
import Goal from "../components/Goal.vue";
import Draggable from "vuedraggable";
import { BoardPhaseDto, PhaseDto } from "@/types/phase";
import { computed, defineComponent, ref } from "@vue/runtime-core";
import { getPhases } from "@/api";
import {
  getBoard,
  getBoardPhases,
  postPhaseAssociation,
  putPhaseOrder,
} from "@/api/board";
import { useQueryClient } from "vue-query";

export default defineComponent({
  name: "Board",
  components: {
    Draggable,
    Phase,
    ContentList,
    CardButton,
    Goal,
  },
  props: {
    boardId: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const queryClient = useQueryClient();
    const board = getBoard(props.boardId);

    const allPhases = getPhases();
    const addPhase = postPhaseAssociation(props.boardId);
    const updatePhase = putPhaseOrder(props.boardId);
    const retrieveBoardPhases = getBoardPhases(props.boardId);
    const boardPhases = computed({
      get: () => {
        return retrieveBoardPhases.data.value;
      },
      set: (phases) => {
        console.log(phases);
        queryClient.setQueryData(["board_phases", props.boardId], phases);
      },
    });

    let showPhasesDialog = ref(false);
    const onAddPhase = (): void => {
      showPhasesDialog.value = true;
    };
    const onPhaseSelect = async (row: PhaseDto): Promise<void> => {
      await addPhase.mutate(
        { id: row.id },
        {
          onSuccess: () => {
            queryClient.invalidateQueries(["board_phases", props.boardId]);
          },
        }
      );

      showPhasesDialog.value = false;
    };

    const onUpdatePhase = (evt: any) => {
      if (evt.newDraggableIndex == evt.oldDraggableIndex) {
        return;
      }

      let afterId = undefined;
      if (evt.newDraggableIndex > 0) {
        const phases = queryClient.getQueryData<BoardPhaseDto[]>([
          "board_phases",
          props.boardId,
        ]);
        if (phases) {
          console.log(evt.item.id);
          afterId = phases[evt.newDraggableIndex - 1].id;
          console.log(afterId);
        }
      }

      updatePhase.mutate(
        {
          phaseId: evt.item.id,
          afterPhaseId: afterId,
        },
        {
          onSuccess: () => {
            queryClient.invalidateQueries(["board_phases", props.boardId]);
          },
        }
      );
    };

    return {
      board,
      boardPhases,
      onAddPhase,
      onPhaseSelect,
      onUpdatePhase,
      allPhases,
      showPhasesDialog,
    };
  },
});
</script>
<style scoped></style>
