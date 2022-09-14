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

export const createLink = async (url: string) => {
  try {
    const resp = await Net.post<LinkPayload>('links', { url });
    return resp.data.payload;
  } catch (error) {
    console.error(`error: ${error}`);
    throw error;
  }
};

export const deleteLink = async (id: string) => {
  try {
    const resp = await Net.delete(`links/${id}`);
    return resp.status === 200;
  } catch (error) {
    console.error(`error: ${error}`);
    throw error;
  }
};
