import config from "@/config";
import { BoardDto, CreateBoardDto } from "@/types/board";
import { BoardMethodDto } from "@/types/method";
import axios from "axios";

export async function getBoard(boardId: string | string[]): Promise<BoardDto> {
  const { data } = await axios.get(
    `${config.CONFIG_API_URL}/v1/boards/${boardId}`
  );

  return data;
}

export async function postBoard(payload: CreateBoardDto): Promise<BoardDto> {
  const { data } = await axios.post(
    `${config.CONFIG_API_URL}/v1/boards`,
    payload
  );

  return data;
}

export async function getBoardMethods({
  boardId,
  phaseId,
}: {
  boardId: string;
  phaseId: string;
}): Promise<BoardMethodDto[]> {
  const { data } = await axios.get(
    `${config.CONFIG_API_URL}/v1/boards/${boardId}/phases/${phaseId}`
  );

  return data;
}
