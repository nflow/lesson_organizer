<template>
  <router-view />
  <VueQueryDevTools v-if="isDevelopment" :initialIsOpen="true" />
</template>
<script lang="ts">
import { computed, defineComponent } from "vue";
import { useQueryProvider } from "vue-query";
import { VueQueryDevTools } from "vue-query/devtools";

export default defineComponent({
  name: "App",
  components: {
    VueQueryDevTools,
  },
  setup() {
    const isDevelopment = computed(() => {
      return process.env.NODE_ENV === "development";
    });
    useQueryProvider();
    return { isDevelopment };
  },
});
</script>
<style></style>
