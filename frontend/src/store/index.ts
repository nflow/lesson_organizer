import { InjectionKey } from "vue";
import {
  createStore,
  useStore as vuexUseStore,
  Store as VuexStore,
  ModuleTree,
} from "vuex";
import { IRootState } from "./interfaces";
import apiModule from "./modules/api";
import { ApiStoreModuleTypes } from "./modules/api/types";

const modules: ModuleTree<IRootState> = {
  apiModule,
};

export default createStore<IRootState>({
  modules: modules,
});

type StoreModules = {
  apiModule: ApiStoreModuleTypes;
};

export type Store = ApiStoreModuleTypes<Pick<StoreModules, "apiModule">>;

export const key: InjectionKey<VuexStore<IRootState>> = Symbol();

export function useStore(): VuexStore<IRootState> {
  return vuexUseStore(key);
}
