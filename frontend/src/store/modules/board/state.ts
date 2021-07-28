import { Request } from "@/types/api-state";
import { BoardDto } from "@/types/board";

export interface BoardState {
  currentBoard?: Request<BoardDto>;
}

const state: BoardState = {
  currentBoard: undefined,
};

export default state;
