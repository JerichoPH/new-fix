import { Notify } from "quasar";

export const getDestroyActions = (handler) => {
  return [
    {
      label: "取消",
      color: "white",
    },
    {
      label: "删除",
      color: "negative",
      handler,
    },
  ];
};

export const getErrorActions = () => {
  return [
    {
      label: "关闭",
      color: "white",
    },
  ];
};

export const loadingNotify = (
  message = "处理中……",
  callback = null,
  params = {}
) => {
  const n = Notify.create({
    type: "info",
    spinner: true,
    position: "top",
    progress: false,
    message,
    timeout: 0,
  });
  if (callback) {
    if (timeout) {
      setTimeout(() => {
        callback(params);
        n();
      }, timeout + 1000);
    } else {
      callback(params);
      n();
    }
  }
  return n;
};

export const infoNotify = (
  message,
  timeout = 1500,
  callback = null,
  params = {}
) => {
  const n = Notify.create({
    type: "info",
    spinner: false,
    position: "top",
    progress: true,
    message,
    timeout,
  });
  if (callback) {
    if (timeout) {
      setTimeout(() => {
        callback(params);
        n();
      }, timeout + 1000);
    } else {
      callback(params);
      n();
    }
  }
  return n;
};

export const successNotify = (
  message,
  timeout = 500,
  callback = null,
  params = {}
) => {
  const n = Notify.create({
    type: "positive",
    spinner: false,
    position: "top",
    progress: true,
    message,
    timeout,
  });
  if (callback) {
    if (timeout) {
      setTimeout(() => {
        callback(params);
        n();
      }, timeout + 1000);
    } else {
      callback(params);
      n();
    }
  }
  return n;
};

export const errorNotify = (message, handler = null) => {
  const n = Notify.create({
    position: "top",
    color: "negative",
    timeout: 0,
    message,
    actions: [
      {
        label: "关闭",
        color: "white",
        handler,
      },
    ],
  });

  return n;
};

/**
 * 交互通知
 * @param {string} message 消息内容
 * @param {[{*}]} actions 行为
 */
export const actionNotify = (
  actions = [],
  message = "删除不可恢复，是否继续？"
) => {
  const n = Notify.create({
    position: "center",
    color: "orange",
    message,
    actions,
  });

  return n;
};
