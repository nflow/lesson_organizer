import { Label } from "./Label";

export interface BaseMethodDto {
  id: string;
  title: string;
  description: string;
  labels: Array<Label>;
}
