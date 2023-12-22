import { api } from "src/boot/axios";

// 定义 API 的基本 URL
const urlPrefix = "/auth";

// 登录的 AJAX 请求函数
export const ajaxLogin = (params) => {
  return api.post(`${urlPrefix}/login`, params);
};

// 注册的 AJAX 请求函数
export const ajaxRegister = (params) => {
  return api.post(`${urlPrefix}/register`, params);
};

// 获取当前用户菜单
export const ajaxGetCurrentAccountMenuList = (params) => {
  return api.get(`${urlPrefix}/menus`, { params });
};
