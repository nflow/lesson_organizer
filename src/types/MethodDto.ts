import { IdeaDto } from "./IdeaDto";

export interface MethodDto {
  id: string;
  title: string;
  description: string;
  ideas: Array<IdeaDto>;
}
