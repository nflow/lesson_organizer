<template>
  <el-dialog title="Create Method" v-model="createMethodDialogVisible">
    <method-form v-model="createMethodModel" />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="createMethodDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="onCreateMethod">Create</el-button>
      </span>
    </template>
  </el-dialog>
  <el-dialog title="Modify Method" v-model="modifyMethodDialogVisible">
    <method-form v-model="modifyMethodModel" />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="modifyMethodDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="onModifyMethod">Update</el-button>
      </span>
    </template>
  </el-dialog>
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

      <div class="pt-2 pl-2 pr-2 bg-gray-300 rounded">
        <el-button
          type="primary"
          icon="el-icon-plus"
          @click="onOpenCreateMethodDialog"
          >Create Method</el-button
        >
      </div>
      <div class="mb-2 p-2 bg-gray-300 rounded">
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
              <el-button size="mini" @click="openModifyMethodDialog(scope.row)"
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
import MethodForm from "@/components/MethodForm.vue";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { RequestState } from "@/types/api-state";

export default defineComponent({
  name: "Manager",
  components: {
    MethodForm,
  },
  setup() {
    const store = useStore();

    onMounted(() => {
      store.dispatch(ApiActionTypes.FETCH_METHODS);
    });

    const search = ref("");

    const emptyCreateMethodModel = (): CreateMethodDto => {
      return {
        title: "",
        description: "",
        labels: [],
      };
    };
    const createMethodModel = ref(emptyCreateMethodModel());
    const createMethodDialogVisible = ref(false);
    const onOpenCreateMethodDialog = (): void => {
      createMethodModel.value = emptyCreateMethodModel();
      createMethodDialogVisible.value = true;
    };
    const onCreateMethod = async (): Promise<void> => {
      await store.dispatch(
        ApiActionTypes.CREATE_METHOD,
        createMethodModel.value
      );

      if (store.state.api.modifyMethod?.state == RequestState.SUCCESS) {
        modifyMethodDialogVisible.value = false;
      }
    };

    const modifyMethodModel = ref({} as MethodDto);
    const modifyMethodDialogVisible = ref(false);
    const openModifyMethodDialog = (method: MethodDto): void => {
      Object.assign(modifyMethodModel.value, method);
      modifyMethodDialogVisible.value = true;
    };
    const onModifyMethod = async (): Promise<void> => {
      await store.dispatch(
        ApiActionTypes.MODIFY_METHOD,
        modifyMethodModel.value
      );

      if (store.state.api.modifyMethod?.state == RequestState.SUCCESS) {
        modifyMethodDialogVisible.value = false;
      }
    };

    const methods = computed(() => store.getters.allMethods);

    const onRemoveMethod = (v: MethodDto): void => {
      store.dispatch(ApiActionTypes.DELETE_METHOD, v.id);
    };

    return {
      methods,
      search,

      createMethodModel,
      createMethodDialogVisible,
      onOpenCreateMethodDialog,
      onCreateMethod,

      modifyMethodModel,
      modifyMethodDialogVisible,
      openModifyMethodDialog,
      onModifyMethod,

      onRemoveMethod,
    };
  },
});
</script>
