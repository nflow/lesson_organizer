import {
  CommitOptions,
  DispatchOptions,
  Module,
  Store as VuexStore,
} from "vuex";
import { getters } from "./getters";
import { actions, ApiActionsTypes } from "./actions";
import { ApiMutationsTypes, mutations } from "./mutations";
import { MethodDto } from "@/types/MethodDto";
import { RootState } from "@/store";
import state from "./state";

export interface ApiStateTypes {
  methods: MethodDto[];
}

// eslint-disable-next-line @typescript-eslint/no-empty-interface
export interface ApiGettersTypes {}

export type ApiStore<S = ApiStateTypes> = Omit<
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

const api: Module<ApiStateTypes, RootState> = {
  state,
  getters,
  mutations,
  actions,
};

export default api;
