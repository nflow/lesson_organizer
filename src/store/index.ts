import { IdeaDto } from "@/types/IdeaDto";
import { MethodDto } from "@/types/MethodDto";
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

export interface DragCallback<T> {
  id: string;
  callback: (id: string) => T;
}

interface State {
  dragCallbackIdea?: DragCallback<IdeaDto>;
  dragCallbackMethod?: DragCallback<MethodDto>;
}

export enum MutationTypes {
  SET_DRAG_CALLBACK_IDEA = "SET_DRAG_CALLBACK_IDEA",
  SET_DRAG_CALLBACK_METHOD = "SET_DRAG_CALLBACK_METHOD",
}

export type Mutations<S = State> = {
  [MutationTypes.SET_DRAG_CALLBACK_IDEA](
    state: S,
    payload: DragCallback<IdeaDto>
  ): void;
  [MutationTypes.SET_DRAG_CALLBACK_METHOD](
    state: S,
    payload: DragCallback<MethodDto>
  ): void;
};

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_DRAG_CALLBACK_IDEA](
    state,
    payload: DragCallback<IdeaDto>
  ) {
    state.dragCallbackIdea = payload;
  },
  [MutationTypes.SET_DRAG_CALLBACK_METHOD](
    state,
    payload: DragCallback<MethodDto>
  ) {
    state.dragCallbackMethod = payload;
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
    dragCallbackIdea: undefined,
    dragCallbackMethod: undefined,
  },
  mutations: mutations,
  actions: actions,
  modules: {},
});

export function useStore(): Store<State> {
  return baseUseStore(key);
}
