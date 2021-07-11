import { Request } from "@/types/api-state";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { MutationTree } from "vuex";
import { ApiMutationTypes } from "./mutation-types";
import { ApiStateTypes } from "./state";

export type ApiMutationsTypes<S = ApiStateTypes> = {
  [ApiMutationTypes.SET_CREATE_METHOD](
    state: S,
    payload: Request<CreateMethodDto>
  ): void;
  [ApiMutationTypes.SET_MODIFY_METHOD](
    state: S,
    payload: Request<MethodDto>
  ): void;
  [ApiMutationTypes.SET_DELETE_METHOD](
    state: S,
    payload: Request<MethodDto>
  ): void;
  [ApiMutationTypes.SET_ALL_METHODS](
    state: S,
    payload: Request<MethodDto[]>
  ): void;
};

export const mutations: MutationTree<ApiStateTypes> & ApiMutationsTypes = {
  [ApiMutationTypes.SET_CREATE_METHOD](
    state,
    payload: Request<CreateMethodDto>
  ) {
    state.createMethod = payload;
  },
  [ApiMutationTypes.SET_MODIFY_METHOD](state, payload: Request<MethodDto>) {
    state.modifyMethod = payload;
  },
  [ApiMutationTypes.SET_DELETE_METHOD](state, payload: Request<MethodDto>) {
    state.deleteMethod = payload;
  },
  [ApiMutationTypes.SET_ALL_METHODS](state, payload: Request<MethodDto[]>) {
    state.allMethods = payload;
  },
};
