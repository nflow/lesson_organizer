import { Request } from "@/types/api-state";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";
import { MutationTree } from "vuex";
import { ApiMutationTypes } from "./mutation-types";
import { ApiState } from "./state";

export type ApiMutationsTypes<S = ApiState> = {
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

  [ApiMutationTypes.SET_CREATE_PHASE](
    state: S,
    payload: Request<CreateMethodDto>
  ): void;
  [ApiMutationTypes.SET_MODIFY_PHASE](
    state: S,
    payload: Request<MethodDto>
  ): void;
  [ApiMutationTypes.SET_DELETE_PHASE](
    state: S,
    payload: Request<MethodDto>
  ): void;
  [ApiMutationTypes.SET_ALL_PHASES](
    state: S,
    payload: Request<MethodDto[]>
  ): void;
};

export const mutations: MutationTree<ApiState> & ApiMutationsTypes = {
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

  [ApiMutationTypes.SET_CREATE_PHASE](state, payload: Request<CreatePhaseDto>) {
    state.createPhase = payload;
  },
  [ApiMutationTypes.SET_MODIFY_PHASE](state, payload: Request<PhaseDto>) {
    state.modifyPhase = payload;
  },
  [ApiMutationTypes.SET_DELETE_PHASE](state, payload: Request<PhaseDto>) {
    state.deletePhase = payload;
  },
  [ApiMutationTypes.SET_ALL_PHASES](state, payload: Request<PhaseDto[]>) {
    state.allPhases = payload;
  },
};
