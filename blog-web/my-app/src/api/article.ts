import { AxiosResponse } from "axios";
import request from "../utils/request";

const baseURL = "api/v1/article";

export const fetchList = (query: any): Promise<AxiosResponse> => {
  return request.get(baseURL, { params: query });
};
