import { api } from "src/boot/axios";
import collect from "collect.js";

const urlPrefix = "/warehouse";

export const ajaxGetWarehouseStorehouses = (params) => {
  return api.get(`${urlPrefix}/storehouse`, { params });
};

export const ajaxGetWarehouseStorehouse = (uuid, params) => {
  return api.get(`${urlPrefix}/storehouse/${uuid}`, { params });
};

export const ajaxStoreWarehouseStorehouse = (params) => {
  return api.post(`${urlPrefix}/storehouse`, params);
};

export const ajaxUpdateWarehouseStorehouse = (uuid, params) => {
  return api.put(`${urlPrefix}/storehouse/${uuid}`, params);
};

export const ajaxDestroyWarehouseStorehouse = (uuid) => {
  return api.delete(`${urlPrefix}/storehouse/${uuid}`);
};

export const ajaxDestroyWarehouseStorehouses = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的仓库-库房");
  return api.post(`${urlPrefix}/storehouse/destroyMany`, { uuids });
};

export const ajaxGetWarehouseAreas = (params) => {
  return api.get(`${urlPrefix}/area`, { params });
};

export const ajaxGetWarehouseArea = (uuid, params) => {
  return api.get(`${urlPrefix}/area/${uuid}`, { params });
};

export const ajaxStoreWarehouseArea = (params) => {
  return api.post(`${urlPrefix}/area`, params);
};

export const ajaxUpdateWarehouseArea = (uuid, params) => {
  return api.put(`${urlPrefix}/area/${uuid}`, params);
};

export const ajaxDestroyWarehouseArea = (uuid) => {
  return api.delete(`${urlPrefix}/area/${uuid}`);
};

export const ajaxDestroyWarehouseAreas = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的仓库-库区");
  return api.post(`${urlPrefix}/area/destroyMany`, { uuids });
};

export const ajaxGetWarehousePlatoons = (params) => {
  return api.get(`${urlPrefix}/platoon`, { params });
};

export const ajaxGetWarehousePlatoon = (uuid, params) => {
  return api.get(`${urlPrefix}/platoon/${uuid}`, { params });
};

export const ajaxStoreWarehousePlatoon = (params) => {
  return api.post(`${urlPrefix}/platoon`, params);
};

export const ajaxUpdateWarehousePlatoon = (uuid, params) => {
  return api.put(`${urlPrefix}/platoon/${uuid}`, params);
};

export const ajaxDestroyWarehousePlatoon = (uuid) => {
  return api.delete(`${urlPrefix}/platoon/${uuid}`);
};

export const ajaxDestroyWarehousePlatoons = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的仓库-排");
  return api.post(`${urlPrefix}/platoon/destroyMany`, { uuids });
};

export const ajaxGetWarehousePlatoonTypeCodesMap = () => {
  return api.get(`${urlPrefix}/platoon/typeCodesMap`);
};

export const ajaxGetWarehouseShelves = (params) => {
  return api.get(`${urlPrefix}/shelf`, { params });
};

export const ajaxGetWarehouseShelf = (uuid, params) => {
  return api.get(`${urlPrefix}/shelf/${uuid}`, { params });
};

export const ajaxStoreWarehouseShelf = (params) => {
  return api.post(`${urlPrefix}/shelf`, params);
};

export const ajaxUpdateWarehouseShelf = (uuid, params) => {
  return api.put(`${urlPrefix}/shelf/${uuid}`, params);
};

export const ajaxDestroyWarehouseShelf = (uuid) => {
  return api.delete(`${urlPrefix}/shelf/${uuid}`);
};

export const ajaxDestroyWarehouseShelves = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的仓库-货架");
  return api.post(`${urlPrefix}/shelf/destroyMany`, { uuids });
};

export const ajaxGetWarehouseTiers = (params) => {
  return api.get(`${urlPrefix}/tier`, { params });
};

export const ajaxGetWarehouseTier = (uuid, params) => {
  return api.get(`${urlPrefix}/tier/${uuid}`, { params });
};

export const ajaxStoreWarehouseTier = (params) => {
  return api.post(`${urlPrefix}/tier`, params);
};

export const ajaxStoreWarehouseTiers = (params) => {
  return api.post(`${urlPrefix}/tier/storeMany`, params);
};

export const ajaxUpdateWarehouseTier = (uuid, params) => {
  return api.put(`${urlPrefix}/tier/${uuid}`, params);
};

export const ajaxDestroyWarehouseTier = (uuid) => {
  return api.delete(`${urlPrefix}/tier/${uuid}`);
};

export const ajaxDestroyWarehouseTiers = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的仓库-层");
  return api.post(`${urlPrefix}/tier/destroyMany`, { uuids });
};

export const ajaxGetWarehouseCells = (params) => {
  return api.get(`${urlPrefix}/cell`, { params });
};

export const ajaxGetWarehouseCell = (uuid, params) => {
  return api.get(`${urlPrefix}/cell/${uuid}`, { params });
};

export const ajaxStoreWarehouseCell = (params) => {
  return api.post(`${urlPrefix}/cell`, params);
};

export const ajaxStoreWarehouseCells = (params) => {
  return api.post(`${urlPrefix}/cell/storeMany`, params);
};

export const ajaxUpdateWarehouseCell = (uuid, params) => {
  return api.put(`${urlPrefix}/cell/${uuid}`, params);
};

export const ajaxDestroyWarehouseCell = (uuid) => {
  return api.delete(`${urlPrefix}/cell/${uuid}`);
};

export const ajaxDestroyWarehouseCells = (uuids) => {
  if (collect(uuids).isEmpty()) new Promise.reject("请选择要删除的仓库-位");
  return api.post(`${urlPrefix}/cell/destroyMany`, { uuids });
};
