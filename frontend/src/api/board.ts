import config from "@/config";
import { CreateBoardDto } from "@/types/board";
import { MethodIdentifierDto } from "@/types/method";
import { PhaseIdentifierDto } from "@/types/phase";
import axios from "axios";
import { Ref } from "vue";
import { useMutation, useQuery } from "vue-query";

export function getBoard(boardId: Ref<string | string[]>) {
  return useQuery("board", async () => {
    if (boardId.value) {
      const { data } = await axios.get(
        `${config.CONFIG_API_URL}/v1/boards/${boardId.value}`
      );

      return data;
    }
  });
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
