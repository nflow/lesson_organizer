import { ContentDto } from "./content";
import { GoalDto } from "./goal";
import { BoardPhaseDto } from "./phase";

export interface CreateBoardDto {
  name: string;
  goals: GoalDto[];
  contents: ContentDto[];
  phases: BoardPhaseDto[];
}
export interface BoardDto {
  id: string;
  name: string;
  goals: GoalDto[];
  contents: ContentDto[];
  phases: BoardPhaseDto[];
}

export interface MovePhaseDto {
  phaseId: string;
  afterPhaseId: string | undefined;
}

export interface MoveMethodDto {
  methodId: string;
  parentPhaseId: string | undefined;
  afterMethodId: string | undefined;
}
