import { MethodDto } from "./MethodDto";

export interface PhaseDto {
  id: string;
  title: string;
}

export interface BoardPhaseDto extends PhaseDto {
  methods: Array<MethodDto>;
}

export interface CreatePhaseDto {
  title: string;
}
