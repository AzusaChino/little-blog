import axios, {AxiosInstance} from "axios";
import {config} from '../config'

const request: AxiosInstance = axios.create({
  baseURL: config.url,
  timeout: 5000
});

export default request;
