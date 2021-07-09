import { ContentDto } from "./ContentDto";

export interface MethodDto {
  id: string;
  title: string;
  description: string;
  labels: Array<Label>;
  ideas: Array<ContentDto>;
}

enum Label {
  METHOD_LABEL_SINGLE = "METHOD_LABEL_SINGLE",
  METHOD_LABEL_PAIR = "METHOD_LABEL_PAIR",
  METHOD_LABEL_GROUP = "METHOD_LABEL_GROUP",
  METHOD_LABEL_PLENUM = "METHOD_LABEL_PLENUM",
}
