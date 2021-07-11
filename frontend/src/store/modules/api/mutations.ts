import { MethodDto } from "@/types/method";
import { MutationTree } from "vuex";
import { ApiStateTypes } from ".";
import { ApiMutationTypes } from "./mutation-types";

export type ApiMutationsTypes<S = ApiStateTypes> = {
  [ApiMutationTypes.SET_METHODS](state: S, payload: MethodDto[]): void;
};

export const mutations: MutationTree<ApiStateTypes> & ApiMutationsTypes = {
  [ApiMutationTypes.SET_METHODS](state, payload: MethodDto[]) {
    state.methods = payload;
  },
};
