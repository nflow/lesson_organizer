import config from "@/config";
import { CreateMethodDto, MethodDto } from "@/types/method";
import { CreatePhaseDto, PhaseDto } from "@/types/phase";
import axios from "axios";
import { useMutation, useQuery } from "vue-query";

export function getMethods() {
  return useQuery(
    "methods",
    async () => {
      const { data } = await axios.get(`${config.CONFIG_API_URL}/v1/methods`);

      return data;
    },
    {
      staleTime: 60000,
    }
  );
}

export function postMethod() {
  return useMutation(async (payload: CreateMethodDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/methods`,
      payload
    );

    return data;
  });
}

export function putMethod() {
  return useMutation(async (payload: MethodDto) => {
    const { data } = await axios.put(
      `${config.CONFIG_API_URL}/v1/methods/${payload.id}`,
      payload
    );

    return data;
  });
}

export function deleteMethod() {
  return useMutation(async (id: string) => {
    const { data } = await axios.delete(
      `${config.CONFIG_API_URL}/v1/methods/${id}`
    );

    return data;
  });
}

export function getPhases() {
  return useQuery(
    "phases",
    async () => {
      const { data } = await axios.get(`${config.CONFIG_API_URL}/v1/phases`);

      return data;
    },
    {
      staleTime: 60000,
    }
  );
}

export function postPhase() {
  return useMutation(async (payload: CreatePhaseDto) => {
    const { data } = await axios.post(
      `${config.CONFIG_API_URL}/v1/phases`,
      payload
    );

    return data;
  });
}

export function putPhase() {
  return useMutation(async (payload: PhaseDto) => {
    const { data } = await axios.put(
      `${config.CONFIG_API_URL}/v1/phases/${payload.id}`,
      payload
    );

    return data;
  });
}

export function deletePhase() {
  return useMutation(async (id: string) => {
    const { data } = await axios.delete(
      `${config.CONFIG_API_URL}/v1/phases/${id}`
    );

    return data;
  });
}
