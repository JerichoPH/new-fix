import { api } from "src/boot/axios";
import collect  from "collect.js";

const apiPrefix = "/organization";

export const ajaxGetOrganizationRailways = (params) => {
  return api.get(`${apiPrefix}/railway`, { params });
};

export const ajaxGetOrganizationRailway = (uuid, params) => {
  return api.get(`${apiPrefix}/railway/${uuid}`, { params });
};

export const ajaxStoreOrganizationRailway = (params) => {
  return api.post(`${apiPrefix}/railway`, params);
};

export const ajaxUpdateOrganizationRailway = (uuid, params) => {
  return api.put(`${apiPrefix}/railway/${uuid}`, params);
};

export const ajaxDestroyOrganizationRailway = (uuid) => {
  return api.delete(`${apiPrefix}/railway/${uuid}`);
};

export const ajaxDestroyOrganizationRailways = (uuids) => {
  if(collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的路局");
  return api.post(`${apiPrefix}/railway/destroyMany`, { uuids });
};

export const ajaxGetOrganizationParagraphs = (params) => {
  return api.get(`${apiPrefix}/paragraph`, { params });
};

export const ajaxGetOrganizationParagraph = (uuid, params) => {
  return api.get(`${apiPrefix}/paragraph/${uuid}`, { params });
};

export const ajaxStoreOrganizationParagraph = (params) => {
  return api.post(`${apiPrefix}/paragraph`, params);
};

export const ajaxUpdateOrganizationParagraph = (uuid, params) => {
  return api.put(`${apiPrefix}/paragraph/${uuid}`, params);
};

export const ajaxDestroyOrganizationParagraph = (uuid) => {
  return api.delete(`${apiPrefix}/paragraph/${uuid}`);
};

export const ajaxDestroyOrganizationParagraphs = (uuids) => {
  if(collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的站段");
  return api.post(`${apiPrefix}/paragraph/destroyMany`, { uuids });
};

export const ajaxGetOrganizationWorkshops = (params) => {
  return api.get(`${apiPrefix}/workshop`, { params });
};

export const ajaxGetOrganizationWorkshop = (uuid, params) => {
  return api.get(`${apiPrefix}/workshop/${uuid}`, { params });
};

export const ajaxStoreOrganizationWorkshop = (params) => {
  return api.post(`${apiPrefix}/workshop`, params);
};

export const ajaxUpdateOrganizationWorkshop = (uuid, params) => {
  return api.put(`${apiPrefix}/workshop/${uuid}`, params);
};

export const ajaxDestroyOrganizationWorkshop = (uuid) => {
  return api.delete(`${apiPrefix}/workshop/${uuid}`);
};

export const ajaxGetOrganizationWorkshopTypeCodesMap = () => {
  return api.get(`${apiPrefix}/workshop/typeCodesMap`);
};

export const ajaxDestroyOrganizationWorkshops = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的车间");
  return api.post(`${apiPrefix}/workshop/destroyMany`, { uuids });
};

export const ajaxGetOrganizationStations = (params) => {
  return api.get(`${apiPrefix}/station`, { params });
};

export const ajaxGetOrganizationStation = (uuid, params) => {
  return api.get(`${apiPrefix}/station/${uuid}`, { params });
};

export const ajaxStoreOrganizationStation = (params) => {
  return api.post(`${apiPrefix}/station`, params);
};

export const ajaxUpdateOrganizationStation = (uuid, params) => {
  return api.put(`${apiPrefix}/station/${uuid}`, params);
};

export const ajaxDestroyOrganizationStation = (uuid) => {
  return api.delete(`${apiPrefix}/station/${uuid}`);
};

export const ajaxDestroyOrganizationStations = (uuids) => {
  if(collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的车站");
  return api.post(`${apiPrefix}/station/destroyMany`, { uuids });
};

export const ajaxGetOrganizationCrossroads = (params) => {
  return api.get(`${apiPrefix}/crossroad`, { params });
};

export const ajaxGetOrganizationCrossroad = (uuid, params) => {
  return api.get(`${apiPrefix}/crossroad/${uuid}`, { params });
};

export const ajaxStoreOrganizationCrossroad = (params) => {
  return api.post(`${apiPrefix}/crossroad`, params);
};

export const ajaxUpdateOrganizationCrossroad = (uuid, params) => {
  return api.put(`${apiPrefix}/crossroad/${uuid}`, params);
};

export const ajaxDestroyOrganizationCrossroad = (uuid) => {
  return api.delete(`${apiPrefix}/crossroad/${uuid}`);
};

export const ajaxDestroyOrganizationCrossroads = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的道口");
  return api.post(`${apiPrefix}/crossroad/destroyMany`, { uuids });
};

export const ajaxGetOrganizationCenters = (params) => {
  return api.get(`${apiPrefix}/center`, { params });
};

export const ajaxGetOrganizationCenter = (uuid, params) => {
  return api.get(`${apiPrefix}/center/${uuid}`, { params });
};

export const ajaxStoreOrganizationCenter = (params) => {
  return api.post(`${apiPrefix}/center`, params);
};

export const ajaxUpdateOrganizationCenter = (uuid, params) => {
  return api.put(`${apiPrefix}/center/${uuid}`, params);
};

export const ajaxDestroyOrganizationCenter = (uuid) => {
  return api.delete(`${apiPrefix}/center/${uuid}`);
};

export const ajaxDestroyOrganizationCenters = (uuids) => {
  if(collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的中心");
  return api.post(`${apiPrefix}/center/destroyMany`, { uuids });
};

export const ajaxGetOrganizationWorkAreas = (params) => {
  return api.get(`${apiPrefix}/workArea`, { params });
};

export const ajaxGetOrganizationWorkArea = (uuid, params) => {
  return api.get(`${apiPrefix}/workArea/${uuid}`, { params });
};

export const ajaxStoreOrganizationWorkArea = (params) => {
  return api.post(`${apiPrefix}/workArea`, params);
};

export const ajaxUpdateOrganizationWorkArea = (uuid, params) => {
  return api.put(`${apiPrefix}/workArea/${uuid}`, params);
};

export const ajaxDestroyOrganizationWorkArea = (uuid) => {
  return api.delete(`${apiPrefix}/workArea/${uuid}`);
};

export const ajaxDestroyOrganizationWorkAreas = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的工区");
  return api.post(`${apiPrefix}/workArea/destroyMany`, { uuids });
};

export const ajaxGetOrganizationWorkAreaTypeCodesMap = () => {
  return api.get(`${apiPrefix}/workArea/typeCodesMap`);
};

export const ajaxGetOrganizationLines = (params) => {
  return api.get(`${apiPrefix}/line`, { params });
};

export const ajaxGetOrganizationLine = (uuid, params) => {
  return api.get(`${apiPrefix}/line/${uuid}`, { params });
};

export const ajaxStoreOrganizationLine = (params) => {
  return api.post(`${apiPrefix}/line`, params);
};

export const ajaxUpdateOrganizationLine = (uuid, params) => {
  return api.put(`${apiPrefix}/line/${uuid}`, params);
};

export const ajaxDestroyOrganizationLine = (uuid) => {
  return api.delete(`${apiPrefix}/line/${uuid}`);
};

export const ajaxDestroyOrganizationLines = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的线别");
  return api.post(`${apiPrefix}/line/destroyMany`, { uuids });
};
