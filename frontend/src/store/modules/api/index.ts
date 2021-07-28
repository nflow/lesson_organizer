import {
  CommitOptions,
  DispatchOptions,
  Module,
  Store as VuexStore,
} from "vuex";
import { ApiGettersTypes, getters } from "./getters";
import { actions, ApiActionsTypes } from "./actions";
import { ApiMutationsTypes, mutations } from "./mutations";
import { RootState } from "@/store";
import state, { ApiState } from "./state";

export type ApiStore<S = ApiState> = Omit<
  VuexStore<S>,
  "commit" | "getters" | "dispatch"
> & {
  commit<
    K extends keyof ApiMutationsTypes,
    P extends Parameters<ApiMutationsTypes[K]>[1]
  >(
    key: K,
    payload?: P,
    options?: CommitOptions
  ): ReturnType<ApiMutationsTypes[K]>;
} & {
  getters: {
    [K in keyof ApiGettersTypes]: ReturnType<ApiGettersTypes[K]>;
  };
} & {
  dispatch<K extends keyof ApiActionsTypes>(
    key: K,
    payload?: Parameters<ApiActionsTypes[K]>[1],
    options?: DispatchOptions
  ): ReturnType<ApiActionsTypes[K]>;
};

const api: Module<ApiState, RootState> = {
  state,
  getters,
  mutations,
  actions,
};

export default api;
