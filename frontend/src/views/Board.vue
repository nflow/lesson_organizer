<template>
  <div
    class="tw-text-5xl"
    v-if="createBoard.isLoading.value || board.isLoading.value"
  >
    Loading board ...
  </div>
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
        v-if="board.isSuccess"
        class="tw-grid tw-grid-flow-col"
        v-model="boardPhases"
        @update="onUpdatePhase"
        item-key="phase-id"
        animation="150"
        group="phases"
        delay="100"
        delayOnTouchOnly="true"
      >
        <template #item="{ element }">
          <phase :phaseId="element.id" :phase="element.phase" />
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
import { PhaseDto } from "@/types/phase";
import { computed, defineComponent, ref } from "@vue/runtime-core";
import { BoardDto, CreateBoardDto } from "@/types/board";
import { getPhases } from "@/api";
import {
  getBoard,
  postBoard,
  postPhaseAssociation,
  putPhaseOrder,
} from "@/api/board";
import { useRoute, useRouter } from "vue-router";
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
  setup() {
    const router = useRouter();
    const route = useRoute();
    const boardId = ref(route.params.boardId);

    const queryClient = useQueryClient();
    const createBoard = postBoard();

    if (!route.params.boardId) {
      createBoard.mutate(
        {
          name: "New Board",
          goals: [],
          contents: [],
          phases: [],
        } as CreateBoardDto,
        {
          onSuccess: (data: BoardDto) => {
            queryClient.setQueryData(["board"], data);
            boardId.value = data.id;
            router.replace({
              name: "Board",
              params: {
                boardId: data ? data.id : "",
              },
            });
          },
        }
      );
    }

    const boardPhases = computed({
      get: () => {
        return board.data.value.phases;
      },
      set: (newPhaseOrder) => {
        const currentBoard: BoardDto = board.data.value;
        // TODO: Refactor to use separate lists
        queryClient.setQueryData("board", {
          id: currentBoard.id,
          name: currentBoard.name,
          goals: currentBoard.goals,
          contents: currentBoard.contents,
          phases: newPhaseOrder,
        } as BoardDto);
        return;
      },
    });
    const board = getBoard(boardId);
    const allPhases = getPhases();
    const addPhase = postPhaseAssociation(boardId);
    const updatePhase = putPhaseOrder(boardId);

    let showPhasesDialog = ref(false);
    const onAddPhase = (): void => {
      showPhasesDialog.value = true;
    };
    const onPhaseSelect = async (row: PhaseDto): Promise<void> => {
      await addPhase.mutateAsync({ id: row.id });

      if (addPhase.isSuccess) {
        queryClient.invalidateQueries("board");
        showPhasesDialog.value = false;
      }
    };

    const onUpdatePhase = (evt: any) => {
      if (evt.newDraggableIndex == evt.oldDraggableIndex || !board.data.value) {
        {
          {
            boardPhases;
          }
        }
        return;
      }

      const currentBoard = queryClient.getQueryData<BoardDto>("board");
      if (!currentBoard) {
        return;
      }

      const elementId = currentBoard.phases[evt.newDraggableIndex].id;
      let afterId = undefined;
      if (evt.newDraggableIndex > 0) {
        afterId = board.data.value.phases[evt.newDraggableIndex - 1].id;
      }

      updatePhase.mutate(
        {
          phaseId: elementId,
          afterPhaseId: afterId,
        },
        {
          onSuccess: () => {
            // TODO: Could return the new list of phases .. or maybe not.
            queryClient.invalidateQueries("board");
          },
        }
      );
    };

    return {
      createBoard,
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
