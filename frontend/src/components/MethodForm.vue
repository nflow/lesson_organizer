<template>
  <q-form>
    <span>Title</span>
    <q-input outlined model="model.title" dense />
    <span>Description</span>
    <q-editor outlined v-model="model.description" min-height="5rem" />
    <span>Category</span>
    <q-input outlined type="text" v-model="model.category" dense />
    <span>Labels</span>
    <q-list>
      <q-item tag="label" v-ripple>
        <q-item-section avatar>
          <q-checkbox v-model="model.labels" val="METHOD_LABEL_SINGLE" />
        </q-item-section>
        <q-item-section>
          <q-item-label>Single Person Working</q-item-label>
          <q-item-label caption>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
            eiusmod
          </q-item-label>
        </q-item-section>
      </q-item>
      <q-item tag="label" v-ripple>
        <q-item-section avatar>
          <q-checkbox v-model="model.labels" val="METHOD_LABEL_PAIR" />
        </q-item-section>
        <q-item-section>
          <q-item-label>Partner Work</q-item-label>
        </q-item-section>
      </q-item>
      <q-item tag="label" v-ripple>
        <q-item-section avatar>
          <q-checkbox v-model="model.labels" val="METHOD_LABEL_GROUP" />
        </q-item-section>
        <q-item-section>
          <q-item-label>Group Work</q-item-label>
        </q-item-section>
      </q-item>
      <q-item tag="label" v-ripple>
        <q-item-section avatar>
          <q-checkbox v-model="model.labels" val="METHOD_LABEL_PLENUM" />
        </q-item-section>
        <q-item-section>
          <q-item-label>Plenum</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
  </q-form>
</template>

<script lang="ts">
import { defineComponent, computed, PropType } from "@vue/runtime-core";
import { CreateMethodDto, MethodDto } from "@/types/method";

export default defineComponent({
  name: "MethodForm",
  props: {
    modelValue: {
      type: Object as PropType<CreateMethodDto | MethodDto>,
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
