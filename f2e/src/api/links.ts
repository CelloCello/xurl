import { Net } from '../utils/net';

export interface LinkPayload {
  id: string;
  url: string;
  code: string;
  created_at: string;
  updated_at: string;
}

export const getLinks = async () => {
  try {
    const resp = await Net.get<LinkPayload[]>('links');
    return resp.data.payload;
  } catch (error) {
    console.error(`error: ${error}`);
    throw error;
  }
};
