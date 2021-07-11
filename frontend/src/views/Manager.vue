<template>
  <div class="flex flex-col bg-gray-200 m-0 h-full">
    <div
      class="
        flex flex-col flex-initial
        bg-gradient-to-br
        from-blue
        to-turquoise
        p-4
        shadow-md
        space-y-4
      "
    >
      <div class="flex flex-row gap-1 flex-auto space-x-2">
        <div
          class="
            flex flex-initial flex-col
            items-center
            p-3
            h-full
            rounded
            text-white
            shadow-xl
            bg-gradient-to-br
            from-red-300
            to-orange
          "
        >
          <span class="flex-auto text-xl text-center font-logo"
            >Lesson Organizer - Manager
          </span>
        </div>
      </div>
    </div>
    <div class="flex flex-col flex-auto p-5">
      <h1
        class="
          text-xl text-white
          font-bold
          p-2
          rounded
          shadow
          bg-gradient-to-r
          from-blue
          to-turquoise
        "
      >
        Manage Methods
      </h1>
      <div class="mt-2 mb-2 p-2 bg-gray-300 rounded">
        <el-form ref="newMethodForm" :model="newMethod" label-width="100px">
          <el-form-item label="Title">
            <el-input v-model="newMethod.title" />
          </el-form-item>
          <el-form-item label="Description">
            <el-input type="textarea" v-model="newMethod.description" />
          </el-form-item>
          <el-form-item label="Labels">
            <el-checkbox-group v-model="newMethod.labels">
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
          <el-form-item size="large">
            <el-button type="primary" @click="onCreateMethod">Create</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div>
        <el-table
          :data="
            methods.filter(
              (data) =>
                !search ||
                data.title.toLowerCase().includes(search.toLowerCase())
            )
          "
        >
          <el-table-column prop="title" label="Name"> </el-table-column>
          <el-table-column prop="description" label="Description" />
          <el-table-column prop="labels" label="Labels" />
          <el-table-column align="right">
            <template #header>
              <el-input
                v-model="search"
                size="mini"
                placeholder="Type to search"
              />
            </template>
            <template #default="scope">
              <el-button
                size="mini"
                @click="onEditMethod(scope.$index, scope.row)"
                >Edit</el-button
              >
              <el-button
                size="mini"
                type="danger"
                @click="onRemoveMethod(scope.row)"
                >Delete</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { computed, defineComponent, onMounted } from "@vue/runtime-core";
import { ref } from "@vue/reactivity";
import { useStore } from "@/store";
import { ApiActionTypes } from "@/store/modules/api/action-types";
import { CreateMethodDto } from "@/types/CreateMethodDto";
import { MethodDto } from "@/types/MethodDto";

export default defineComponent({
  name: "Manager",
  setup() {
    const store = useStore();
    const search = ref("");
    const newMethod = ref({
      title: "",
      description: "",
      labels: [],
    } as CreateMethodDto);

    onMounted(() => {
      store.dispatch(ApiActionTypes.FETCH_METHODS);
    });

    const methods = computed(() => store.state.api.methods);

    const onCreateMethod = (): void => {
      store.dispatch(ApiActionTypes.CREATE_METHOD, newMethod.value);
    };

    const onRemoveMethod = (v: MethodDto): void => {
      store.dispatch(ApiActionTypes.DELETE_METHOD, v.id);
    };

    const onEditMethod = (index: number, row: number): void => {
      console.log(index);
      console.log(row);
    };

    return {
      newMethod,
      methods,
      search,
      onCreateMethod,
      onRemoveMethod,
      onEditMethod,
    };
  },
});
</script>
