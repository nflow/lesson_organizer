<template>
  <router-view />
  <VueQueryDevTools :initialIsOpen="true" />
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { useQueryProvider } from "vue-query";
import { VueQueryDevTools } from "vue-query/devtools";

export default defineComponent({
  name: "App",
  components: {
    VueQueryDevTools,
  },
  setup() {
    useQueryProvider();
  },
});
</script>
<style></style>
