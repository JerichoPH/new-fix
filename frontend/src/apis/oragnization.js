import { api } from "src/boot/axios";

const apiPrefix = "/organization";

/**
 * 获取路局列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetOrganizationRailways = params => {
  return api.get(`${apiPrefix}/railway`, { params });
};

/**
 * 获取路局详情
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxGetOrganizationRailway = (uuid, params) => {
  return api.get(`${apiPrefix}/railway/${uuid}`, { params });
};

/**
 * 新建路局
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreOrganizationRailway = params => {
  return api.post(`${apiPrefix}/railway`, params);
};

/**
 * 编辑路局
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateOrganizationRailway = (uuid, params) => {
  return api.put(`${apiPrefix}/railway/${uuid}`, params);
};

/**
 * 删除路局
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyOrganizationRailway = uuid => {
  return api.delete(`${apiPrefix}/railway/${uuid}`);
}

/**
 * 批量删除路局
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyOrganizationRailways = uuids => {
  if (uuids.length === 0) return Promise.reject('路局编号列表不能为空');
  return api.post(`${apiPrefix}/railway/destroyMany`, { uuids });
}
