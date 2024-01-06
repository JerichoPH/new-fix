import { api } from "src/boot/axios";
import collect from "collect.js";
import { notifies } from "src/utils/notify";

const urlPrefix = "/install";

export const ajaxGetInstallIndoorRoomTypes = (params) => {
  return api.get(`${urlPrefix}/indoorRoomType`, { params });
};

export const ajaxGetInstallIndoorRoomType = (uuid, params) => {
  return api.get(`${urlPrefix}/indoorRoomType/${uuid}`, { params });
};

export const ajaxStoreInstallIndoorRoomType = (params) => {
  return api.post(`${urlPrefix}/indoorRoomType`, params);
};

export const ajaxUpdateInstallIndoorRoomType = (uuid, params) => {
  return api.put(`${urlPrefix}/indoorRoomType/${uuid}`, params);
};

export const ajaxDestroyInstallIndoorRoomType = (uuid) => {
  return api.delete(`${urlPrefix}/indoorRoomType/${uuid}`);
};

export const ajaxDestroyInstallIndoorRoomTypes = (uuids) => {
  if (uuids.isEmpty()) return Promise.reject({ msg: "请选择需要删除的内容" });
  return api.post(`${urlPrefix}/indoorRoomType/destroyMany`, { uuids });
};

export const ajaxGetInstallIndoorRooms = (params) => {
  return api.get(`${urlPrefix}/indoorRoom`, { params });
};

export const ajaxGetInstallIndoorRoom = (uuid, params) => {
  return api.get(`${urlPrefix}/indoorRoom/${uuid}`, { params });
};

export const ajaxStoreInstallIndoorRoom = (params) => {
  return api.post(`${urlPrefix}/indoorRoom`, params);
};

export const ajaxUpdateInstallIndoorRoom = (uuid, params) => {
  return api.put(`${urlPrefix}/indoorRoom/${uuid}`, params);
};

export const ajaxDestroyInstallIndoorRoom = (uuid) => {
  return api.delete(`${urlPrefix}/indoorRoom/${uuid}`);
};

export const ajaxDestroyInstallIndoorRooms = (uuids) => {
  if (uuids.isEmpty()) return Promise.reject({ msg: "请选择需要删除的内容" });
  return api.post(`${urlPrefix}/indoorRoom/destroyMany`, { uuids });
};

export const ajaxGetInstallIndoorPlatoons = (params) => {
  return api.get(`${urlPrefix}/indoorPlatoon`, { params });
};

export const ajaxGetInstallIndoorPlatoon = (uuid, params) => {
  return api.get(`${urlPrefix}/indoorPlatoon/${uuid}`, { params });
};

export const ajaxStoreInstallIndoorPlatoon = (params) => {
  return api.post(`${urlPrefix}/indoorPlatoon`, params);
};

export const ajaxUpdateInstallIndoorPlatoon = (uuid, params) => {
  return api.put(`${urlPrefix}/indoorPlatoon/${uuid}`, params);
};

export const ajaxDestroyInstallIndoorPlatoon = (uuid) => {
  return api.delete(`${urlPrefix}/indoorPlatoon/${uuid}`);
};

export const ajaxDestroyInstallIndoorPlatoons = (uuids) => {
  if (uuids.isEmpty()) return Promise.reject({ msg: "请选择需要删除的内容" });
  return api.post(`${urlPrefix}/indoorPlatoon/destroyMany`, { uuids });
};

export const ajaxGetInstallIndoorShelves = (params) => {
  return api.get(`${urlPrefix}/indoorShelf`, { params });
};

export const ajaxGetInstallIndoorShelf = (uuid, params) => {
  return api.get(`${urlPrefix}/indoorShelf/${uuid}`, { params });
};

export const ajaxStoreInstallIndoorShelf = (params) => {
  return api.post(`${urlPrefix}/indoorShelf`, params);
};

export const ajaxUpdateInstallIndoorShelf = (uuid, params) => {
  return api.put(`${urlPrefix}/indoorShelf/${uuid}`, params);
};

export const ajaxDestroyInstallIndoorShelf = (uuid) => {
  return api.delete(`${urlPrefix}/indoorShelf/${uuid}`);
};

export const ajaxDestroyInstallIndoorShelves = (uuids) => {
  if (uuids.isEmpty()) return Promise.reject({ msg: "请选择需要删除的内容" });
  return api.post(`${urlPrefix}/indoorShelf/destroyMany`, { uuids });
};

export const ajaxGetInstallIndoorTiers = (params) => {
  return api.get(`${urlPrefix}/indoorTier`, { params });
};

export const ajaxGetInstallIndoorTier = (uuid, params) => {
  return api.get(`${urlPrefix}/indoorTier/${uuid}`, { params });
};

export const ajaxStoreInstallIndoorTier = (params) => {
  return api.post(`${urlPrefix}/indoorTier`, params);
};

export const ajaxUpdateInstallIndoorTier = (uuid, params) => {
  return api.put(`${urlPrefix}/indoorTier/${uuid}`, params);
};

export const ajaxDestroyInstallIndoorTier = (uuid) => {
  return api.delete(`${urlPrefix}/indoorTier/${uuid}`);
};

export const ajaxDestroyInstallIndoorTiers = (uuids) => {
  if (uuids.isEmpty()) return Promise.reject({ msg: "请选择需要删除的内容" });
  return api.post(`${urlPrefix}/indoorTier/destroyMany`, { uuids });
};

export const ajaxGetInstallIndoorCells = (params) => {
  return api.get(`${urlPrefix}/indoorCell`, { params });
};

export const ajaxGetInstallIndoorCell = (uuid, params) => {
  return api.get(`${urlPrefix}/indoorCell/${uuid}`, { params });
};

export const ajaxStoreInstallIndoorCell = (params) => {
  return api.post(`${urlPrefix}/indoorCell`, params);
};

export const ajaxUpdateInstallIndoorCell = (uuid, params) => {
  return api.put(`${urlPrefix}/indoorCell/${uuid}`, params);
};

export const ajaxDestroyInstallIndoorCell = (uuid) => {
  return api.delete(`${urlPrefix}/indoorCell/${uuid}`);
};

export const ajaxDestroyInstallIndoorCells = (uuids) => {
  if (uuids.isEmpty()) return Promise.reject({ msg: "请选择需要删除的内容" });
  return api.post(`${urlPrefix}/indoorCell/destroyMany`, { uuids });
};
