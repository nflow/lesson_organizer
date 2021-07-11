import axios from "axios";
import { ActionContext, ActionTree } from "vuex";
import { ApiActionTypes } from "./action-types";
import { ApiMutationTypes } from "./mutation-types";
import config from "@/config";
import { ApiMutationsTypes } from "./mutations";
import { RootState } from "@/store";
import { ApiStateTypes } from ".";
import { CreateMethodDto } from "@/types/method";

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
    payload: CreateMethodDto
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
  [ApiActionTypes.CREATE_METHOD](
    { commit, dispatch },
    payload: CreateMethodDto
  ) {
    axios
      .post(config.CONFIG_API_URL + "/v1/methods", payload)
      .then((response) => {
        dispatch(ApiActionTypes.FETCH_METHODS);
      });
  },
  [ApiActionTypes.DELETE_METHOD]({ commit, dispatch }, id: string) {
    axios
      .delete(config.CONFIG_API_URL + "/v1/methods/" + id)
      .then((response) => {
        dispatch(ApiActionTypes.FETCH_METHODS);
      });
  },
};
