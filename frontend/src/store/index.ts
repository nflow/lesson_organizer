import { ContentDto } from "@/types/ContentDto";
import { InjectionKey } from "vue";
import {
  createStore,
  useStore as baseUseStore,
  MutationTree,
  Store,
  ActionTree,
  ActionContext,
} from "vuex";
import { IRootState } from "./interfaces";
import ApiModule from "./modules/api";
import { mutations } from "./modules/api/mutations";

export const key: InjectionKey<Store<IRootState>> = Symbol();

export default createStore<IRootState>({
  modules: {
    api: ApiModule,
  },
});

export function useStore(): Store<IRootState> {
  return baseUseStore(key);
}
