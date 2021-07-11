<template>
  <el-form ref="newMethodForm" :model="model" label-width="100px">
    <el-form-item label="Title">
      <el-input v-model="model.title" />
    </el-form-item>
    <el-form-item label="Description">
      <el-input type="textarea" v-model="model.description" />
    </el-form-item>
    <el-form-item label="Labels">
      <el-checkbox-group v-model="model.labels">
        <el-checkbox-button label="METHOD_LABEL_SINGLE" name="type"
          >Single Person Working</el-checkbox-button
        >
        <el-checkbox-button label="METHOD_LABEL_PAIR" name="type"
          >Partner Work</el-checkbox-button
        >
        <el-checkbox-button label="METHOD_LABEL_GROUP" name="type"
          >Group Work</el-checkbox-button
        >
        <el-checkbox-button label="METHOD_LABEL_PLENUM" name="type"
          >Plenum</el-checkbox-button
        >
      </el-checkbox-group>
    </el-form-item>
  </el-form>
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
