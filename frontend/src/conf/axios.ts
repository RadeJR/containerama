import axios, { type AxiosRequestConfig, type AxiosInstance } from "axios"

let customaxios: AxiosInstance

export function getAxios(): AxiosInstance {
  let config: AxiosRequestConfig = {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json'
    },
  }
  
  console.log(customaxios)
  if (customaxios == undefined) {
    customaxios = axios.create(config)
  }

  return customaxios
}

