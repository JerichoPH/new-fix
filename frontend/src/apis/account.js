import { api } from "src/boot/axios";

const urlPrefix = "/account";

export const ajaxAccountList = (params) => {
  return api.get(urlPrefix, { params });
};

export const ajaxAccountDetail = (uuid, params) => {
  return api.get(`${urlPrefix}/${uuid}`, { params });
};

export const ajaxAccountStore = (params) => {
  return api.post(urlPrefix, params);
};

export const ajaxAccountUpdate = (uuid, params) => {
  return api.put(`${urlPrefix}/${uuid}`, params);
};

export const ajaxAccountDestroy = (uuid) => {
  return api.delete(`${urlPrefix}/${uuid}`);
};

export const ajaxAccountDestroyMany = uuids => {
  if (uuids.length === 0) return Promise.reject('账号编号列表不能为空');
  return api.post(`${urlPrefix}/destroyMany`, { uuids });
};

export const ajaxAccountUpdatePassword = (uuid, params) => {
  return api.put(`${urlPrefix}/${uuid}/password`, params);
};
