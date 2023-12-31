import { api } from "/src/boot/axios";
import collect from "collect.js";

const urlPrefix = "/breakdown";

export const ajaxGetBreakdownTypes = (params) => {
  return api.get(`${urlPrefix}/type`, { params });
};

export const ajaxGetBreakdownType = (uuid, params) => {
  return api.get(`${urlPrefix}/type/${uuid}`, { params });
};

export const ajaxStoreBreakdownType = (params) => {
  return api.post(`${urlPrefix}/type`, params);
};

export const ajaxUpdateBreakdownType = (uuid, params) => {
  return api.put(`${urlPrefix}/type/${uuid}`, params);
};

export const ajaxDestroyBreakdownType = (uuid) => {
  return api.delete(`${urlPrefix}/type/${uuid}`);
};

export const ajaxDestroyBreakdownTypes = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的故障类型");
  return api.post(`${urlPrefix}/type/destroyMany`, { uuids });
};

export const ajaxGetBreakdownLogs = (params) => {
  return api.get(`${urlPrefix}/log`, { params });
};

export const ajaxGetBreakdownLog = (uuid) => {
  return api.get(`${urlPrefix}/log/${uuid}`);
};