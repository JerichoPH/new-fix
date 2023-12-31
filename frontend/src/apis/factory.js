import { api } from "/src/boot/axios";
import collect from "collect.js";

const urlPrefix = "/factory";

export const ajaxGetFactories = (params) => {
  return api.get(`${urlPrefix}`, { params });
};

export const ajaxGetFactory = (uuid, params) => {
  return api.get(`${urlPrefix}/${uuid}`, { params });
};

export const ajaxStoreFactory = (params) => {
  return api.post(`${urlPrefix}`, params);
};

export const ajaxUpdateFactory = (uuid, params) => {
  return api.put(`${urlPrefix}/${uuid}`, params);
};

export const ajaxDestroyFactory = (uuid) => {
  return api.delete(`${urlPrefix}/${uuid}`);
};

export const ajaxDestroyFactories = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的厂家");
  return api.post(`${urlPrefix}/destroyMany`, {uuids});
};
