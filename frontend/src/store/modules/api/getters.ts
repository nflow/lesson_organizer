import { RootState } from "@/store";
import { GetterTree } from "vuex";
import { ApiGettersTypes, ApiStateTypes } from ".";

export const getters: GetterTree<ApiStateTypes, RootState> & ApiGettersTypes =
  {};
