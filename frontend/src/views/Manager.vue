<template>
  <q-dialog v-model="createMethodDialogVisible" persistent>
    <q-card class="sm:w-full md:w-10/12">
      <q-card-section>
        <span class="font-bold text-xl">Create Method</span>
      </q-card-section>
      <q-card-section>
        <method-form v-model="createMethodModel" />
      </q-card-section>
      <q-card-section class="space-x-2">
        <q-btn color="primary" @click="onCreateMethod" label="Create" />
        <q-btn @click="createMethodDialogVisible = false" label="Cancel" />
      </q-card-section>
    </q-card>
  </q-dialog>
  <q-dialog v-model="modifyMethodDialogVisible" persistent>
    <q-card class="sm:w-full md:w-10/12">
      <q-card-section>
        <span class="font-bold text-xl">Modify Method</span>
      </q-card-section>
      <q-card-section>
        <method-form v-model="modifyMethodModel" />
      </q-card-section>
      <q-card-section class="space-x-2">
        <q-btn color="primary" @click="onModifyMethod" label="Update" />
        <q-btn @click="modifyMethodDialogVisible = false" label="Cancel" />
      </q-card-section>
    </q-card>
  </q-dialog>
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
        <q-btn
          color="primary"
          icon="note_add"
          @click="onOpenCreateMethodDialog"
          label="Create Method"
        />
      </div>
      <div class="mb-2 p-2 bg-gray-300 rounded">
        <q-table
          :loading="inProgressFetch()"
          :rows="methods"
          :columns="methodColumns"
          row-key="id"
          :rows-per-page-options="[0]"
          hide-pagination
        >
          <template v-slot:body-cell-labels="props">
            <td :props="props">
              <q-chip v-for="(v, i) in props.row.labels" :key="i">
                {{ resolveLabelName(v) }}
              </q-chip>
            </td>
          </template>
          <template v-slot:body-cell-actions="props">
            <td :props="props">
              <div class="text-right space-x-2">
                <q-btn
                  dense
                  color="primary"
                  size="sm"
                  @click="openModifyMethodDialog(props.row)"
                  label="Modify"
                />
                <q-btn
                  dense
                  size="sm"
                  @click="onRemoveMethod(props.row)"
                  label="Delete"
                />
              </div>
            </td>
          </template>
        </q-table>
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
import { CreateMethodDto, MethodDto, resolveLabelName } from "@/types/method";
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

    const emptyCreateMethodModel = (): CreateMethodDto => {
      return {
        title: "",
        description: "",
        category: "",
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

      if (store.state.api.createMethod?.state == RequestState.SUCCESS) {
        createMethodDialogVisible.value = false;
      }
    };

    const modifyMethodModel = ref({} as MethodDto);
    const modifyMethodDialogVisible = ref(false);
    const openModifyMethodDialog = (method: MethodDto): void => {
      Object.assign(modifyMethodModel.value, method);
      console.log(modifyMethodModel.value);
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

    const inProgressFetch = (): boolean => {
      if (store.state.api.allMethods?.state == RequestState.PENDING) {
        return true;
      }

      return false;
    };

    const methods = computed(() => store.getters.allMethods);
    const methodColumns = [
      { name: "title", label: "Title", field: "title", sortable: true },
      {
        name: "description",
        label: "Description",
        field: "description",
        sortable: true,
      },
      {
        name: "category",
        label: "Category",
        field: "category",
        sortable: true,
      },
      { name: "labels", label: "Labels", field: "labels", sortable: true },
      { name: "actions", label: "Actions", sortable: true },
    ];

    const onRemoveMethod = (v: MethodDto): void => {
      store.dispatch(ApiActionTypes.DELETE_METHOD, v.id);
    };

    return {
      methodColumns,
      methods,
      inProgressFetch,
      resolveLabelName,

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
