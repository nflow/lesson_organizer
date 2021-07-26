<template>
  <q-dialog v-model="createMethodDialogVisible" persistent>
    <q-card class="sm:tw-w-full md:tw-w-10/12">
      <q-card-section>
        <span class="tw-font-bold tw-text-xl">Create Method</span>
      </q-card-section>
      <q-card-section>
        <method-form v-model="createMethodModel" />
      </q-card-section>
      <q-card-section class="tw-space-x-2">
        <q-btn color="primary" @click="onCreateMethod" label="Create" />
        <q-btn @click="createMethodDialogVisible = false" label="Cancel" />
      </q-card-section>
    </q-card>
  </q-dialog>
  <q-dialog v-model="modifyMethodDialogVisible" persistent>
    <q-card class="sm:tw-w-full md:tw-w-10/12">
      <q-card-section>
        <span class="tw-font-bold tw-text-xl">Modify Method</span>
      </q-card-section>
      <q-card-section>
        <method-form v-model="modifyMethodModel" />
      </q-card-section>
      <q-card-section class="tw-space-x-2">
        <q-btn color="primary" @click="onModifyMethod" label="Update" />
        <q-btn @click="modifyMethodDialogVisible = false" label="Cancel" />
      </q-card-section>
    </q-card>
  </q-dialog>
  <q-dialog v-model="createPhaseDialogVisible" persistent>
    <q-card class="sm:tw-w-full md:tw-w-10/12">
      <q-card-section>
        <span class="tw-font-bold tw-text-xl">Create Phase</span>
      </q-card-section>
      <q-card-section>
        <phase-form v-model="createPhaseModel" />
      </q-card-section>
      <q-card-section class="tw-space-x-2">
        <q-btn color="primary" @click="onCreatePhase" label="Create" />
        <q-btn @click="createPhaseDialogVisible = false" label="Cancel" />
      </q-card-section>
    </q-card>
  </q-dialog>
  <q-dialog v-model="modifyPhaseDialogVisible" persistent>
    <q-card class="sm:tw-w-full md:tw-w-10/12">
      <q-card-section>
        <span class="tw-font-bold tw-text-xl">Modify Phase</span>
      </q-card-section>
      <q-card-section>
        <phase-form v-model="modifyPhaseModel" />
      </q-card-section>
      <q-card-section class="tw-space-x-2">
        <q-btn color="primary" @click="onModifyPhase" label="Update" />
        <q-btn @click="modifyPhaseDialogVisible = false" label="Cancel" />
      </q-card-section>
    </q-card>
  </q-dialog>
  <div class="tw-flex tw-flex-col tw-bg-gray-200 tw-m-0 tw-h-full">
    <div
      class="
        tw-flex
        tw-flex-col
        tw-flex-initial
        tw-bg-gradient-to-br
        tw-from-blue
        tw-to-turquoise
        tw-p-4
        tw-shadow-md
        tw-space-y-4
      "
    >
      <div class="tw-flex tw-flex-row tw-gap-1 tw-flex-auto tw-space-x-2">
        <div
          class="
            tw-flex
            tw-flex-initial
            tw-flex-col
            tw-items-center
            tw-p-3
            tw-h-full
            tw-rounded
            tw-text-white
            tw-shadow-xl
            tw-bg-gradient-to-br
            tw-from-red-300
            tw-to-orange
          "
        >
          <span class="tw-flex-auto tw-text-xl tw-text-center tw-font-logo"
            >Lesson Organizer - Manager
          </span>
        </div>
      </div>
    </div>
    <div class="tw-flex tw-flex-col tw-flex-initial tw-p-5">
      <h1
        class="
          tw-text-xl
          tw-text-white
          tw-font-bold
          tw-p-2
          tw-rounded
          tw-shadow
          tw-bg-gradient-to-r
          tw-from-blue
          tw-to-turquoise
        "
      >
        Manage Methods
      </h1>

      <div class="tw-pt-2 tw-pl-2 tw-pr-2 tw-bg-gray-300 tw-rounded">
        <q-btn
          color="primary"
          icon="note_add"
          @click="onOpenCreateMethodDialog"
          label="Create Method"
        />
      </div>
      <div class="tw-mb-2 tw-p-2 tw-bg-gray-300 tw-rounded">
        <q-table
          :loading="fetchMethodsInProgress()"
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
              <div class="tw-text-right tw-space-x-2">
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

    <div class="tw-flex tw-flex-col tw-flex-initial tw-p-5">
      <h1
        class="
          tw-ext-xl
          tw-text-white
          tw-font-bold
          tw-p-2
          tw-rounded
          tw-shadow
          tw-bg-gradient-to-r
          tw-from-blue
          tw-to-turquoise
        "
      >
        Manage Phases
      </h1>

      <div class="tw-pt-2 tw-pl-2 tw-pr-2 tw-bg-gray-300 tw-rounded">
        <q-btn
          color="primary"
          icon="note_add"
          @click="onOpenCreatePhaseDialog"
          label="Create Phase"
        />
      </div>
      <div class="tw-mb-2 tw-p-2 tw-bg-gray-300 tw-rounded">
        <q-table
          :loading="fetchMethodsInProgress()"
          :rows="phases"
          :columns="phaseColumns"
          row-key="id"
          :rows-per-page-options="[0]"
          hide-pagination
        >
          <template v-slot:body-cell-actions="props">
            <td :props="props">
              <div class="tw-text-right tw-space-x-2">
                <q-btn
                  dense
                  color="primary"
                  size="sm"
                  @click="openModifyPhaseDialog(props.row)"
                  label="Modify"
                />
                <q-btn
                  dense
                  size="sm"
                  @click="onRemovePhase(props.row)"
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
import PhaseForm from "@/components/PhaseForm.vue";
import { CreateMethodDto, MethodDto, resolveLabelName } from "@/types/method";
import { RequestState } from "@/types/api-state";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";

