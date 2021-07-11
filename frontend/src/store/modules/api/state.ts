import { Request } from "@/types/api-state";
import { CreateMethodDto, MethodDto } from "@/types/method";

export interface ApiStateTypes {
  createMethod?: Request<CreateMethodDto>;
  modifyMethod?: Request<MethodDto>;
  deleteMethod?: Request<MethodDto>;
  allMethods?: Request<MethodDto[]>;
}

const state: ApiStateTypes = {
  createMethod: undefined,
  modifyMethod: undefined,
  deleteMethod: undefined,
  allMethods: undefined,
};

export default state;
