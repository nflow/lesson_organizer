import config from "@/config";
import {
  CreateBoardDto,
  MoveContentDto,
  MoveMethodDto,
  MovePhaseDto,
} from "@/types/board";
import { CreateContentDto } from "@/types/content";
import { MethodIdentifierDto } from "@/types/method";
import { PhaseIdentifierDto } from "@/types/phase";
import axios from "axios";
import { Ref } from "vue";
import { useMutation, useQuery } from "vue-query";
import { stringifyQuery } from "vue-router";

export function getBoard(boardId: Ref<string | string[]>) {
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
    {
      staleTime: 10000,
    }
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
  return useMutation(async (payload: MovePhaseDto) => {
    const { data } = await axios.put(
      `${config.CONFIG_API_URL}/v1/boards/${boardId.value}/phases`,
      payload
    );

    return data;
  });
}

export function postMethodAssociation(phaseId: string) {
  return useMutation(async (payload: MethodIdentifierDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/phases/${phaseId}/methods`,
      payload
    );

    return data;
  });
}

export function getPhaseMethods(phaseId: string) {
  return useQuery(
    ["phase_methods", phaseId],
    async () => {
      const { data } = await axios.get(
        `${config.CONFIG_API_URL}/v1/phases/${phaseId}/methods`
      );

      return data;
    },
    {
      staleTime: 10000,
    }
  );
}

export function putMethodOrder(phaseId: string) {
  return useMutation(async (payload: MoveMethodDto) => {
    const { data } = await axios.put(
      `${config.CONFIG_API_URL}/v1/phases/${phaseId}/methods`,
      payload
    );

    return data;
  });
}

export function getBoardContents(boardId: Ref<string | undefined>) {
  return useQuery(
    ["board_contents", boardId.value],
    async () => {
      const { data } = await axios.get(
        `${config.CONFIG_API_URL}/v1/boards/${boardId.value}/contents`
      );

      return data;
    },
    {
      staleTime: 10000,
    }
  );
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

export function getMethodContents(methodId: string) {
  return useQuery(
    ["method_contents", methodId],
    async () => {
      const { data } = await axios.get(
        `${config.CONFIG_API_URL}/v1/methods/${methodId}/contents`
      );

      return data;
    },
    {
      staleTime: 10000,
    }
  );
}

export function postMethodContent(methodId: string) {
  return useMutation(async (payload: CreateContentDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/methods/${methodId}/contents`,
      payload
    );

    return data;
  });
}

export function putContent() {
  return useMutation(async (payload: MoveContentDto) => {
    const { data } = await axios.put(
      `${config.CONFIG_API_URL}/v1/contents/${payload.contentId}`,
      payload
    );

    return data;
  });
}

export function deleteContent() {
  return useMutation(async (contentId: string) => {
    const { data } = await axios.delete(
      `${config.CONFIG_API_URL}/v1/contents/${contentId}`
    );

    return data;
  });
}

export function deletePhase() {
  return useMutation(async ([boardId, phaseId]: string[]) => {
    const { data } = await axios.delete(
      `${config.CONFIG_API_URL}/v1/boards/${boardId}/phases/${phaseId}`
    );

    return data;
  });
}

export function deleteMethod() {
  return useMutation(async ([boardId, phaseId, methodId]: string[]) => {
    const { data } = await axios.delete(
      `${config.CONFIG_API_URL}/v1/boards/${boardId}/phases/${phaseId}/methods/${methodId}`
    );

    return data;
  });
}
