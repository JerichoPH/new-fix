import { api } from "src/boot/axios";

const urlPrefix = "/rbac";

/**
 * 角色列表
 * @param {{ * }} params 参数
 * @returns
 */
export const ajaxRbacRoleList = (params) => {
  return api.get(`${urlPrefix}/role`, { params });
};

/**
 * 角色详情
 * @param {string} uuid 唯一编号
 * @returns
 */
export const ajaxRbacRoleDetail = (uuid) => {
  return api.get(`${urlPrefix}/role/${uuid}`);
};

/**
 * 新建角色
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacRoleStore = (params) => {
  return api.post(`${urlPrefix}/role`, params);
};

/**
 * 编辑角色
 * @param {string} uuid 唯一编号
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacRoleUpdate = (uuid, params) => {
  return api.put(`${urlPrefix}/role/${uuid}`, params);
};

/**
 * 删除角色
 * @param {string} uuid 唯一编号
 * @returns
 */
export const ajaxRbacRoleDestroy = (uuid) => {
  return api.delete(`${urlPrefix}/role/${uuid}`);
};

/**
 * 批量删除角色
 * @param {[string]} uuids 唯一编号列表
 * @returns
 */
export const ajaxRbacRoleDestroyMany = uuids => {
  if (uuids.length === 0) return Promise.reject('角色编号列表不能为空');
  return api.post(`${urlPrefix}/role/destroyMany`, { uuids });
};

/**
 * 权限列表
 * @param {{*}} params 权限列表
 * @returns
 */
export const ajaxRbacPermissionList = (params) => {
  return api.get(`${urlPrefix}/permission`, { params });
};

/**
 * 权限详情
 * @param {string} uuid 唯一编号
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacPermissionDetail = (uuid, params) => {
  return api.get(`${urlPrefix}/permission/${uuid}`, { params });
};

/**
 * 新建权限
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacPermissionStore = (params) => {
  return api.post(`${urlPrefix}/permission`, params);
};

/**
 * 更新权限
 * @param {string} uuid 唯一编号
 * @param {{*}} params 参数
 */
export const ajaxRbacPermissionUpdate = (uuid, params) => {
  return api.put(`${urlPrefix}/permission/${uuid}`, params);
};

/**
 * 删除权限
 * @param {string} uuid 唯一编号
 * @returns
 */
export const ajaxRbacPermissionDestroy = (uuid) => {
  return api.delete(`${urlPrefix}/permission/${uuid}`);
};

/**
 * 批量删除权限
 * @param {[string]} uuids 权限编号列表
 * @returns
 */
export const ajaxRbacPermissionDestroyMany = uuids => {
  if (uuids.length === 0) return Promise.reject('权限编号列表不能为空');
  return api.post(`${urlPrefix}/permission/destroyMany`, { uuids });
}

/**
 * 菜单列表
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacMenuList = (params) => {
  return api.get(`${urlPrefix}/menu`, { params });
};

/**
 * 菜单详情
 * @param {string} uuid 唯一编号
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacMenuDetail = (uuid, params) => {
  return api.get(`${urlPrefix}/menu/${uuid}`, { params });
};

/**
 * 新建菜单
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacMenuStore = (params) => {
  return api.post(`${urlPrefix}/menu`, params);
};

/**
 * 编辑菜单
 * @param {string} uuid 唯一编号
 * @param {{*}} params 参数
 * @returns
 */
export const ajaxRbacMenuUpdate = (uuid, params) => {
  return api.put(`${urlPrefix}/menu/${uuid}`, params);
};

/**
 * 删除菜单
 * @param {string} uuid 唯一编号
 * @returns
 */
export const ajaxRbacMenuDestroy = (uuid) => {
  return api.delete(`${urlPrefix}/menu/${uuid}`);
};

/**
 * 批量删除菜单
 * @param {[string]} uuids 唯一编号列表
 * @returns
 */
export const ajaxRbacMenuDestroyMany = uuids => {
  if (uuids.length === 0) return Promise.reject('菜单编号列表不能为空');
  return api.post(`${urlPrefix}/menu/destroyMany`, { uuids });
}
