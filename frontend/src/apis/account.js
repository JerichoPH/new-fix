import { api } from "src/boot/axios";
import collect from "collect.js";

const urlPrefix = "/account";

/**
 * 获取用户列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetAccounts = (params) => {
  return api.get(urlPrefix, { params });
};

/**
 * 获取用户详情
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxGetAccount = (uuid, params) => {
  return api.get(`${urlPrefix}/${uuid}`, { params });
};

/**
 * 新建用户
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreAccount = (params) => {
  return api.post(urlPrefix, params);
};

/**
 * 编辑用户信息
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateAccount = (uuid, params) => {
  return api.put(`${urlPrefix}/${uuid}`, params);
};

/**
 * 删除用户
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyAccount = (uuid) => {
  return api.delete(`${urlPrefix}/${uuid}`);
};

/**
 * 批量删除用户
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyAccounts = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的用户");
  return api.post(`${urlPrefix}/destroyMany`, { uuids });
};

/**
 * 修改用户密码
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateAccountPassword = (uuid, params) => {
  return api.put(`${urlPrefix}/${uuid}/password`, params);
};
