<template>
  <div class="flex flex-col bg-gray-200 m-0 h-full">
    <div
      class="
        flex flex-col flex-initial
        bg-gradient-to-br
        from-blue
        to-turquoise
        p-4
        space-y-4
      "
    >
      <div class="flex flex-row gap-1 flex-auto space-x-2">
        <div
          class="
            flex flex-initial flex-col
            items-center
            p-3
            h-full
            rounded
            text-white
            shadow-xl
            bg-gradient-to-br
            from-red-300
            to-orange
          "
        >
          <span class="flex-auto text-xl text-center font-logo"
            >Lesson Organizer</span
          >
        </div>
        <div class="flex-auto flex space-x-1">
          <Draggable
            class="flex flex-initial space-x-1"
            v-model="goals"
            item-key="goal-id"
            animation="150"
            group="goals"
          >
            <template #item="{ element }">
              <Goal
                :order_id="element.order_id"
                :text="element.text"
                :bgColor="element.color"
              />
            </template>
          </Draggable>
          <CardButton @click="console.log('foo')" class="flex-initial" />
        </div>
      </div>
    </div>
    <Draggable
      class="grid grid-flow-col overflow-x-auto flex-auto overflow-y-auto"
      v-model="phases"
      item-key="phase-id"
      animation="150"
      group="phases"
    >
      <template #item="{ element }">
        <Phase v-model:methods="element.methods" :title="element.title" />
      </template>
    </Draggable>
    <div class="flex-initial">
      <list class="self-start" v-model="ideas" />
    </div>
  </div>
</template>

<script lang="ts">
import List from "../components/List.vue";
import Phase from "../components/Phase.vue";
import CardButton from "../components/CardButton.vue";
import Goal from "../components/Goal.vue";
import Draggable from "vuedraggable";
import { ContentDto } from "@/types/ContentDto";
import { PhaseDto } from "@/types/PhaseDto";
import { defineComponent, ref, Ref } from "@vue/runtime-core";
import { GoalDto } from "@/types/GoalDto";

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

    const phases: Ref<Array<PhaseDto>> = ref([
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
          },
          {
            id: "method_id_2",
            title: "Blitzlicht",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [
              { id: "idea_id_5", value: "Schaubild Tsunami-Frühwarnsystem" },
            ],
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
          },
        ],
      },
      {
        id: "phase_id_9",
        title: "Übung",
        methods: [],
      },
    ]);

    return { ideas, phases, goals };
  },
});
</script>
<style scoped></style>
