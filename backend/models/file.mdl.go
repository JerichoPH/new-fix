package models

import (
	"encoding/base64"
	"fmt"
	"new-fix/tools"
	"new-fix/wrongs"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type (
	// FileMdl 文件模型
	FileMdl struct {
		MySqlMdl
		SavePath          string        `gorm:"type:text;comment:文件保存路径;" json:"save_path"`
		Filename          string        `gorm:"unique;type:varchar(128);comment:文件保存名;" json:"filename"`
		OriginalFilename  string        `gorm:"type:varchar(255);not null;default:'';comment:文件原始名;" json:"original_filename"`
		OriginalExtension string        `gorm:"type:varchar(255);not null;default:'';comment:文件原始扩展名;" json:"original_extension"`
		Accounts          []*AccountMdl `gorm:"foreignKey:avatar_uuid;references:uuid" json:"accounts"`
		Url               string        `gorm:"-" json:"url"`
	}
)

// TableName 文件表名称
func (FileMdl) TableName() string {
	return "files"
}

// NewFileMdl 新建文件模型
func NewFileMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(FileMdl{})
}

// GetListByQuery 根据Query获取文件列表
func (receiver FileMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewFileMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "f.name like ?",
		}).
		SetWheresDateBetween("f.created_at", "f.updated_at", "f.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("files as f")
}

// AfterFind 查询后填充url
func (receiver *FileMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.Url = fmt.Sprintf("%s/%s", receiver.SavePath, receiver.Filename)
	return
}

// CreateByBase64 新建文件(base64)
func (receiver *FileMdl) CreateByBase64(savePath string, originalFilename, originalExtension, base64Str string) (*FileMdl, error) {
	// 保存文件
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return receiver.CreateByByte(savePath, originalFilename, originalExtension, data)
}

// CreateByByte 新建文件(byte)
func (receiver *FileMdl) CreateByByte(savePath string, originalFilename, originalExtension string, values []byte) (*FileMdl, error) {
	avatarUuid := uuid.NewV4().String()
	filename := fmt.Sprintf("%s.%s", avatarUuid, originalExtension)
	file := &FileMdl{
		MySqlMdl:          MySqlMdl{Uuid: avatarUuid},
		SavePath:          savePath,
		Filename:          filename,
		OriginalFilename:  originalFilename,
		OriginalExtension: originalExtension,
	}

	// 保存用户头像
	if ret := NewFileMdl().GetDb("").Create(&file); ret.Error != nil {
		return nil, ret.Error
	}
	// 创建文件夹
	if fs := tools.NewFileSystem(tools.GetRootPath()).Join([]string{savePath}); !fs.IsExist() {
		fs.CreateDir()
	}
	if err := tools.NewFileSystem(tools.GetRootPath()).
		Join([]string{savePath, filename}).
		WriteByte(values); err != nil {
		return nil, err
	}

	return file, nil
}

// DeleteFile 删除文件
func (receiver *FileMdl) DeleteFile(uuid string) error {
	ret := NewFileMdl().
		GetDb("").
		Where("uuid = ?", uuid).
		First(&receiver)

	if !wrongs.ThrowWhenEmpty(ret, "") {
		p := path.Join(tools.GetRootPath(), receiver.SavePath, receiver.Filename)
		os.Remove(p)

		NewFileMdl().
			GetDb("").
			Where("uuid = ?", receiver.Uuid).
			Delete(nil)
	}

	return nil
}
