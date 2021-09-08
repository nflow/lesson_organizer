import config from "@/config";
import { CreateBoardDto, MoveElementDto } from "@/types/board";
import { CreateContentDto } from "@/types/content";
import { MethodIdentifierDto } from "@/types/method";
import { PhaseIdentifierDto } from "@/types/phase";
import axios from "axios";
import { Ref } from "vue";
import { useMutation, useQuery } from "vue-query";
import type { UseQueryOptions } from "react-query/types/react/types";

export function getBoard<
  TQueryFnData = unknown,
  TError = unknown,
  TData = TQueryFnData
>(
  boardId: Ref<string | string[]>,
  options: UseQueryOptions<TQueryFnData, TError, TData>
) {
  return useQuery(
    "board",
    async () => {
      if (boardId.value) {
        const { data } = await axios.get(
          `${config.CONFIG_API_URL}/v1/boards/${boardId.value}`
        );

        return data;
      }
    },
    options
  );
}

export function postBoard() {
  return useMutation(async (payload: CreateBoardDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/boards`,
      payload
    );

    return data;
  });
}

export function postPhaseAssociation(boardId: Ref<string | string[]>) {
  return useMutation(async (payload: PhaseIdentifierDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/boards/${boardId.value}/phases`,
      payload
    );

    return data;
  });
}

export function putPhaseOrder(boardId: Ref<string | string[]>) {
  return useMutation(async (payload: MoveElementDto) => {
    const { data } = await axios.put(
      `${config.CONFIG_API_URL}/v1/boards/${boardId.value}/phases`,
      payload
    );

    return data;
  });
}

export function postMethodAssociation(
  boardId: Ref<string | string[]>,
  phaseId: string
) {
  return useMutation(async (payload: MethodIdentifierDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/boards/${boardId.value}/phases/${phaseId}/methods`,
      payload
    );

    return data;
  });
}

export function postBoardContent(boardId: Ref<string | undefined>) {
  return useMutation(async (payload: CreateContentDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/boards/${boardId.value}/contents`,
      payload
    );

    return data;
  });
}

export function postMethodContent(
  boardId: Ref<string | undefined>,
  phaseId: Ref<string>,
  methodId: Ref<string>
) {
  return useMutation(async (payload: CreateContentDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/boards/${boardId.value}/phases/${phaseId.value}/methods/${methodId.value}/contents`,
      payload
    );

    return data;
  });
}
