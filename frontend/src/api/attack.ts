import client from './client';
import { ATTACKMapping } from '../types';

export const getATTACKMappings = async (entityType?: string, entityId?: string): Promise<ATTACKMapping[]> => {
  const params = new URLSearchParams();
  if (entityType) params.append('entity_type', entityType);
  if (entityId) params.append('entity_id', entityId);

  const response = await client.get<ATTACKMapping[]>(`/attack/mappings?${params.toString()}`);
  return response.data;
};

export const createATTACKMapping = async (mapping: Partial<ATTACKMapping>): Promise<ATTACKMapping> => {
  const response = await client.post<ATTACKMapping>('/attack/mappings', mapping);
  return response.data;
};

export const getATTACKTechniques = async (): Promise<Record<string, ATTACKMapping[]>> => {
  const response = await client.get<Record<string, ATTACKMapping[]>>('/attack/techniques');
  return response.data;
};

export const deleteATTACKMapping = async (id: string): Promise<void> => {
  await client.delete(`/attack/mappings/${id}`);
};
