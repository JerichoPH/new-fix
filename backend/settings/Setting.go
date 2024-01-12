package settings

import (
	"new-fix/utils"
	"new-fix/wrongs"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

// Setting 设置
type Setting struct {
	App      *ini.File
	DB       *ini.File
	Url      *ini.File
	Time     time.Time
	Timezone *time.Location
	RootPath string
}

var settingIns *Setting
var once sync.Once

// NewSetting 初始化设置
func NewSetting() *Setting {
	if settingIns == nil {
		once.Do(func() { settingIns = &Setting{} })

		settingIns.RootPath = utils.GetRootPath()

		appConfigFile, appConfigErr := ini.Load(filepath.Join(settingIns.RootPath, "/settings/app.ini"))
		if appConfigErr != nil {
			wrongs.ThrowForbidden("启动错误：加载程序配置文件失败")
		}

		dbConfigFile, dbConfigErr := ini.Load(filepath.Join(settingIns.RootPath, "/settings/db.ini"))
		if dbConfigErr != nil {
			wrongs.ThrowForbidden("启动错误：加载数据库配置文件失败")
		}

		urlConfigFile, urlConfigErr := ini.Load(filepath.Join(settingIns.RootPath, "/settings/url.ini"))
		if urlConfigErr != nil {
			wrongs.ThrowForbidden("启动错误：加载url配置文件失败")
		}

		settingIns.App = appConfigFile
		settingIns.DB = dbConfigFile
		settingIns.Url = urlConfigFile
		settingIns.Timezone, _ = time.LoadLocation(settingIns.App.Section("app").Key("timezone").MustString("Asia/Shanghai"))
		settingIns.Time = (&time.Time{}).In(settingIns.Timezone)

	}

	return settingIns
}
