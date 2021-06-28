<template>
  <div class="flex flex-col bg-gray-200 m-0 h-full">
    <div class="flex flex-col flex-initial bg-green-500 p-4 space-y-4">
      <div class="flex flex-row gap-1 flex-auto">
        <div
          class="
            flex-auto
            text-center
            bg-green-700
            p-3
            text-white
            self-center
            rounded
          "
        >
          <h1 class="text-xl font-bold">Lesson Organizer - Demo</h1>
          <span
            >Data is not persited! Start by draging methods and ideas and have
            fun!</span
          >
        </div>
        <div
          class="
            flex-initial flex flex-row
            text-center
            bg-red-500
            p-4
            text-white
            rounded
          "
        >
          <div class="flex-initial text-4xl align-middle pr-4">1</div>
          <div class="flex items-center">My great goal!</div>
        </div>

        <div
          class="
            flex-initial flex flex-row
            text-center
            bg-yellow-500
            p-4
            text-white
            rounded
          "
        >
          <div class="flex-initial text-4xl align-middle pr-4">2</div>
          <div class="flex items-center">My great goal!</div>
        </div>
      </div>
      <div class="flex-1">
        <list class="self-start" v-model="ideas" />
      </div>
    </div>
    <Draggable
      class="
        grid grid-flow-col
        max-w-full
        overflow-x-auto
        flex-auto
        h-0
        overflow-y-auto
      "
      v-model="phases"
      item-key="phase-id"
      animation="150"
      group="phases"
    >
      <template #item="{ element }">
        <Phase v-model:methods="element.methods" :title="element.title" />
      </template>
    </Draggable>
  </div>
</template>

<script lang="ts">
import List from "../components/List.vue";
import Phase from "../components/Phase.vue";
import Draggable from "vuedraggable";
import { IdeaDto } from "@/types/IdeaDto";
import { PhaseDto } from "@/types/PhaseDto";
import { defineComponent, ref, Ref } from "@vue/runtime-core";

export default defineComponent({
  name: "Main",
  components: {
    Draggable,
    Phase,
    List,
  },
  setup() {
    const ideas: Ref<Array<IdeaDto>> = ref([
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

    return { ideas, phases };
  },
});
</script>
<style scoped></style>
