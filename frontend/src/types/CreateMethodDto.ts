import { Label } from "./Label";

export interface CreateMethodDto {
  title: string;
  description: string;
  labels: Array<Label>;
}
