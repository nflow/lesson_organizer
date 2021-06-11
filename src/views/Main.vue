<template>
  <div class="bg-gray-200">
    <div class="bg-green-500 w-full p-4">
      <card />
    </div>
    <div class="flex flex-row bg-green-400 w-full p-4 space-x-10">
      <card title="foo" content="contetn" color="#ff0000" />
      <card-button />
      <list />
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import Card from "@/components/Card.vue";
import CardButton from "@/components/CardButton.vue";
import List from "@/components/List.vue";

@Options({
  components: {
    Card,
    CardButton,
    List,
  },
})
export default class Main extends Vue {}
</script>
<style scoped></style>
