<template>
  <div class="bg-gray-200">
    <div class="flex flex-row bg-green-500 w-full p-4">
      <div class="flex-1">
        <list v-model="ideaList" />
      </div>
      <div class="flex-"></div>
    </div>
    <div class="flex flex-row w-full">
      <div class="flex flex-col hover:bg-gray-300">
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
          Phase 1
        </div>
        <div class="flex-1 flex flex-row p-4 space-x-10">
          <card
            title="Methode 1"
            content="Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua."
          >
            <list v-model="subIdeaList" />
          </card>
        </div>
      </div>
      <div class="flex flex-col hover:bg-gray-300">
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
          Phase 2
        </div>
        <div class="flex-1 flex flex-row p-4 space-x-10">
          <card
            title="Methode2"
            content="At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."
          >
            <list v-model="subIdeaList2" />
          </card>
        </div>
      </div>

      <card-button />
    </div>
  </div>
</template>

<script lang="ts">
import Card from "../components/Card.vue";
import CardButton from "../components/CardButton.vue";
import List from "../components/List.vue";
import { Ref, ref } from "vue";
import { ListEntry } from "@/types/List";

export default {
  name: "Main",
  components: {
    Card,
    CardButton,
    List,
  },
  setup(): Record<string, unknown> {
    const ideaList = ref({
      list: [
        { id: "id1", value: "Test1" },
        { id: "id2", value: "Test2" },
      ],
    });
    const subIdeaList = ref({
      list: [
        { id: "id3", value: "Test3" },
        { id: "id4", value: "Test4" },
      ],
    });

    const subIdeaList2 = ref({
      list: [
        { id: "id3", value: "Test3" },
        { id: "id4", value: "Test4" },
      ],
    });
    return { ideaList, subIdeaList, subIdeaList2 };
  },
};
</script>
<style scoped></style>
