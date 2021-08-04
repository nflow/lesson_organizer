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
          :loading="methods.isLoading"
          :rows="methods.data.value"
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
          :loading="phases.isLoading"
          :rows="phases.data.value"
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
import { defineComponent } from "@vue/runtime-core";
import { ref } from "@vue/reactivity";
import MethodForm from "@/components/MethodForm.vue";
import PhaseForm from "@/components/PhaseForm.vue";
import { CreateMethodDto, MethodDto, resolveLabelName } from "@/types/method";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";
import { useQuery, useMutation, useQueryClient } from "vue-query";
import {
  getMethods,
  getPhases,
  postMethod,
  putMethod,
  deleteMethod,
  postPhase,
  putPhase,
  deletePhase,
} from "@/api";

export default defineComponent({
  name: "Manager",
  components: {
    MethodForm,
    PhaseForm,
  },
  setup() {
    const queryClient = useQueryClient();

    const methods = useQuery("methods", getMethods);
    const createMethod = useMutation(postMethod, {
      onSuccess: () => {
        queryClient.invalidateQueries("methods");
      },
    });
    const updateMethod = useMutation(putMethod, {
      onSuccess: () => {
        queryClient.invalidateQueries("methods");
      },
    });
    const removeMethod = useMutation(deleteMethod, {
      onSuccess: () => {
        queryClient.invalidateQueries("methods");
      },
    });

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
      createMethod.mutate(createMethodModel.value);

      if (createMethod.isSuccess) {
        createMethodDialogVisible.value = false;
      }
    };

    const modifyMethodModel = ref({} as MethodDto);
    const modifyMethodDialogVisible = ref(false);
    const openModifyMethodDialog = (method: MethodDto): void => {
      Object.assign(modifyMethodModel.value, method);
      modifyMethodDialogVisible.value = true;
    };
    const onModifyMethod = async (): Promise<void> => {
      updateMethod.mutate(modifyMethodModel.value);

      if (updateMethod.isSuccess) {
        modifyMethodDialogVisible.value = false;
      }
    };

    const onRemoveMethod = (v: MethodDto): void => {
      removeMethod.mutate(v.id);
    };

    const phases = useQuery("phases", getPhases);
    const createPhase = useMutation(postPhase, {
      onSuccess: () => {
        queryClient.invalidateQueries("phases");
      },
    });
    const updatePhase = useMutation(putPhase, {
      onSuccess: () => {
        queryClient.invalidateQueries("phases");
      },
    });
    const removePhase = useMutation(deletePhase, {
      onSuccess: () => {
        queryClient.invalidateQueries("phases");
      },
    });

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
      createPhase.mutate(createPhaseModel.value);

      if (createPhase.isSuccess) {
        createPhaseDialogVisible.value = false;
      }
    };

    const modifyPhaseModel = ref({} as PhaseDto);
    const modifyPhaseDialogVisible = ref(false);
    const openModifyPhaseDialog = (phase: PhaseDto): void => {
      Object.assign(modifyPhaseModel.value, phase);
      modifyPhaseDialogVisible.value = true;
    };
    const onModifyPhase = async (): Promise<void> => {
      updatePhase.mutate(modifyPhaseModel.value);

      if (updatePhase.isSuccess) {
        modifyPhaseDialogVisible.value = false;
      }
    };

    const onRemovePhase = (v: PhaseDto): void => {
      removePhase.mutate(v.id);
    };

    return {
      methodColumns,
      methods,
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
