import { IdeaDto } from "@/types/IdeaDto";
import { validate } from "uuid";
import { InjectionKey } from "vue";
import {
  createStore,
  useStore as baseUseStore,
  MutationTree,
  Store,
  ActionTree,
  ActionContext,
  mapGetters,
  GetterTree,
} from "vuex";

export interface DragCallback {
  id: string;
  callback: (id: string) => IdeaDto;
}

interface State {
  dragCallback?: DragCallback;
}

export enum MutationTypes {
  SET_DRAG_CALLBACK = "SET_DRAG_CALLBACK",
}

export type Mutations<S = State> = {
  [MutationTypes.SET_DRAG_CALLBACK](state: S, payload: DragCallback): void;
};

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_DRAG_CALLBACK](state, payload: DragCallback) {
    state.dragCallback = payload;
  },
};

export enum ActionTypes {
  FETCH_CARDS = "FETCH_CARDS",
}

type AugmentedActionContext = {
  commit<K extends keyof Mutations>(
    key: K,
    payload: Parameters<Mutations[K]>[1]
  ): ReturnType<Mutations[K]>;
} & Omit<ActionContext<State, State>, "commit">;

export interface Actions {
  [ActionTypes.FETCH_CARDS](
    { commit }: AugmentedActionContext,
    payload: string
  ): Promise<void>;
}

export const actions: ActionTree<State, State> & Actions = {
  async [ActionTypes.FETCH_CARDS]({ commit }, payload) {
    return Promise.resolve();
  },
};

export const key: InjectionKey<Store<State>> = Symbol();

export default createStore<State>({
  state: {
    dragCallback: undefined,
  },
  mutations: mutations,
  actions: actions,
  modules: {},
});

export function useStore(): Store<State> {
  return baseUseStore(key);
}
