import { RootState } from "@/store";
import { MethodDto } from "@/types/method";
import { PhaseDto } from "@/types/phase";
import { GetterTree } from "vuex";
import { ApiStateTypes } from "./state";

export interface ApiGettersTypes {
  allMethods(state: ApiStateTypes): MethodDto[];
  allPhases(state: ApiStateTypes): PhaseDto[];
}

export const getters: GetterTree<ApiStateTypes, RootState> & ApiGettersTypes = {
  allMethods: (state: ApiStateTypes) => {
    return state.allMethods && state.allMethods.data
      ? state.allMethods.data
      : [];
  },
  allPhases: (state: ApiStateTypes) => {
    return state.allPhases && state.allPhases.data ? state.allPhases.data : [];
  },
};
