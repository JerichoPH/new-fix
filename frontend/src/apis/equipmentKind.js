import { api } from "src/boot/axios";

const urlPrefix = "/equipmentKind";

/**
 * 器材种类列表
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindCategoryList = params => {
  return api.get(`${urlPrefix}/category`, { params });
};

/**
 * 器材种类详情
 * @param {string} uuid
 * @returns
 */
export const ajaxEquipmentKindCategoryDetail = uuid => {
  return api.get(`${urlPrefix}/category/${uuid}`);
};

/**
 * 新建器材种类
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindCategoryStore = params => {
  return api.post(`${urlPrefix}/category`, params);
}

/**
 * 编辑器材种类
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindCategoryUpdate = (uuid, params) => {
  return api.put(`${urlPrefix}/category/${uuid}`, params);
}

/**
 * 删除器材种类
 * @param {string} uuid
 * @returns
 */
export const ajaxEquipmentKindCategoryDestroy = uuid => {
  return api.delete(`${urlPrefix}/category/${uuid}`);
}

/**
 * 批量删除器材种类
 * @param {[string]} uuids
 * @returns
 */
export const ajaxEquipmentKindCategoryDestroyMany = uuids => {
  if (uuids.length === 0) return Promise.reject('器材种类编号列表不能为空');
  return api.post(`${urlPrefix}/category/destroyMany`, { uuids });
}

/**
 * 器材类型列表
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindTypeList = params => {
  return api.get(`${urlPrefix}/type`, { params });
}

/**
 * 器材类型详情
 * @param {string} uuid
 * @returns
 */
export const ajaxEquipmentKindTypeDetail = uuid => {
  return api.get(`${urlPrefix}/type/${uuid}`);
}

/**
 * 新建器材类型
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindTypeStore = params => {
  return api.post(`${urlPrefix}/type`, params);
}

/**
 * 编辑器材类型
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindTypeUpdate = (uuid, params) => {
  return api.put(`${urlPrefix}/type/${uuid}`, params);
}

/**
 * 删除器材类型
 * @param {string} uuid
 * @returns
 */
export const ajaxEquipmentKindTypeDestroy = uuid => {
  return api.delete(`${urlPrefix}/type/${uuid}`);
}

/**
 * 批量删除器材类型
 * @param {[string]} uuids
 * @returns
 */
export const ajaxEquipmentKindTypeDestroyMany = uuids => {
  if (uuids.length === 0) return Promise.reject('器材类型编号列表不能为空');
  return api.post(`${urlPrefix}/type/destroyMany`, { uuids });
}

/**
 * 器材型号列表
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindModelList = params => {
  return api.get(`${urlPrefix}`, { params });
}

/**
 * 器材型号详情
 * @param {string} uuid
 * @returns
 */
export const ajaxEquipmentKindModelDetail = uuid => {
  return api.get(`${urlPrefix}/${uuid}`);
}

/**
 * 新建器材型号
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindModelStore = params => {
  return api.post(`${urlPrefix}`, params);
}

/**
 * 编辑器材型号
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxEquipmentKindModelUpdate = (uuid, params) => {
  return api.put(`${urlPrefix}/${uuid}`, params);
}

/**
 * 删除器材型号
 * @param {string} uuid
 * @returns
 */
export const ajaxEquipmentKindModelDestroy = uuid => {
  return api.delete(`${urlPrefix}/${uuid}`);
}

/**
 * 批量删除器材型号
 * @param {[string]} uuids
 * @returns
 */
export const ajaxEquipmentKindModelDestroyMany = uuids => {
  if (uuids.length === 0) return Promise.reject('器材型号编号列表不能为空');
  return api.post(`${urlPrefix}/destroyMany`, { uuids });
}
