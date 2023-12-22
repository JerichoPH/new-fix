package tools

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"new-fix/wrongs"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type FileSystem struct {
	path string
}

// NewFileSystem 构造函数
func NewFileSystem(path string) *FileSystem {
	if path == "" {
		path = GetRootPath()
	}
	return &FileSystem{path: path}
}

func GetRootPath() string {
	rootPath, _ := filepath.Abs(".")
	return rootPath
}

// GetCurrentPath 最终方案-全兼容
func GetCurrentPath(paths ...string) string {
	dir := getGoBuildPath()
	if strings.Contains(dir, getTmpDir()) {
		return getGoRunPath()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getGoBuildPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getGoRunPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// GetPath 获取路径
func (receiver FileSystem) GetPath() string {
	return receiver.path
}

// SetPath 设置路径
func (receiver FileSystem) SetPath(path string) FileSystem {
	receiver.path = path
	return receiver
}

// Join 拼接路径
func (receiver FileSystem) Join(paths []string) FileSystem {
	if len(paths) > 0 {
		for _, pathDatum := range paths {
			receiver.path = path.Join(receiver.path, pathDatum)
		}
	}
	return receiver
}

// IsExist 判断路径是否存在
func (receiver FileSystem) IsExist() bool {
	_, err := os.Stat(receiver.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// GetDirStruct 获取目录结构
func (receiver FileSystem) GetDirStruct() ([]string, error) {
	var files []string
	err := filepath.Walk(receiver.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// CopyFile 复制文件
func (receiver FileSystem) CopyFile(dst string) error {
	srcFile, err := os.Open(receiver.path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

// MoveFile 移动文件
func (receiver FileSystem) MoveFile(dst string) error {
	err := os.Rename(receiver.path, dst)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile 删除文件
func (receiver FileSystem) DeleteFile() error {
	return os.Remove(receiver.path)
}

// CreateDir 创建目录
func (receiver FileSystem) CreateDir() error {
	return os.MkdirAll(receiver.path, 0755)
}

// DeleteDir 删除目录
func (receiver FileSystem) DeleteDir() error {
	err := os.RemoveAll(receiver.path)
	if err != nil {
		return err
	}
	return nil
}

// RenameFile 重命名文件
func (receiver FileSystem) RenameFile(newName string) error {
	newPath := filepath.Join(filepath.Dir(receiver.path), newName)
	return os.Rename(receiver.path, newPath)
}

// ReadStr 以文字读取文件内容
func (receiver FileSystem) ReadStr() string {
	f, err := ioutil.ReadFile(receiver.path)
	if err != nil {
		wrongs.ThrowForbidden("文件读取失败：%s", err.Error())
	}
	return string(f)
}

// ReadJson 读取json文件
func (receiver FileSystem) ReadJson() interface{} {
	f, err := ioutil.ReadFile(receiver.path)
	if err != nil {
		wrongs.ThrowForbidden("文件读取失败：%s", err.Error())
	}

	var t interface{}
	err = json.Unmarshal(f, &t)
	if err != nil {
		wrongs.ThrowForbidden("JSON解析失败: %s", err.Error())
	}

	return t
}

// WriteJson 写入json文件
func (receiver FileSystem) WriteJson(value any) FileSystem {
	if bytes, err := json.Marshal(value); err != nil {
		wrongs.ThrowForbidden("JSON解析失败: %s", err.Error())
	} else {
		receiver.WriteStr(string(bytes))
	}
	return receiver
}

// WriteByte 写入二进制文件
func (receiver FileSystem) WriteByte(value []byte) error {
	file, err := os.OpenFile(receiver.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	// 及时关闭file句柄
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			wrongs.ThrowForbidden("文件关闭失败: %s", err.Error())
		}
	}(file)

	if _, err := file.Write(value); err != nil {
		return err
	}

	return nil
}

// WriteStr 写入string文件
func (receiver FileSystem) WriteStr(value string) FileSystem {
	file, err := os.OpenFile(receiver.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		wrongs.ThrowForbidden("文件打开失败: %s", err.Error())
	}

	// 及时关闭file句柄
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			wrongs.ThrowForbidden("文件关闭失败: %s", err.Error())
		}
	}(file)

	if _, err := file.WriteString(value); err != nil {
		wrongs.ThrowForbidden("写入文件失败：%s", err.Error())
	}

	return receiver
}
