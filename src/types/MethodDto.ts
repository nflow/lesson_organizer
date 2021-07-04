import { ContentDto } from "./ContentDto";

export interface MethodDto {
  id: string;
  title: string;
  description: string;
  ideas: Array<ContentDto>;
}
