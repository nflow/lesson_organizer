import { Request } from "@/types/api-state";
import { BoardDto } from "@/types/board";
import { MutationTree } from "vuex";
import { BoardMutationTypes } from "./mutation-types";
import { BoardState } from "./state";

export type BoardMutationsTypes<S = BoardState> = {
  [BoardMutationTypes.SET_CURRENT_BOARD](
    state: S,
    payload: Request<BoardDto>
  ): void;
};

export const mutations: MutationTree<BoardState> & BoardMutationsTypes = {
  [BoardMutationTypes.SET_CURRENT_BOARD](state, payload: Request<BoardDto>) {
    state.currentBoard = payload;
  },
};
