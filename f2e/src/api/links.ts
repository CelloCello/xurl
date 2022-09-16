import { Net } from '../utils/net';
import { LinkPayload } from '../types';

export const fetchLinks = async () => {
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
