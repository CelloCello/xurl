import axios, { Method } from 'axios';

export interface ResponseData<T = any> {
  msg: string;
  payload: T;
  error: string;
}

export interface Response<T = any> {
  data: ResponseData<T>;
  status: number;
  //   request?: any;
}

export type ReqMethod = Method;

const request = async <Payload>(
  method: ReqMethod,
  url: string,
  params?: any,
  data?: any,
): Promise<Response<Payload>> => {
  const baseURL = 'http://localhost:8080/api';
  try {
    const resp = await axios.request<ResponseData<Payload>>({
      url: url,
      method: method,
      baseURL: baseURL,
      // headers?: AxiosRequestHeaders,
      // headers: { "Access-Control-Allow-Origin": "*" },
      params: params,
      data: data,
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
    return request<T>('get', url, params);
  },

  post: <T = any>(url: string, params?: any, data?: any) => {
    return request<T>('post', url, params, data);
  },

  put: <T = any>(url: string, params?: any, data?: any) => {
    return request<T>('put', url, params, data);
  },

  delete: <T = any>(url: string, params?: any, data?: any) => {
    return request<T>('delete', url, params, data);
  },
};
