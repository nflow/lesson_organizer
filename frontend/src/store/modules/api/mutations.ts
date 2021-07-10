import { ApiMutationsTypes, ApiStateTypes } from "@/store/interfaces";
import { MethodDto } from "@/types/MethodDto";
import { MutationTree } from "vuex";
import { ApiMutationTypes } from "./mutation-types";

export const mutations: MutationTree<ApiStateTypes> & ApiMutationsTypes = {
  [ApiMutationTypes.SET_METHODS](state, payload: MethodDto[]) {
    state.methods = payload;
  },
};
