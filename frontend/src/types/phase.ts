import { BoardMethodDto } from "./method";

export interface PhaseDto {
  id: string;
  title: string;
}

export interface BoardPhaseDto extends PhaseDto {
  methods: Array<BoardMethodDto>;
}

export interface CreatePhaseDto {
  title: string;
}
