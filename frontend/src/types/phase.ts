import { BoardMethodDto } from "./method";

export interface PhaseDto {
  id: string;
  title: string;
}

export interface BoardPhaseDto {
  id: string;
  phase: PhaseDto;
  methods: Array<BoardMethodDto>;
  order: number;
}

export interface CreatePhaseDto {
  title: string;
}

export interface PhaseIdentifierDto {
  id: string;
}
