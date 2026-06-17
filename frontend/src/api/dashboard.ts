import client from './client';
import { DashboardStats } from '../types';

export const getDashboardStats = async (): Promise<DashboardStats> => {
  const response = await client.get<DashboardStats>('/dashboard/stats');
  return response.data;
};
