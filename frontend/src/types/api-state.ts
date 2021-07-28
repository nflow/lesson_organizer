import { AxiosError } from "axios";

export enum RequestState {
  PENDING,
  SUCCESS,
  ERROR,
}
export interface Request<T> {
  state: RequestState;
  data?: T;
  error?: AxiosError;
}
