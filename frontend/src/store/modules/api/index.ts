import { Module } from "vuex";
import { ApiStateTypes, IRootState } from "@/store/interfaces";
import { getters } from "./getters";
import { actions } from "./actions";
import { mutations } from "./mutations";
import { state } from "./state";

const apiModule: Module<ApiStateTypes, IRootState> = {
  state,
  getters,
  mutations,
  actions,
};

export default apiModule;
