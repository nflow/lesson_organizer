import { createStore } from "vuex";
import api, { ApiStore } from "./modules/api";
import { ApiStateTypes } from "./modules/api/state";

export type RootState = {
  api: ApiStateTypes;
};

export type Store = ApiStore<Pick<RootState, "api">>;

export const store = createStore({
  modules: {
    api,
  },
});

export function useStore(): Store {
  return store as Store;
}
