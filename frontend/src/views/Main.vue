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
        <Draggable
          class="tw-grid tw-grid-flow-row sm:tw-grid-flow-col tw-gap-2"
          v-model="goals"
          item-key="goal-id"
          animation="150"
          group="goals"
          delay="100"
          delayOnTouchOnly="true"
        >
          <template #item="{ element }">
            <Goal
              :order_id="element.order_id"
              :text="element.text"
              :bgColor="element.color"
            />
          </template>
        </Draggable>
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
      <Draggable
        class="tw-grid tw-grid-flow-col"
        v-model="phases"
        item-key="phase-id"
        animation="150"
        group="phases"
        delay="100"
        delayOnTouchOnly="true"
      >
        <template #item="{ element }">
          <Phase v-model:methods="element.methods" :title="element.title" />
        </template>
      </Draggable>
    </div>
    <div>
      <list class="tw-self-start" v-model="ideas" />
    </div>
  </div>
</template>

<script lang="ts">
import List from "../components/List.vue";
import Phase from "../components/Phase.vue";
import CardButton from "../components/CardButton.vue";
import Goal from "../components/Goal.vue";
import Draggable from "vuedraggable";
import { ContentDto } from "@/types/content";
import { BoardPhaseDto, PhaseDto } from "@/types/phase";
import { computed, defineComponent, ref, Ref } from "@vue/runtime-core";
import { GoalDto } from "@/types/goal";
import { ApiActionTypes } from "@/store/modules/api/action-types";
import { useStore } from "vuex";

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

    const goals: Ref<Array<GoalDto>> = ref([
      {
        id: "goal_id_1",
        order_id: 1,
        text: "A nice goal!",
        color: "#fcba03",
      },
      {
        id: "goal_id_2",
        order_id: 2,
        text: "A loooooooooong goal!",
        color: "#bafc03",
      },
    ]);

    const ideas: Ref<Array<ContentDto>> = ref([
      { id: "idea_id_1", value: "Karte von Südost-Asien" },
      { id: "idea_id_2", value: "Nachrichtenbeitrag" },
      { id: "idea_id_6", value: "Zeitungsartikel Tsunami" },
    ]);

    const phases: Ref<Array<BoardPhaseDto>> = ref([]);

    let showPhasesDialog = ref(false);
    const onAddPhase = (): void => {
      store.dispatch(ApiActionTypes.FETCH_PHASES);
      showPhasesDialog.value = true;
    };
    const allPhases = computed(() => store.getters.allPhases);
    const onPhaseSelect = (row: PhaseDto): void => {
      console.log(row);
      phases.value.push({
        ...row,
        methods: [],
      });
      showPhasesDialog.value = false;
    };

    return {
      ideas,
      phases,
      goals,
      onAddPhase,
      onPhaseSelect,
      allPhases,
      showPhasesDialog,
    };
  },
});
</script>
<style scoped></style>
