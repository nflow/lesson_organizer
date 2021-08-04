import config from "@/config";
import { BoardMethodDto } from "@/types/method";
import axios from "axios";

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
