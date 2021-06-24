<template>
  <div class="flex flex-col bg-gray-200 m-0 h-full">
    <div class="w-full flex flex-row flex-initial bg-green-500 p-4">
      <div class="flex-1">
        <list v-model="ideas" />
      </div>
    </div>
    <div class="max-w-full overflow-x-auto flex-auto">
      <div class="flex flex-row flex-auto">
        <Phase
          v-for="phase in phases"
          :key="phase.id"
          v-model:methods="phase.methods"
          :title="phase.title"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import List from "../components/List.vue";
import Phase from "../components/Phase.vue";
import { IdeaDto } from "@/types/IdeaDto";
import { PhaseDto } from "@/types/PhaseDto";
import { defineComponent, ref, Ref, toRefs } from "@vue/runtime-core";

export default defineComponent({
  name: "Main",
  components: {
    Phase,
    List,
  },
  setup() {
    const ideas: Ref<Array<IdeaDto>> = ref([
      { id: "idea_id_1", value: "Test1" },
      { id: "idea_id_2", value: "Test2" },
    ]);

    const phases: Ref<Array<PhaseDto>> = ref([
      {
        id: "phase_id_1",
        title: "Einstieg",
        methods: [
          {
            id: "method_id_1",
            title: "Gruppenarbeit",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [],
          },
          {
            id: "method_id_2",
            title: "Gruppenarbeit",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [],
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
            ideas: [],
          },
        ],
      },
      {
        id: "phase_id_3",
        title: "Sicherung",
        methods: [
          {
            id: "method_id_4",
            title: "Schläge auf den Hinterkopf",
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
        id: "phase_id_4",
        title: "Übung",
        methods: [],
      },
      {
        id: "phase_id_4",
        title: "Übung",
        methods: [],
      },
      {
        id: "phase_id_4",
        title: "Übung",
        methods: [],
      },
      {
        id: "phase_id_4",
        title: "Übung",
        methods: [],
      },
      {
        id: "phase_id_4",
        title: "Übung",
        methods: [],
      },
      {
        id: "phase_id_4",
        title: "Übung",
        methods: [],
      },
    ]);

    return { ideas, phases };
  },
});
</script>
<style scoped></style>
