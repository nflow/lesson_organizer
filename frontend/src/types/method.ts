import { ContentDto } from "./content";

export enum Label {
  METHOD_LABEL_SINGLE = "METHOD_LABEL_SINGLE",
  METHOD_LABEL_PAIR = "METHOD_LABEL_PAIR",
  METHOD_LABEL_GROUP = "METHOD_LABEL_GROUP",
  METHOD_LABEL_PLENUM = "METHOD_LABEL_PLENUM",
}

export function resolveLabelName(label: Label): string {
  switch (label) {
    case Label.METHOD_LABEL_SINGLE:
      return "Single Person Working";
    case Label.METHOD_LABEL_PAIR:
      return "Partner Work";
    case Label.METHOD_LABEL_GROUP:
      return "Group Work";
    case Label.METHOD_LABEL_PLENUM:
      return "Plenum";
    default:
      return label;
  }
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
