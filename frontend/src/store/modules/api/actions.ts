import { ApiActionsTypes, ApiStateTypes, IRootState } from "@/store/interfaces";
import axios from "axios";
import { MethodDto } from "@/types/MethodDto";
import { ActionTree } from "vuex";
import { ApiActionTypes } from "./action-types";
import { ApiMutationTypes } from "./mutation-types";
import config from "@/config";

export const actions: ActionTree<ApiStateTypes, IRootState> & ApiActionsTypes =
  {
    [ApiActionTypes.FETCH_METHODS]({ commit }) {
      axios.get(config.CONFIG_API_URL + "/v1/methods").then((response) => {
        commit(ApiMutationTypes.SET_METHODS, response.data);
      });
    },
    [ApiActionTypes.CREATE_METHOD]({ commit }, payload: MethodDto) {
      axios.get(config.CONFIG_API_URL + "/v1/methods").then((response) => {
        commit(ApiMutationTypes.SET_METHODS, response.data);
      });
    },
    [ApiActionTypes.DELETE_METHOD]({ commit }, payload: string) {
      console.log("payload", payload);
      console.log("commit", commit);
    },
  };
