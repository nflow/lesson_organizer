import { ContentDto } from "./content";
import { GoalDto } from "./goal";
import { BoardPhaseDto } from "./phase";

export interface BoardDto {
  id: string;
  name: string;
  goals: GoalDto[];
  contents: ContentDto[];
  phases: BoardPhaseDto[];
}
