import { ContentDto } from "./content";

export enum Label {
  METHOD_LABEL_SINGLE = "METHOD_LABEL_SINGLE",
  METHOD_LABEL_PAIR = "METHOD_LABEL_PAIR",
  METHOD_LABEL_GROUP = "METHOD_LABEL_GROUP",
  METHOD_LABEL_PLENUM = "METHOD_LABEL_PLENUM",
}

export interface MethodDto {
  id: string;
  title: string;
  description: string;
  category: string;
  labels: Array<Label>;
}

export interface BoardMethodDto extends MethodDto {
  ideas: Array<ContentDto>;
}

export interface CreateMethodDto {
  title: string;
  description: string;
  category: string;
  labels: Array<Label>;
}
