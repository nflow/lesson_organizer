<template>
  <div class="bg-gray-200">
    <div class="flex flex-row bg-green-500 w-full p-4">
      <div class="flex-1">
        <list v-model="ideas" />
      </div>
    </div>
    <div class="flex flex-row w-full min-h-screen">
      <div
        v-for="phase in phases"
        :key="phase.id"
        class="flex-1 flex flex-col hover:bg-gray-300"
      >
        <div
          class="
            flex-initial
            font-extrabold
            hover:bg-green-500
            bg-green-400
            p-4
            text-center
          "
        >
          {{ phase.title }}
        </div>
        <div
          class="flex-1 flex flex-col p-4 space-y-10 items-start self-center"
        >
          <method
            v-for="method in phase.methods"
            :key="method.id"
            :title="method.title"
            :description="method.description"
            :ideas="method.ideas"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Method from "../components/Method.vue";
import List from "../components/List.vue";
import { IdeaDto } from "@/types/IdeaDto";
import { PhaseDto } from "@/types/PhaseDto";
import { defineComponent, ref, Ref, toRefs } from "@vue/runtime-core";

export default defineComponent({
  name: "Main",
  components: {
    Method,
    List,
  },
  setup(): Record<string, unknown> {
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
            id: "method_id_1",
            title: "Gruppenarbeit",
            description:
              "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
            ideas: [],
          },
        ],
      },
    ]);

    return { ideas, phases };
  },
});
</script>
<style scoped></style>
