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
