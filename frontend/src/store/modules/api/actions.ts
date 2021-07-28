import axios, { AxiosError } from "axios";
import { ActionContext, ActionTree } from "vuex";
import { ApiActionTypes } from "./action-types";
import { ApiMutationTypes } from "./mutation-types";
import config from "@/config";
import { ApiMutationsTypes } from "./mutations";
import { RootState } from "@/store";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { RequestState } from "@/types/api-state";
import { ApiState } from "./state";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";

export type AugmentedActionContext = {
  commit<K extends keyof ApiMutationsTypes>(
    key: K,
    payload: Parameters<ApiMutationsTypes[K]>[1]
  ): ReturnType<ApiMutationsTypes[K]>;
} & Omit<ActionContext<ApiState, RootState>, "commit">;

export interface ApiActionsTypes {
  [ApiActionTypes.CREATE_METHOD](
    { commit }: AugmentedActionContext,
    payload: CreateMethodDto
  ): Promise<void>;
  [ApiActionTypes.MODIFY_METHOD](
    { commit }: AugmentedActionContext,
    payload: MethodDto
  ): Promise<void>;
  [ApiActionTypes.DELETE_METHOD](
    { commit }: AugmentedActionContext,
    uuid: string
  ): Promise<void>;
  [ApiActionTypes.FETCH_METHODS]({ commit }: AugmentedActionContext): void;

  [ApiActionTypes.CREATE_PHASE](
    { commit }: AugmentedActionContext,
    payload: CreatePhaseDto
  ): Promise<void>;
  [ApiActionTypes.MODIFY_PHASE](
    { commit }: AugmentedActionContext,
    payload: PhaseDto
  ): Promise<void>;
  [ApiActionTypes.DELETE_PHASE](
    { commit }: AugmentedActionContext,
    uuid: string
  ): Promise<void>;
  [ApiActionTypes.FETCH_PHASES]({ commit }: AugmentedActionContext): void;
}

export const actions: ActionTree<ApiState, RootState> & ApiActionsTypes = {
  async [ApiActionTypes.CREATE_METHOD](
    { commit, dispatch },
    payload: CreateMethodDto
  ) {
    commit(ApiMutationTypes.SET_CREATE_METHOD, {
      state: RequestState.PENDING,
    });
    await axios
      .post(config.CONFIG_API_URL + "v1/methods", payload)
      .then((response) => {
        commit(ApiMutationTypes.SET_CREATE_METHOD, {
          state: RequestState.SUCCESS,
          data: response.data,
        });

        dispatch(ApiActionTypes.FETCH_METHODS);
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_CREATE_METHOD, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },
  async [ApiActionTypes.MODIFY_METHOD](
    { commit, dispatch },
    payload: MethodDto
  ) {
    commit(ApiMutationTypes.SET_MODIFY_METHOD, {
      state: RequestState.PENDING,
    });
    await axios
      .put(config.CONFIG_API_URL + "v1/methods/" + payload.id, payload)
      .then((response) => {
        commit(ApiMutationTypes.SET_MODIFY_METHOD, {
          state: RequestState.SUCCESS,
          data: response.data,
        });

        dispatch(ApiActionTypes.FETCH_METHODS);
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_MODIFY_METHOD, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },
  async [ApiActionTypes.DELETE_METHOD]({ commit, dispatch }, uuid: string) {
    commit(ApiMutationTypes.SET_DELETE_METHOD, {
      state: RequestState.PENDING,
    });
    await axios
      .delete(config.CONFIG_API_URL + "v1/methods/" + uuid)
      .then((response) => {
        commit(ApiMutationTypes.SET_ALL_METHODS, {
          state: RequestState.SUCCESS,
          data: response.data,
        });

        dispatch(ApiActionTypes.FETCH_METHODS);
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_DELETE_METHOD, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },
  async [ApiActionTypes.FETCH_METHODS]({ commit }) {
    commit(ApiMutationTypes.SET_ALL_METHODS, {
      state: RequestState.PENDING,
    });
    await axios
      .get(config.CONFIG_API_URL + "v1/methods")
      .then((response) => {
        commit(ApiMutationTypes.SET_ALL_METHODS, {
          state: RequestState.SUCCESS,
          data: response.data,
        });
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_ALL_METHODS, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },

  async [ApiActionTypes.CREATE_PHASE](
    { commit, dispatch },
    payload: CreatePhaseDto
  ) {
    commit(ApiMutationTypes.SET_CREATE_PHASE, {
      state: RequestState.PENDING,
    });
    await axios
      .post(config.CONFIG_API_URL + "v1/phases", payload)
      .then((response) => {
        commit(ApiMutationTypes.SET_CREATE_PHASE, {
          state: RequestState.SUCCESS,
          data: response.data,
        });

        dispatch(ApiActionTypes.FETCH_PHASES);
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_CREATE_PHASE, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },
  async [ApiActionTypes.MODIFY_PHASE]({ commit, dispatch }, payload: PhaseDto) {
    commit(ApiMutationTypes.SET_MODIFY_PHASE, {
      state: RequestState.PENDING,
    });
    await axios
      .put(config.CONFIG_API_URL + "v1/phases/" + payload.id, payload)
      .then((response) => {
        commit(ApiMutationTypes.SET_MODIFY_PHASE, {
          state: RequestState.SUCCESS,
          data: response.data,
        });

        dispatch(ApiActionTypes.FETCH_PHASES);
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_MODIFY_PHASE, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },
  async [ApiActionTypes.DELETE_PHASE]({ commit, dispatch }, uuid: string) {
    commit(ApiMutationTypes.SET_DELETE_PHASE, {
      state: RequestState.PENDING,
    });
    await axios
      .delete(config.CONFIG_API_URL + "v1/phases/" + uuid)
      .then((response) => {
        commit(ApiMutationTypes.SET_ALL_PHASES, {
          state: RequestState.SUCCESS,
          data: response.data,
        });

        dispatch(ApiActionTypes.FETCH_PHASES);
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_DELETE_PHASE, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },
  async [ApiActionTypes.FETCH_PHASES]({ commit }) {
    commit(ApiMutationTypes.SET_ALL_PHASES, {
      state: RequestState.PENDING,
    });
    await axios
      .get(config.CONFIG_API_URL + "v1/phases")
      .then((response) => {
        commit(ApiMutationTypes.SET_ALL_PHASES, {
          state: RequestState.SUCCESS,
          data: response.data,
        });
      })
      .catch((error: AxiosError) => {
        commit(ApiMutationTypes.SET_ALL_PHASES, {
          state: RequestState.ERROR,
          error: error,
        });
      });
  },
};
