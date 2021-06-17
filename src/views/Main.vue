<template>
  <div class="bg-gray-200">
    <div class="bg-green-500 w-full p-4">
      <card />

      <list v-model="ideaList" />
    </div>
    <div
      v-for="(entry, index) in ideaList"
      :key="index"
      class="flex flex-row bg-green-400 w-full p-4 space-x-10"
    >
      <card :title="entry" content="contetn" color="#ff0000" />
      <card-button />
    </div>
  </div>
</template>

<script lang="ts">
import Card from "../components/Card.vue";
import CardButton from "../components/CardButton.vue";
import List from "../components/List.vue";
import { ref } from "vue";

export default {
  name: "Main",
  components: {
    Card,
    CardButton,
    List,
  },
  setup(): Record<string, unknown> {
    const ideaList = ref(["Test1", "Test2"]);

    return { ideaList };
  },
};
</script>
<style scoped></style>
