import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";
import api, { ApiStore } from "./modules/api";
import board, { BoardStore } from "./modules/board";
import { ApiState } from "./modules/api/state";
import { BoardState } from "./modules/board/state";

export type RootState = {
  board: BoardState;
  api: ApiState;
};

export type Store = BoardStore<Pick<RootState, "board">> &
  ApiStore<Pick<RootState, "api">>;

export const store = createStore({
  modules: {
    api,
    board,
  },
  plugins: [createPersistedState()],
});

export function useStore(): Store {
  return store as Store;
}
