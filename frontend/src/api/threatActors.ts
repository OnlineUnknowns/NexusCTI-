import client from './client';
import { ThreatActor } from '../types';

export const getThreatActors = async (): Promise<ThreatActor[]> => {
  const response = await client.get<ThreatActor[]>('/threat-actors');
  return response.data;
};

export const getThreatActor = async (id: string): Promise<ThreatActor> => {
  const response = await client.get<ThreatActor>(`/threat-actors/${id}`);
  return response.data;
};

export const createThreatActor = async (actor: Partial<ThreatActor>): Promise<ThreatActor> => {
  const response = await client.post<ThreatActor>('/threat-actors', actor);
  return response.data;
};

export const updateThreatActor = async (id: string, actor: Partial<ThreatActor>): Promise<ThreatActor> => {
  const response = await client.put<ThreatActor>(`/threat-actors/${id}`, actor);
  return response.data;
};

export const deleteThreatActor = async (id: string): Promise<void> => {
  await client.delete(`/threat-actors/${id}`);
};
