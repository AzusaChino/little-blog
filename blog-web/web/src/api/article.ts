import request from "../utils/request";

const baseURL = "api/v1/article";

export const fetchList = (query: any) => {
  return request.get(baseURL, { params: query });
};

export const fetchArticleDetail = (id : string) => {
  return request.get(`${baseURL}/${id}`)
}
