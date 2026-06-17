import client from './client';
import { Campaign } from '../types';

export const getCampaigns = async (threatActorId?: string): Promise<Campaign[]> => {
  const url = threatActorId ? `/campaigns?threat_actor_id=${threatActorId}` : '/campaigns';
  const response = await client.get<Campaign[]>(url);
  return response.data;
};

export const getCampaign = async (id: string): Promise<Campaign> => {
  const response = await client.get<Campaign>(`/campaigns/${id}`);
  return response.data;
};

export const createCampaign = async (campaign: Partial<Campaign>): Promise<Campaign> => {
  const response = await client.post<Campaign>('/campaigns', campaign);
  return response.data;
};

export const updateCampaign = async (id: string, campaign: Partial<Campaign>): Promise<Campaign> => {
  const response = await client.put<Campaign>(`/campaigns/${id}`, campaign);
  return response.data;
};

export const deleteCampaign = async (id: string): Promise<void> => {
  await client.delete(`/campaigns/${id}`);
};
