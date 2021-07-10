import axios from "axios";
import { MethodDto } from "@/types/MethodDto";
import { ActionContext, ActionTree } from "vuex";
import { ApiActionTypes } from "./action-types";
import { ApiMutationTypes } from "./mutation-types";
import config from "@/config";
import { ApiMutationsTypes } from "./mutations";
import { RootState } from "@/store";
import { ApiStateTypes } from ".";

export type AugmentedActionContext = {
  commit<K extends keyof ApiMutationsTypes>(
    key: K,
    payload: Parameters<ApiMutationsTypes[K]>[1]
  ): ReturnType<ApiMutationsTypes[K]>;
} & Omit<ActionContext<ApiStateTypes, RootState>, "commit">;

export interface ApiActionsTypes {
  [ApiActionTypes.FETCH_METHODS](
    { commit }: AugmentedActionContext,
    payload: number
  ): void;
  [ApiActionTypes.CREATE_METHOD](
    { commit }: AugmentedActionContext,
    payload: MethodDto
  ): void;
  [ApiActionTypes.DELETE_METHOD](
    { commit }: AugmentedActionContext,
    payload: string
  ): void;
}

export const actions: ActionTree<ApiStateTypes, RootState> & ApiActionsTypes = {
  [ApiActionTypes.FETCH_METHODS]({ commit }) {
    axios.get(config.CONFIG_API_URL + "/v1/methods").then((response) => {
      commit(ApiMutationTypes.SET_METHODS, response.data);
    });
  },
  [ApiActionTypes.CREATE_METHOD]({ commit }, payload: MethodDto) {
    axios.get(config.CONFIG_API_URL + "/v1/methods").then((response) => {
      commit(ApiMutationTypes.SET_METHODS, response.data);
    });
  },
  [ApiActionTypes.DELETE_METHOD]({ commit }, payload: string) {
    console.log("payload", payload);
    console.log("commit", commit);
  },
};
