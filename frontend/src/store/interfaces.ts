import { MethodDto } from "@/types/MethodDto";
import { ActionContext } from "vuex";
import { ApiActionTypes } from "./modules/api/action-types";
import { ApiMutationTypes } from "./modules/api/mutation-types";

export interface IRootState {}

export interface ApiStateTypes {
  methods: MethodDto[];
}

export interface ApiGettersTypes {}

export type ApiMutationsTypes<S = ApiStateTypes> = {
  [ApiMutationTypes.SET_METHODS](state: S, payload: MethodDto[]): void;
};

export type AugmentedActionContext = {
  commit<K extends keyof ApiMutationsTypes>(
    key: K,
    payload: Parameters<ApiMutationsTypes[K]>[1]
  ): ReturnType<ApiMutationsTypes[K]>;
} & Omit<ActionContext<ApiStateTypes, IRootState>, "commit">;

export interface ApiActionsTypes {
  [ApiActionTypes.FETCH_METHODS](
    { commit }: AugmentedActionContext,
    payload: number
  ): void;
  [ApiActionTypes.CREATE_METHOD](
    { commit }: AugmentedActionContext,
    payload: MethodDto
  ): void;
  [ApiActionTypes.DELETE_METHOD](
    { commit }: AugmentedActionContext,
    payload: string
  ): void;
}

export interface StoreActions extends ApiActionsTypes {}

export interface StoreGetters extends ApiGettersTypes {}
