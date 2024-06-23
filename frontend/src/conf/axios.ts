import axios, { type AxiosRequestConfig, type AxiosInstance } from "axios"
import { isAuthorized } from "$store";

let customaxios: AxiosInstance

function configureAxios(): AxiosInstance {

  let config: AxiosRequestConfig = {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json'
    },
  }
  customaxios = axios.create(config)

  customaxios.interceptors.response.use(null, function(error) {
    if (error.response.status == 401) {
      isAuthorized.set(false);
    }
    return Promise.reject(error);
  });
  return customaxios
}

export function getAxios(): AxiosInstance {
  if (customaxios == undefined) {
    return configureAxios()
  }
  return customaxios
}

