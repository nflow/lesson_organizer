import { Request } from "@/types/api-state";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";

export interface ApiState {
  createMethod?: Request<CreateMethodDto>;
  modifyMethod?: Request<MethodDto>;
  deleteMethod?: Request<MethodDto>;
  allMethods?: Request<MethodDto[]>;

  createPhase?: Request<CreatePhaseDto>;
  modifyPhase?: Request<PhaseDto>;
  deletePhase?: Request<PhaseDto>;
  allPhases?: Request<PhaseDto[]>;
}

const state: ApiState = {
  createMethod: undefined,
  modifyMethod: undefined,
  deleteMethod: undefined,
  allMethods: undefined,

  createPhase: undefined,
  modifyPhase: undefined,
  deletePhase: undefined,
  allPhases: undefined,
};

export default state;
