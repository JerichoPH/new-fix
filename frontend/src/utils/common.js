/**
 * 表格排序
 * @param {*} event 事件对象
 * @param {{*}} props 作用域参数
 * @param {string} sortBy 当前排序字段
 */
export const fnColumnReverseSort = (event, props, sortBy) => {
  const name = event.target.getAttribute("name");
  const isCurrent = sortBy === name;
  sortBy = isCurrent ? "" : name;
  props.sort(name, !isCurrent);
};

export const isEmpty = (val) => {
  if (!val) return true;
  return false;
};

export const getVal = (obj = {}, keys = "") => {
  keys = keys.split(".");
  if (!obj || !keys.length) {
    return null;
  }

  let key = keys.shift();
  if (obj.hasOwnProperty(key)) {
    return keys.length ? getVal(obj[key], keys.join(".")) : obj[key];
  } else {
    return null;
  }
};

export const setVal = (obj, keys, val, defaultVal = null) => {
  keys = keys.split(".");
  let key = keys.shift();

  if (keys.length) {
    if (!obj.hasOwnProperty(key) || typeof obj[key] !== "object") {
      obj[key] = {};
    }
    setVal(obj[key], keys.join("."), val || defaultVal);
  } else {
    obj[key] = val;
  }
};
