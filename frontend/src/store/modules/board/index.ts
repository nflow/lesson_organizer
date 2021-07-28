import { CommitOptions, Module, Store as VuexStore } from "vuex";
import { BoardMutationsTypes, mutations } from "./mutations";
import { RootState } from "@/store";
import state, { BoardState } from "./state";

export type BoardStore<S = BoardState> = Omit<
  VuexStore<S>,
  "commit" | "getters" | "dispatch"
> & {
  commit<
    K extends keyof BoardMutationsTypes,
    P extends Parameters<BoardMutationsTypes[K]>[1]
  >(
    key: K,
    payload?: P,
    options?: CommitOptions
  ): ReturnType<BoardMutationsTypes[K]>;
};

const board: Module<BoardState, RootState> = {
  state,
  mutations,
};

export default board;
