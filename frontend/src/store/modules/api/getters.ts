import { ApiGettersTypes, ApiStateTypes, IRootState } from "@/store/interfaces";
import { GetterTree } from "vuex";

export const getters: GetterTree<ApiStateTypes, IRootState> & ApiGettersTypes =
  {};
