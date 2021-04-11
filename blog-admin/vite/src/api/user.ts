import {LoginParam, RegisterParam} from "/@/models/user";
import request from "/@/utils/request";

const baseUrl = 'api/v1/user'

/**
 * 登录接口
 * @param lp 登录凭证
 */
export function login(lp: LoginParam) {
  return request.post(`${baseUrl}/login`, lp)
}

export function register(rp: RegisterParam) {
  return request.post(`${baseUrl}/register`, rp)
}
