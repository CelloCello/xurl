import axios, { Method } from 'axios';
import { ROOT_URL } from '../constants';

export interface ResponseData<T = any> {
  msg: string;
  payload: T;
  error: string;
}

export interface Request {
  method: ReqMethod;
  url: string;
  params?: any;
  data?: any;
}

export interface Response<T = any> {
  data: ResponseData<T>;
  status: number;
  //   request?: any;
}

export type ReqMethod = Method;

const request = async <Payload>(req: Request): Promise<Response<Payload>> => {
  const baseURL = `${ROOT_URL}/api`;
  try {
    const resp = await axios.request<ResponseData<Payload>>({
      url: req.url,
      method: req.method,
      baseURL: baseURL,
      // headers?: AxiosRequestHeaders,
      // headers: { 'Access-Control-Allow-Origin': '*' },
      params: req.params,
      data: req.data,
      timeout: 60 * 1000,
    });
    return {
      data: resp.data,
      status: resp.status,
    };
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const Net = {
  request,
  get: <T = any>(url: string, params?: any) => {
    return request<T>({ method: 'get', url, params });
  },

  post: <T = any>(url: string, data?: any, params?: any) => {
    return request<T>({ method: 'post', url, params, data });
  },

  put: <T = any>(url: string, data?: any, params?: any) => {
    return request<T>({ method: 'put', url, params, data });
  },

  delete: <T = any>(url: string, data?: any, params?: any) => {
    return request<T>({ method: 'delete', url, params, data });
  },
};
