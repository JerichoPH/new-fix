import { boot } from "quasar/wrappers";
import axios from "axios";
import settings from "src/settings";
import {errorNotify} from "src/utils/notify";

// Be careful when using SSR for cross-request state pollution
// due to creating a Singconston instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({
  baseURL: settings.apiUrl,
  headers: {
    "Content-Type": "application/json", // 设置请求头中的 Content-Type
    Accept: "application/json", // 设置请求头中的 Accept
  },
});

// 在请求发送之前进行拦截
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("auth.token"); // 从 localStorage 中获取 token
  if (token) config.headers.Authorization = `JWT ${token}`; // 如果有 token，则将其添加到请求头中

  return config;
});

// 在响应拦截之后进行处理
api.interceptors.response.use(
  (resp) => {
    return resp.data;
  },
  (e) => {
    if (e.code === "ERR_NETWORK") {
      return Promise.reject({
        msg: "网络错误",
      });
    }
    if (e.response.status === 401) {
      // 如果响应的状态码为 401（未登陆）
      localStorage.removeItem("auth.token"); // 从 localStorage 中删除 token
      errorNotify("未登录", 1000); // 显示通知，提示用户未登录
    } else {
      return Promise.reject({ msg: e.response.data.msg || e.response.data });
    }
  }
);

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});

export { api };
