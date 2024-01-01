import { api } from "src/boot/axios";
import collect from "collect.js";

const urlPrefix = "/source";

export const ajaxGetSourceTypes = (params) => {
  return api.get(`${urlPrefix}/type`, { params });
};

export const ajaxGetSourceType = (uuid, params) => {
  return api.get(`${urlPrefix}/type/${uuid}`, { params });
};

export const ajaxStoreSourceType = (params) => {
  return api.post(`${urlPrefix}/type`, params);
};

export const ajaxUpdateSoruceType = (uuid, params) => {
  return api.put(`${urlPrefix}/type/${uuid}`, params);
};

export const ajaxDestroySourceType = (uuid) => {
  return api.delete(`${urlPrefix}/type/${uuid}`);
};

export const ajaxDestroySourceTypes = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的来源类型");
  return api.delete(`${urlPrefix}/type`, { uuids });
};

export const ajaxGetSourceProjects = (params) => {
  return api.get(`${urlPrefix}/project`, { params });
};

export const ajaxGetSourceProject = (uuid, params) => {
  return api.get(`${urlPrefix}/project/${uuid}`, { params });
};

export const ajaxStoreSourceProject = (params) => {
  return api.post(`${urlPrefix}/project`, params);
};

export const ajaxUpdateSourceProject = (uuid, params) => {
  return api.put(`${urlPrefix}/project/${uuid}`, params);
};

export const ajaxDestroySourceProject = (uuid) => {
  return api.delete(`${urlPrefix}/project/${uuid}`);
};

export const ajaxDestroySourceProjects = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的来源项目");
  return api.post(`${urlPrefix}/project/destroyMany`, { uuids });
};