export default defineComponent({
  name: "Manager",
  components: {
    MethodForm,
    PhaseForm,
  },
  setup() {
    const store = useStore();

    onMounted(() => {
      store.dispatch(ApiActionTypes.FETCH_METHODS);
      store.dispatch(ApiActionTypes.FETCH_PHASES);
    });

    const methods = computed(() => store.getters.allMethods);
    const methodColumns = [
      {
        name: "title",
        label: "Title",
        field: "title",
        sortable: true,
        align: "left",
      },
      {
        name: "description",
        label: "Description",
        field: "description",
        sortable: true,
        align: "left",
      },
      {
        name: "category",
        label: "Category",
        field: "category",
        sortable: true,
        align: "left",
      },
      {
        name: "labels",
        label: "Labels",
        field: "labels",
        sortable: true,
        align: "center",
      },
      { name: "actions", label: "Actions" },
    ];

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

    const onRemoveMethod = (v: MethodDto): void => {
      store.dispatch(ApiActionTypes.DELETE_METHOD, v.id);
    };

    const phases = computed(() => store.getters.allPhases);
    const phaseColumns = [
      {
        name: "title",
        label: "Title",
        field: "title",
        sortable: true,
        align: "left",
      },
      { name: "actions", label: "Actions" },
    ];

    const fetchMethodsInProgress = (): boolean => {
      if (store.state.api.allMethods?.state == RequestState.PENDING) {
        return true;
      }

      return false;
    };

    const emptyCreatePhaseModel = (): CreatePhaseDto => {
      return {
        title: "",
      };
    };
    const createPhaseModel = ref(emptyCreatePhaseModel());
    const createPhaseDialogVisible = ref(false);
    const onOpenCreatePhaseDialog = (): void => {
      createPhaseModel.value = emptyCreatePhaseModel();
      createPhaseDialogVisible.value = true;
    };
    const onCreatePhase = async (): Promise<void> => {
      await store.dispatch(ApiActionTypes.CREATE_PHASE, createPhaseModel.value);

      if (store.state.api.createPhase?.state == RequestState.SUCCESS) {
        createPhaseDialogVisible.value = false;
      }
    };

    const modifyPhaseModel = ref({} as PhaseDto);
    const modifyPhaseDialogVisible = ref(false);
    const openModifyPhaseDialog = (phase: PhaseDto): void => {
      Object.assign(modifyPhaseModel.value, phase);
      console.log(modifyPhaseModel.value);
      modifyPhaseDialogVisible.value = true;
    };
    const onModifyPhase = async (): Promise<void> => {
      await store.dispatch(ApiActionTypes.MODIFY_PHASE, modifyPhaseModel.value);

      if (store.state.api.modifyPhase?.state == RequestState.SUCCESS) {
        modifyPhaseDialogVisible.value = false;
      }
    };

    const onRemovePhase = (v: PhaseDto): void => {
      store.dispatch(ApiActionTypes.DELETE_PHASE, v.id);
    };

    const fetchPhasesInProgress = (): boolean => {
      if (store.state.api.allPhases?.state == RequestState.PENDING) {
        return true;
      }

      return false;
    };

    return {
      methodColumns,
      methods,
      fetchMethodsInProgress,
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

      phaseColumns,
      phases,
      fetchPhasesInProgress,
      createPhaseModel,
      createPhaseDialogVisible,
      onOpenCreatePhaseDialog,
      onCreatePhase,
      modifyPhaseModel,
      modifyPhaseDialogVisible,
      openModifyPhaseDialog,
      onModifyPhase,
      onRemovePhase,
    };
  },
});
</script>
