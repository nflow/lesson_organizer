import config from "@/config";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";
import axios from "axios";

export async function getMethods(): Promise<MethodDto[]> {
  const { data } = await axios.get(`${config.CONFIG_API_URL}/v1/methods`);

  return data;
}

export async function postMethod(
  payload: CreateMethodDto
): Promise<MethodDto[]> {
  const { data } = await axios.post(
    `${config.CONFIG_API_URL}/v1/methods`,
    payload
  );

  return data;
}

export async function putMethod(payload: MethodDto): Promise<MethodDto> {
  const { data } = await axios.put(
    `${config.CONFIG_API_URL}/v1/methods/${payload.id}`,
    payload
  );

  return data;
}

export async function deleteMethod(id: string): Promise<MethodDto> {
  const { data } = await axios.delete(
    `${config.CONFIG_API_URL}/v1/methods/${id}`
  );

  return data;
}

export async function getPhases(): Promise<PhaseDto[]> {
  const { data } = await axios.get(`${config.CONFIG_API_URL}/v1/phases`);

  return data;
}

export async function postPhase(payload: CreatePhaseDto): Promise<PhaseDto[]> {
  const { data } = await axios.post(
    `${config.CONFIG_API_URL}/v1/phases`,
    payload
  );

  return data;
}

export async function putPhase(payload: PhaseDto): Promise<PhaseDto> {
  const { data } = await axios.put(
    `${config.CONFIG_API_URL}/v1/phases/${payload.id}`,
    payload
  );

  return data;
}

export async function deletePhase(id: string): Promise<PhaseDto> {
  const { data } = await axios.delete(
    `${config.CONFIG_API_URL}/v1/phases/${id}`
  );

  return data;
}
