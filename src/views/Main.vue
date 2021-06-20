<template>
  <div class="bg-gray-200">
    <div class="bg-green-500 w-full p-4">
      <card />

      <list v-model="ideaList" />
    </div>
    <div class="flex flex-row bg-green-400 w-full p-4 space-x-10">
      <card title="Editable Card" content="Test">
        <list v-model="subIdeaList" />
      </card>
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

    return { ideaList, subIdeaList };
  },
};
</script>
<style scoped></style>
