import axios, {AxiosInstance} from "axios";
import {config} from '../config'

const request: AxiosInstance = axios.create({
  baseURL: config.url,
  timeout: 5000,
  withCredentials: false // 以便access-control-allow-origin可以设置成*
});

export default request;
