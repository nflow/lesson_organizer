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
          item-key="id"
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
              hide-pagination
              grid
            >
              <template v-slot:item="props">
                <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4 tw-flex">
                  <div
                    @click="onPhaseSelect(props.row)"
                    class="
                      tw-flex
                      tw-flex-col
                      tw-cursor-pointer
                      tw-bg-white
                      tw-shadow-lg
                      tw-rounded-lg
                      hover:tw-shadow-2xl
                      tw-m-2 tw-p-4 tw-w-full tw-text-gray-800
                    "
                  >
                    <span>Name</span>
                    <span class="tw-text-3xl tw-font-semibold">
                      {{ props.row.title }}
                    </span>
                  </div>
                </div>
              </template>
            </q-table>
          </q-card-section>
        </q-card>
      </q-dialog>
      <draggable
        class="tw-grid tw-grid-flow-col"
        v-model="boardPhases"
        @end="onUpdatePhase"
        item-key="id"
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
            @onRemove="removePhase"
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
  deletePhase,
  getBoard,
  getBoardPhases,
  postPhaseAssociation,
  putPhaseOrder,
} from "@/api/board";
import { useQueryClient } from "vue-query";
import { useRoute, useRouter } from "vue-router";

export default defineComponent({
  name: "Board",
  components: {
    Draggable,
    Phase,
    ContentList,
    CardButton,
    Goal,
  },
  setup(props) {
    const router = useRouter();
    const route = useRoute();

    // const uuidRegex =
    //   "/^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$/i";

    const boardId: string = route.params.boardId.toString();
    // if (!boardId.match(uuidRegex)) {
    //   router.push("/");
    //   return;
    // }

    const queryClient = useQueryClient();
    const board = getBoard(boardId);

    const allPhases = getPhases();
    const addPhase = postPhaseAssociation(boardId);
    const updatePhase = putPhaseOrder(boardId);
    const retrieveBoardPhases = getBoardPhases(boardId);
    const boardPhases = computed({
      get: () => {
        return retrieveBoardPhases.data.value;
      },
      set: (phases) => {
        queryClient.setQueryData(["board_phases", boardId], phases);
      },
    });

    let showPhasesDialog = ref(false);
    const onAddPhase = (): void => {
      showPhasesDialog.value = true;
    };
    const onPhaseSelect = (row: PhaseDto) => {
      addPhase.mutate(
        { id: row.id },
        {
          onSuccess: (data) => {
            const phases: BoardPhaseDto[] | undefined =
              queryClient.getQueryData(["board_phases", boardId]);
            if (phases) {
              phases.push(data);
              queryClient.setQueryData(["board_phases", boardId], phases);
            }
            queryClient.invalidateQueries(["board_phases", boardId]);
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
          boardId,
        ]);
        if (phases) {
          afterId = phases[evt.newDraggableIndex - 1].id;
        }
      }

      updatePhase.mutate(
        {
          phaseId: evt.item.id,
          afterPhaseId: afterId,
        },
        {
          onSuccess: () => {
            queryClient.invalidateQueries(["board_phases", boardId]);
          },
        }
      );
    };

    const removePhaseMutation = deletePhase();
    const removePhase = (phaseId: string): void => {
      removePhaseMutation.mutate([boardId, phaseId], {
        onSuccess: () => {
          queryClient.invalidateQueries(["board_phases", boardId]);
        },
      });

      const newList = [...retrieveBoardPhases.data.value];
      newList.splice(
        newList.find((element) => element.id == phaseId),
        1
      );
      queryClient.setQueryData(["board_phases", boardId], newList);
    };

    return {
      removePhase,
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
