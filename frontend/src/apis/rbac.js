import { api } from "src/boot/axios";
import collect from "collect.js";

const urlPrefix = "/rbac";

/**
 * 角色列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetRbacRoles = (params) => {
  return api.get(`${urlPrefix}/role`, { params });
};

/**
 * 角色详情
 * @param {string} uuid
 * @returns
 */
export const ajaxGetRbacRole = (uuid) => {
  return api.get(`${urlPrefix}/role/${uuid}`);
};

/**
 * 新建角色
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreRbacRole = (params) => {
  return api.post(`${urlPrefix}/role`, params);
};

/**
 * 编辑角色
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateRbacRole = (uuid, params) => {
  return api.put(`${urlPrefix}/role/${uuid}`, params);
};

/**
 * 删除角色
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyRbacRole = (uuid) => {
  return api.delete(`${urlPrefix}/role/${uuid}`);
};

/**
 * 批量删除角色
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyRbacRoles = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的角色");
  return api.post(`${urlPrefix}/role/destroyMany`, { uuids });
};

/**
 * 权限列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetRbacPermissions = (params) => {
  return api.get(`${urlPrefix}/permission`, { params });
};

/**
 * 权限详情
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxGetRbacPermission = (uuid, params) => {
  return api.get(`${urlPrefix}/permission/${uuid}`, { params });
};

/**
 * 新建权限
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreRbacPermission = (params) => {
  return api.post(`${urlPrefix}/permission`, params);
};

/**
 * 更新权限
 * @param {string} uuid
 * @param {{*}} params
 */
export const ajaxUpdateRbacPermission = (uuid, params) => {
  return api.put(`${urlPrefix}/permission/${uuid}`, params);
};

/**
 * 删除权限
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyRbacPermission = (uuid) => {
  return api.delete(`${urlPrefix}/permission/${uuid}`);
};

/**
 * 批量删除权限
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyRbacPermissions = (uuids) => {
  if(collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的权限");
  return api.post(`${urlPrefix}/permission/destroyMany`, { uuids });
};

/**
 * 菜单列表
 * @param {{*}} params
 * @returns
 */
export const ajaxGetRbacMenus = (params) => {
  return api.get(`${urlPrefix}/menu`, { params });
};

/**
 * 菜单详情
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxGetRbacMenu = (uuid, params) => {
  return api.get(`${urlPrefix}/menu/${uuid}`, { params });
};

/**
 * 新建菜单
 * @param {{*}} params
 * @returns
 */
export const ajaxStoreRbacMenu = (params) => {
  return api.post(`${urlPrefix}/menu`, params);
};

/**
 * 编辑菜单
 * @param {string} uuid
 * @param {{*}} params
 * @returns
 */
export const ajaxUpdateRbacMenu = (uuid, params) => {
  return api.put(`${urlPrefix}/menu/${uuid}`, params);
};

/**
 * 删除菜单
 * @param {string} uuid
 * @returns
 */
export const ajaxDestroyRbacMenu = (uuid) => {
  return api.delete(`${urlPrefix}/menu/${uuid}`);
};

/**
 * 批量删除菜单
 * @param {[string]} uuids
 * @returns
 */
export const ajaxDestroyRbacMenus = (uuids) => {
  if (collect(uuids).isEmpty()) return Promise.reject("请选择需要删除的菜单");
  return api.post(`${urlPrefix}/menu/destroyMany`, { uuids });
};
