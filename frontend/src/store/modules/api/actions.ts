import axios, { AxiosError } from "axios";
import { ActionContext, ActionTree } from "vuex";
import { ApiActionTypes } from "./action-types";
import { ApiMutationTypes } from "./mutation-types";
import config from "@/config";
import { ApiMutationsTypes } from "./mutations";
import { RootState } from "@/store";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { RequestState } from "@/types/api-state";
import { ApiStateTypes } from "./state";

export type AugmentedActionContext = {
  commit<K extends keyof ApiMutationsTypes>(
    key: K,
    payload: Parameters<ApiMutationsTypes[K]>[1]
  ): ReturnType<ApiMutationsTypes[K]>;
} & Omit<ActionContext<ApiStateTypes, RootState>, "commit">;

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
}

export const actions: ActionTree<ApiStateTypes, RootState> & ApiActionsTypes = {
  async [ApiActionTypes.CREATE_METHOD](
    { commit, dispatch },
    payload: CreateMethodDto
  ) {
    commit(ApiMutationTypes.SET_CREATE_METHOD, {
      state: RequestState.PENDING,
    });
    await axios
      .post(config.CONFIG_API_URL + "/v1/methods", payload)
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
      .put(config.CONFIG_API_URL + "/v1/methods/" + payload.id, payload)
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
      .delete(config.CONFIG_API_URL + "/v1/methods/" + uuid)
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
      .get(config.CONFIG_API_URL + "/v1/methods")
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
};
