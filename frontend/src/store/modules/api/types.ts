import {
  ApiStateTypes,
  ApiMutationsTypes,
  ApiGettersTypes,
  ApiActionsTypes,
} from "@/store/interfaces";
import { Store as VuexStore, CommitOptions, DispatchOptions } from "vuex";

export type ApiStoreModuleTypes<S = ApiStateTypes> = Omit<
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
