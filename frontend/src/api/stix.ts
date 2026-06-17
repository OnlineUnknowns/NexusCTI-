import client from './client';

export const exportSTIXBundle = async (): Promise<any> => {
  const response = await client.post<any>('/stix/export');
  return response.data;
};

export const importSTIXBundle = async (bundle: any): Promise<{ message: string; stats: any }> => {
  const response = await client.post<{ message: string; stats: any }>('/stix/import', bundle);
  return response.data;
};
