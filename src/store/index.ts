import { IdeaDto } from "@/types/IdeaDto";
import { InjectionKey } from "vue";
import {
  createStore,
  useStore as baseUseStore,
  MutationTree,
  Store,
  ActionTree,
  ActionContext,
} from "vuex";

interface State {
  ideas: IdeaDto[];
}

export enum MutationTypes {
  SET_IDEAS = "SET_IDEAS",
}

export type Mutations<S = State> = {
  [MutationTypes.SET_IDEAS](state: S, payload: IdeaDto[]): void;
};

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_IDEAS](state, payload: IdeaDto[]) {
    state.ideas = payload;
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
    ideas: [],
  },
  mutations: mutations,
  actions: actions,
  modules: {},
});

export function useStore(): Store<State> {
  return baseUseStore(key);
}
