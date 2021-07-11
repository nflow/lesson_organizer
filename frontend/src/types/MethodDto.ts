import { BaseMethodDto } from "./BaseMethodDto";
import { ContentDto } from "./ContentDto";

export interface MethodDto extends BaseMethodDto {
  ideas: Array<ContentDto>;
}
