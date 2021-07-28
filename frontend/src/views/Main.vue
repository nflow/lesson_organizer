<template>
  <div class="tw-flex tw-flex-col tw-flex-nowrap tw-h-full">
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
              :rows="allPhases"
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
        v-model="board.phases"
        item-key="phase-id"
        animation="150"
        group="phases"
        delay="100"
        delayOnTouchOnly="true"
      >
        <template #item="{ element }">
          <phase v-model:methods="element.methods" :title="element.title" />
        </template>
      </draggable>
    </div>
    <div>
      <list class="tw-self-start" v-model="board.contents" />
    </div>
  </div>
</template>

<script lang="ts">
import List from "../components/List.vue";
import Phase from "../components/Phase.vue";
import CardButton from "../components/CardButton.vue";
import Goal from "../components/Goal.vue";
import Draggable from "vuedraggable";
import { PhaseDto } from "@/types/phase";
import { computed, defineComponent, onMounted, ref } from "@vue/runtime-core";
import { ApiActionTypes } from "@/store/modules/api/action-types";
import { useStore } from "vuex";
import { BoardDto } from "@/types/board";
import { BoardMutationTypes } from "@/store/modules/board/mutation-types";

export default defineComponent({
  name: "Main",
  components: {
    Draggable,
    Phase,
    List,
    CardButton,
    Goal,
  },
  setup() {
    const store = useStore();

    onMounted(() => {
      if (store.state.board.currentBoard == undefined) {
        store.commit(BoardMutationTypes.SET_CURRENT_BOARD, {
          id: "id",
          name: "name",
          goals: [],
          contents: [],
          phases: [],
        });
      }
    });
    const board = computed({
      get: (): BoardDto => {
        console.log("read ..");
        return store.state.board.currentBoard;
      },
      set: (value: BoardDto) => {
        console.log("write back ..");
        store.commit(BoardMutationTypes.SET_CURRENT_BOARD, value);
      },
    });

    let showPhasesDialog = ref(false);
    const onAddPhase = (): void => {
      store.dispatch(ApiActionTypes.FETCH_PHASES);
      showPhasesDialog.value = true;
    };
    const allPhases = computed(() => store.getters.allPhases);
    const onPhaseSelect = (row: PhaseDto): void => {
      const board = store.state.board.currentBoard;
      board.phases.push({
        ...row,
        methods: [],
      });
      store.commit(BoardMutationTypes.SET_CURRENT_BOARD, board);
      showPhasesDialog.value = false;
    };

    return {
      board,
      onAddPhase,
      onPhaseSelect,
      allPhases,
      showPhasesDialog,
    };
  },
});
</script>
<style scoped></style>
