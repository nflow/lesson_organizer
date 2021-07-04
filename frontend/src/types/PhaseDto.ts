import { MethodDto } from "./MethodDto";

export interface PhaseDto {
  id: string;
  title: string;
  methods: Array<MethodDto>;
}
