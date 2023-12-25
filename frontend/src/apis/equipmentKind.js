import { api } from "src/boot/axios";

const urlPrefix = "/equipmentKind";

/**
 * 器材种类列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetEquipmentKindCategories = params => {
  return api.get(`${urlPrefix}/category`, { params });
};

/**
 * 器材种类详情
 * @param {string} uuid
 * @returns
 */
export const ajaxGetEquipmentKindCategory = uuid => {
  return api.get(`${urlPrefix}/category/${uuid}`);
};

/**
 * 新建器材种类
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreEquipmentKindCategory = params => {
  return api.post(`${urlPrefix}/category`, params);
}

/**
 * 编辑器材种类
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateEquipmentKindCategory = (uuid, params) => {
  return api.put(`${urlPrefix}/category/${uuid}`, params);
}

/**
 * 删除器材种类
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyEquipmentKindCategory = uuid => {
  return api.delete(`${urlPrefix}/category/${uuid}`);
}

/**
 * 批量删除器材种类
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyEquipmentKindCategories = uuids => {
  if (uuids.length === 0) return Promise.reject('器材种类编号列表不能为空');
  return api.post(`${urlPrefix}/category/destroyMany`, { uuids });
}

/**
 * 器材类型列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetEquipmentKindTypes = params => {
  return api.get(`${urlPrefix}/type`, { params });
}

/**
 * 器材类型详情
 * @param {string} uuid
 * @params {{*}} params
 * @returns
 */
export const ajaxGetEquipmentKindType = (uuid, params) => {
  return api.get(`${urlPrefix}/type/${uuid}`, { params });
}

/**
 * 新建器材类型
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreEquipmentKindType = params => {
  return api.post(`${urlPrefix}/type`, params);
}

/**
 * 编辑器材类型
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateEquipmentKindType = (uuid, params) => {
  return api.put(`${urlPrefix}/type/${uuid}`, params);
}

/**
 * 删除器材类型
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyEquipmentKindType = uuid => {
  return api.delete(`${urlPrefix}/type/${uuid}`);
}

/**
 * 批量删除器材类型
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyEquipmentKindTypes = uuids => {
  if (uuids.length === 0) return Promise.reject('器材类型编号列表不能为空');
  return api.post(`${urlPrefix}/type/destroyMany`, { uuids });
}

/**
 * 器材型号列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetEquipmentKindModels = params => {
  return api.get(`${urlPrefix}/model`, { params });
}

/**
 * 器材型号详情
 * @param {string} uuid
 * @returns
 */
export const ajaxGetEquipmentKindModel = uuid => {
  return api.get(`${urlPrefix}/model/${uuid}`);
}

/**
 * 新建器材型号
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreEquipmentKindModel = params => {
  return api.post(`${urlPrefix}/model`, params);
}

/**
 * 编辑器材型号
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateEquipmentKindModel = (uuid, params) => {
  return api.put(`${urlPrefix}/model/${uuid}`, params);
}

/**
 * 删除器材型号
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyEquipmentKindModel = uuid => {
  return api.delete(`${urlPrefix}/model/${uuid}`);
}

/**
 * 批量删除器材型号
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyEquipmentKindModels = uuids => {
  if (uuids.length === 0) return Promise.reject('器材型号编号列表不能为空');
  return api.post(`${urlPrefix}/model/destroyMany`, { uuids });
}
