package commands

import (
	"fmt"
	"time"

	"new-fix/database"
	"new-fix/utils"
	"new-fix/wrongs"

	uuid "github.com/satori/go.uuid"
)

// TempCmd 临时终端
type TempCmd struct{}

// NewTempCmd 构造函数
func NewTempCmd() *TempCmd {
	return &TempCmd{}
}

func (receiver TempCmd) uuid() []string {
	c := make(chan string)
	go func(c chan string) {
		uuidStr := uuid.NewV4().String()
		c <- uuidStr
	}(c)
	go utils.NewTimer(5).Ticker()
	return []string{<-c}
}

func (receiver TempCmd) ls() []string {
	_, res := (&utils.Cmd{}).Process("ls", "-la")
	return []string{res}
}

func (receiver TempCmd) redis() []string {
	if _, err := database.NewRedis(0).SetValue("test", "AAA", 15*time.Minute); err != nil {
		wrongs.ThrowForbidden(err.Error())
	}

	for i := 0; i < 100000; i++ {
		if val, err := database.NewRedis(0).GetValue("test"); err != nil {
			wrongs.ThrowForbidden(err.Error())
		} else {
			fmt.Println(i, val)
		}
	}

	return []string{""}
}

func (TempCmd) test() []string {
	std := utils.NewStdoutHelper("OK")
	std.EchoLineDebug(utils.GetCurrentPath())
	std.EchoLineDebug(utils.GetRootPath())

	return []string{}
}

// Handle 执行命令
func (receiver TempCmd) Handle(params []string) []string {
	switch params[0] {
	case "uuid":
		return receiver.uuid()
	case "ls":
		return receiver.ls()
	case "redis":
		return receiver.redis()
	case "test":
		return receiver.test()
	default:
		return []string{"没有找到命令"}
	}
}
