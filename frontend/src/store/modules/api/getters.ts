import { RootState } from "@/store";
import { MethodDto } from "@/types/method";
import { PhaseDto } from "@/types/phase";
import { GetterTree } from "vuex";
import { ApiState } from "./state";

export interface ApiGettersTypes {
  allMethods(state: ApiState): MethodDto[];
  allPhases(state: ApiState): PhaseDto[];
}

export const getters: GetterTree<ApiState, RootState> & ApiGettersTypes = {
  allMethods: (state: ApiState) => {
    return state.allMethods && state.allMethods.data
      ? state.allMethods.data
      : [];
  },
  allPhases: (state: ApiState) => {
    return state.allPhases && state.allPhases.data ? state.allPhases.data : [];
  },
};
