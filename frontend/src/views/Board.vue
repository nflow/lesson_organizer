<template>
  <div class="tw-text-5xl" v-if="board.isLoading.value">Loading board ...</div>
  <div v-else class="tw-flex tw-flex-col tw-flex-nowrap tw-h-full">
    <div class="tw-flex-initial tw-flex tw-flex-col tw-bg-gray-200 tw-m-0">
      <div
        class="
          isLoading
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
          v-model="boardGoals"
          @end="onMoveGoal"
          item-key="id"
          animation="150"
          group="goals"
          delay="100"
          delayOnTouchOnly="true"
        >
          <template #item="{ element }">
            <goal
              :id="element.id"
              :order_id="element.order_id"
              :text="element.text"
              :bgColor="element.color"
              @onModify="modifyGoalDialog.onOpen(element)"
              @onDelete="onRemoveGoal(element)"
            />
          </template>
        </draggable>
        <card-button @click="createGoalDialog.onOpen" class="flex-initial" />
        <q-dialog v-model="createGoalDialog.visible.value" persistent>
          <q-card class="sm:tw-w-full md:tw-w-10/12">
            <form @submit.prevent.stop="createGoalDialog.onSubmit">
              <q-card-section>
                <span class="tw-font-bold tw-text-xl">Create Goal</span>
              </q-card-section>

              <q-card-section>
                <q-input
                  :ref="createGoalDialog.components.text"
                  v-model="createGoalDialog.model.value.text"
                  square
                  outlined
                  lazy-rules
                  autofocus
                  :rules="createGoalDialog.rules.text"
                  label="Text"
                />
              </q-card-section>
              <q-card-section class="tw-space-x-2">
                <q-btn
                  color="primary"
                  type="submit"
                  label="Create"
                  :loading="createGoal.isLoading.value"
                />
                <q-btn @click="createGoalDialog.onCancle" label="Cancel" />
              </q-card-section>
            </form>
          </q-card>
        </q-dialog>
        <q-dialog v-model="modifyGoalDialog.visible.value" persistent>
          <q-card class="sm:tw-w-full md:tw-w-10/12">
            <form @submit.prevent.stop="modifyGoalDialog.onSubmit">
              <q-card-section>
                <span class="tw-font-bold tw-text-xl">Modify Goal</span>
              </q-card-section>

              <q-card-section>
                <q-input
                  :ref="modifyGoalDialog.components.text"
                  v-model="modifyGoalDialog.model.value.text"
                  square
                  outlined
                  lazy-rules
                  autofocus
                  :rules="modifyGoalDialog.rules.text"
                  label="Text"
                />
              </q-card-section>
              <q-card-section class="tw-space-x-2">
                <q-btn
                  color="primary"
                  type="submit"
                  label="Create"
                  :loading="modifyGoal.isLoading.value"
                />
                <q-btn @click="modifyGoalDialog.onCancle" label="Cancel" />
              </q-card-section>
            </form>
          </q-card>
        </q-dialog>
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
              :loading="addPhase.isLoading.value"
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
        @end="onMovePhase"
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
            @onRemove="onRemovePhase"
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
  deleteGoal,
  deletePhase,
  getBoard,
  getBoardPhases,
  getGoals,
  postGoal,
  postPhaseAssociation,
  putModifyGoal,
  putMoveGoal,
  putPhaseOrder,
} from "@/api/board";
import { useQueryClient } from "vue-query";
import { useRoute, useRouter } from "vue-router";
import { CreateGoalDto, GoalDto } from "@/types/goal";
import { GenericDialog } from "@/types/dialog";

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
      if (addPhase.isLoading.value) {
        return;
      }

      addPhase.mutate(
        { id: row.id },
        {
          onSuccess: (data) => {
            const phases: BoardPhaseDto[] | undefined =
              queryClient.getQueryData(["board_phases", boardId]);
            if (phases) {
              queryClient.setQueryData(
                ["board_phases", boardId],
                [...phases, data]
              );
            }
            queryClient.invalidateQueries(["board_phases", boardId]);
            showPhasesDialog.value = false;
          },
        }
      );
    };

    const onMovePhase = (evt: any) => {
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

    const removePhase = deletePhase();
    const onRemovePhase = (phaseId: string): void => {
      removePhase.mutate([boardId, phaseId], {
        onSuccess: () => {
          queryClient.invalidateQueries(["board_phases", boardId]);
        },
      });

      const newList = [...retrieveBoardPhases.data.value];
      newList.splice(
        newList.findIndex((element) => element.id == phaseId),
        1
      );
      queryClient.setQueryData(["board_phases", boardId], newList);
    };

    const retrieveGoals = getGoals(boardId);
    const boardGoals = computed({
      get: () => {
        return retrieveGoals.data.value;
      },
      set: (goals) => {
        queryClient.setQueryData(["board_goals", boardId], goals);
      },
    });

    const removeGoal = deleteGoal(boardId);
    const onRemoveGoal = (goal: GoalDto) => {
      removeGoal.mutate(goal.id, {
        onSuccess: () => {
          queryClient.invalidateQueries(["board_goals", boardId]);
        },
      });

      const newList = [...boardGoals.value];
      newList.splice(
        newList.findIndex((element) => element.id == goal.id),
        1
      );
      queryClient.setQueryData(["board_goals", boardId], newList);
    };

    const modifyGoal = putModifyGoal(boardId);
    const modifyGoalDialog: GenericDialog<GoalDto> = {
      visible: ref(false),
      model: ref({ id: "", text: "", rank: 0 }),
      components: {
        text: ref<any>({}),
      },
      rules: {
        text: [
          (val: string | undefined) =>
            (val && val.length > 0) || "Please type something",
        ],
      },
      onOpen: (element: GoalDto) => {
        modifyGoalDialog.model.value = element;
        modifyGoalDialog.visible.value = true;
      },
      onSubmit: () => {
        if (!modifyGoalDialog.components) {
          return;
        }

        modifyGoalDialog.components.text.value.validate();

        if (!modifyGoalDialog.components.text.value.hasError) {
          modifyGoal.mutate(
            {
              goalId: modifyGoalDialog.model.value.id,
              payload: { text: modifyGoalDialog.model.value.text },
            },
            {
              onSuccess: () => {
                queryClient.invalidateQueries(["board_goals", boardId]);
                modifyGoalDialog.visible.value = false;
              },
            }
          );
        }
      },
      onCancle: () => {
        modifyGoalDialog.visible.value = false;
      },
    };

    const moveGoal = putMoveGoal(boardId);
    const onMoveGoal = (evt: any) => {
      if (evt.newDraggableIndex == evt.oldDraggableIndex) {
        return;
      }

      let afterId = undefined;
      if (evt.newDraggableIndex > 0) {
        const goals = queryClient.getQueryData<BoardPhaseDto[]>([
          "board_goals",
          boardId,
        ]);
        if (goals) {
          afterId = goals[evt.newDraggableIndex - 1].id;
        }
      }

      moveGoal.mutate(
        {
          goalId: evt.item.id,
          afterGoalId: afterId,
        },
        {
          onSuccess: () => {
            queryClient.invalidateQueries(["board_goals", boardId]);
          },
        }
      );
    };

    const createGoal = postGoal(boardId);
    const createGoalDialog: GenericDialog<CreateGoalDto> = {
      visible: ref(false),
      model: ref({ text: "" }),
      components: {
        text: ref<any>({}),
      },
      rules: {
        text: [
          (val: string | undefined) =>
            (val && val.length > 0) || "Please type something",
        ],
      },
      onOpen: () => {
        createGoalDialog.model.value = { text: "" };
        createGoalDialog.visible.value = true;
      },
      onSubmit: () => {
        if (!createGoalDialog.components) {
          return;
        }

        createGoalDialog.components.text.value.validate();

        if (!createGoalDialog.components.text.value.hasError) {
          createGoal.mutate(createGoalDialog.model.value, {
            onSuccess: (data: GoalDto) => {
              queryClient.setQueryData(
                ["board_goals", boardId],
                [...boardGoals.value, data]
              );
              createGoalDialog.visible.value = false;
            },
          });
        }
      },
      onCancle: () => {
        createGoalDialog.model.value = { text: "" };
        createGoalDialog.visible.value = false;
      },
    };

    return {
      onMoveGoal,
      onRemoveGoal,
      boardGoals,
      createGoal,
      createGoalDialog,
      modifyGoal,
      modifyGoalDialog,

      onRemovePhase,
      board,
      boardPhases,
      addPhase,
      onAddPhase,
      onPhaseSelect,
      onMovePhase,
      allPhases,
      showPhasesDialog,
    };
  },
});
</script>
<style scoped></style>
