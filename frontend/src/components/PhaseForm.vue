<template>
  <q-form>
    <span>Title</span>
    <q-input outlined v-model="model.title" dense />
  </q-form>
</template>

<script lang="ts">
import { defineComponent, computed, PropType } from "@vue/runtime-core";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";

export default defineComponent({
  name: "MethodForm",
  props: {
    modelValue: {
      type: Object as PropType<CreatePhaseDto | PhaseDto>,
      required: true,
    },
  },
  setup(props, { emit }) {
    const model = computed({
      get: () => {
        return props.modelValue;
      },
      set: (value) => {
        emit("update:modelValue", value);
      },
    });

    return { model };
  },
});
</script>

<style></style>
