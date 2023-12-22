/**
 * 上传文件转base64
 * @param {*} file 上传文件
 * @returns string
 */
export const getBase64 = (file) => {
  return new Promise((resolve, reject) => {
    let fileResult = "";
    let reader = new FileReader();
      reader.readAsDataURL(file);
      //开始转
      reader.onload = function () {
        fileResult = reader.result;
      };
      //转 失败
      reader.onerror = function (error) {
        reject(error);
      };
      //转 结束  咱就 resolve 出去
      reader.onloadend = function () {
        resolve(fileResult);
      };
  });
};
