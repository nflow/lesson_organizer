import config from "@/config";
import { CreateBoardDto } from "@/types/board";
import { PhaseIdentifierDto } from "@/types/phase";
import axios from "axios";
import { useMutation, useQuery } from "vue-query";

export function getBoard(boardId: string | string[]) {
  return useQuery("board", async () => {
    const { data } = await axios.get(
      `${config.CONFIG_API_URL}/v1/boards/${boardId}`
    );

    return data;
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

export function postPhaseAssociation(boardId: string | string[]) {
  return useMutation(async (payload: PhaseIdentifierDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/boards/${boardId}/phases`,
      payload
    );

    return data;
  });
}
