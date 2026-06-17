import client from './client';
import { IOC, PaginatedResponse } from '../types';

export interface IOCFilters {
  type?: string;
  tlp?: string;
  tag?: string;
  q?: string;
}

export const getIOCs = async (filters: IOCFilters, page = 1, limit = 20): Promise<PaginatedResponse<IOC>> => {
  const params = new URLSearchParams();
  if (filters.type) params.append('type', filters.type);
  if (filters.tlp) params.append('tlp', filters.tlp);
  if (filters.tag) params.append('tag', filters.tag);
  if (filters.q) params.append('q', filters.q);
  params.append('page', page.toString());
  params.append('limit', limit.toString());

  const response = await client.get<PaginatedResponse<IOC>>(`/iocs?${params.toString()}`);
  return response.data;
};

export const getIOC = async (id: string): Promise<IOC> => {
  const response = await client.get<IOC>(`/iocs/${id}`);
  return response.data;
};

export const createIOC = async (ioc: Partial<IOC>): Promise<IOC> => {
  const response = await client.post<IOC>('/iocs', ioc);
  return response.data;
};

export const updateIOC = async (id: string, ioc: Partial<IOC>): Promise<IOC> => {
  const response = await client.put<IOC>(`/iocs/${id}`, ioc);
  return response.data;
};

export const deleteIOC = async (id: string): Promise<void> => {
  await client.delete(`/iocs/${id}`);
};

export const bulkCreateIOCs = async (iocs: Partial<IOC>[]): Promise<{ message: string; inserted: number }> => {
  const response = await client.post<{ message: string; inserted: number }>('/iocs/bulk', { iocs });
  return response.data;
};
