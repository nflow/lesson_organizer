import { RootState } from "@/store";
import { MethodDto } from "@/types/method";
import { GetterTree } from "vuex";
import { ApiStateTypes } from "./state";

export interface ApiGettersTypes {
  allMethods(state: ApiStateTypes): MethodDto[];
}

export const getters: GetterTree<ApiStateTypes, RootState> & ApiGettersTypes = {
  allMethods: (state: ApiStateTypes) => {
    return state.allMethods && state.allMethods.data
      ? state.allMethods.data
      : [];
  },
};
