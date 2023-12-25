import { api } from "src/boot/axios";

// 定义 API 的基本 URL
const urlPrefix = "/auth";

/**
 * 登陆
 * @param {{*}} params
 * @returns
 */
export const ajaxLogin = (params) => {
  return api.post(`${urlPrefix}/login`, params);
};

/**
 * 注册
 * @param {{*}} params
 * @returns
 */
export const ajaxRegister = (params) => {
  return api.post(`${urlPrefix}/register`, params);
};

/**
 * 获取当前用户所有菜单
 * @param {{*}} params
 * @returns
 */
export const ajaxGetCurrentAccountMenus = (params) => {
  return api.get(`${urlPrefix}/menus`, { params });
};
