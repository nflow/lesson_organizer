<template>
  <div class="flex flex-col bg-gray-200 m-0">
    <div
      class="
        bg-gradient-to-br
        from-blue
        to-turquoise
        p-4
        shadow-mds
        flex-none flex
        sm:flex-row
        flex-col
        gap-2
      "
    >
      <div
        class="
          flex-initial
          items-center
          p-2
          rounded
          text-white
          shadow-xl
          bg-gradient-to-br
          from-red-300
          to-orange
          text-xl text-center
          font-logo
        "
      >
        Lesson Organizer
      </div>
      <Draggable
        class="grid grid-flow-row sm:grid-flow-col gap-2"
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
  <div class="grid grid-flow-col overflow-x-auto overflow-y-auto bg-gray-200">
    <card-button @click="onAddPhase" class="auto-height m-1" />
    <q-dialog full-width full-height v-model="showPhasesDialog"
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
            hide-pagination
            grid
          />
        </q-card-section>
      </q-card>
    </q-dialog>
    <Draggable
      class="grid grid-flow-col"
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
  <div class="flex-initial">
    <list class="self-start" v-model="ideas" />
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

    const phases: Ref<Array<BoardPhaseDto>> = ref([
      {
        id: "phase_id_1",
        title: "Einstieg",
        methods: [
          {
            id: "method_id_1",
            title: "Mind-Map",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [],
            labels: [],
          },
          {
            id: "method_id_2",
            title: "Blitzlicht",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [
              { id: "idea_id_5", value: "Schaubild Tsunami-Frühwarnsystem" },
            ],
            labels: [],
          },
        ],
      },
      {
        id: "phase_id_2",
        title: "Erarbeitung",
        methods: [
          {
            id: "method_id_3",
            title: "Internetrecherche",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [
              { id: "idea_id_3", value: "Fotos der Zerstörung" },
              { id: "idea_id_4", value: "Augenzeugenbericht" },
            ],
            labels: [],
          },
        ],
      },
      {
        id: "phase_id_3",
        title: "Sicherung",
        methods: [
          {
            id: "method_id_4",
            title: "Soziometrische Abfrage",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [],
            labels: [],
          },
        ],
      },
      {
        id: "phase_id_4",
        title: "Überleitung",
        methods: [],
      },
      {
        id: "phase_id_5",
        title: "Übung",
        methods: [
          {
            id: "method_id_5",
            title: "Soziometrische Abfrage",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [],
            labels: [],
          },
        ],
      },
      {
        id: "phase_id_6",
        title: "Erarbeitung",
        methods: [],
      },
      {
        id: "phase_id_7",
        title: "Sicherung",
        methods: [],
      },
      {
        id: "phase_id_8",
        title: "Überleitung",
        methods: [
          {
            id: "method_id_6",
            title: "Fishbowl-Diskussion",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [],
            labels: [],
          },
        ],
      },
      {
        id: "phase_id_9",
        title: "Übung",
        methods: [],
      },
    ]);

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
