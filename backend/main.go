package main

import (
	"flag"
	"fmt"
	"log"
	"new-fix/models"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"new-fix/commands"
	"new-fix/database"
	"new-fix/middlewares"
	"new-fix/providers"
	"new-fix/routes/apiRout"
	"new-fix/routes/webRout"
	"new-fix/settings"
	"new-fix/tools"
	"new-fix/types"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
)

type (
	// WebServiceCommand web-service 服务
	// WebServiceCommand struct{}

	// WebServiceCommandSetting web-service 服务配置
	WebServiceCommandSetting struct {
		WebAddr           string
		TcpServerEnable   bool
		TcpServerAddr     string
		TcpClientEnable   bool
		TcpClientAddr     string
		RabbitMqEnable    bool
		RabbitMqAddr      string
		RabbitMqUsername  string
		RabbitMqPassword  string
		RabbitMqVhost     string
		RabbitMqQueueName string
		KafkaServerEnable bool
		KafkaServerAddr   string
		KafkaServerTopic  string
		KafkaClientEnable bool
		KafakClientAddr   string
		KafkaClientTopic  string
		RocketMqEnable    bool
		RocketMqAddr      string
	}
)

// 自动迁移列表
var autoMigrateList = []any{
	models.AccountMdl{},               // 用户
	models.RbacRoleMdl{},              // 权限管理-角色
	models.RbacPermissionMdl{},        // 权限管理-权限
	models.RbacMenuMdl{},              // 权限管理-菜单
	models.FileMdl{},                  // 文件
	models.EquipmentKindCategoryMdl{}, // 器材种类
	models.EquipmentKindTypeMdl{},     // 器材类型
	models.EquipmentKindModelMdl{},    // 器材型号
	models.OrganizationRailwayMdl{},   // 组织结构-路局
	models.OrganizationParagraphMdl{}, // 组织结构-站段
	models.OrganizationWorkshopMdl{},  // 组织结构-车间
	models.OrganizationStationMdl{},   // 组织结构-站场
	models.OrganizationCrossroadMdl{}, // 组织结构-道口
	models.OrganizationCenterMdl{},    // 组织结构-中心
	models.OrganizationWorkAreaMdl{},  // 组织结构-工区
	models.OrganizationLineMdl{},      // 组织结构-线别
	models.BreakdownTypeMdl{},         // 故障-故障类型
	models.BreakdownLogMdl{},          // 故障-故障日志
	models.SourceTypeMdl{},            // 来源-来源类型
	models.SourceProjectMdl{},         // 来源-来源项目
	models.FactoryMdl{},               // 厂家
	models.EquipmentMdl{},             // 器材
	models.WarehouseStorehouseMdl{},   // 库房-仓库
	models.WarehouseAreaMdl{},         // 库房-库区
	models.WarehousePlatoonMdl{},      // 库房-排
	models.WarehouseShelfMdl{},        // 库房-架
	models.WarehouseTierMdl{},         // 库房-层
	models.WarehouseCellMdl{},         // 库房-格
	models.WarehouseScanItemMdl{},     // 库房-扫描器材项
	models.WarehouseOrderMdl{},        // 库房-出入库单
	models.InstallIndoorRoomTypeMdl{}, // 室内上道位置-机房类型
	models.InstallIndoorRoomMdl{},     // 室内上道位置-机房
	models.InstallIndoorPlatoonMdl{},  // 室内上道位置-排
	models.InstallIndoorShelfMdl{},    // 室内上道位置-架
	models.InstallIndoorTierMdl{},     // 室内上道位置-层
	models.InstallIndoorCellMdl{},     // 室内上道位置-格
}

// 启动守护进程
func launchDaemon(title string) {
	if syscall.Getppid() == 1 {
		if err := os.Chdir("./"); err != nil {
			panic(err)
		}
		// syscall.Umask(0) // TODO TEST
		return
	}

	currentDir := os.Getenv("PWD")
	logDir := fmt.Sprintf("%s/logs", currentDir)
	if !tools.NewFileSystem(logDir).IsExist() {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			fmt.Println("创建日志目录错误：" + err.Error())
			return
		}
	}
	filenameOnToday := fmt.Sprintf("%s/logs/%s.log", currentDir, time.Now().Format(string(types.TIME_FORMAT_DATE)))
	fp, err := os.OpenFile(filenameOnToday, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = fp.Close()
	}()
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdout = fp
	cmd.Stderr = fp
	cmd.Stdin = nil
	// cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true} // TODO TEST
	if err = cmd.Start(); err != nil {
		panic(err)
	}

	startLogContent := fmt.Sprintf("%s 程序启动成功 [进程号%d] 启动于：%s\r\n", title, cmd.Process.Pid, time.Now().Format(string(types.TIME_FORMAT_DATETIME)))
	fmt.Println(startLogContent)

	_, _ = fp.WriteString("\r\n\r\n\r\n>>>>>>>>>>>>>>>>>>>>\r\n" + startLogContent)
	os.Exit(0)
}

// 执行command命令
func launchCommand(commandName string, commandParams, tmp []string, daemon bool) {
	var (
		commandResults []string
		commandSetting WebServiceCommandSetting
		appSetting     = settings.NewSetting().App
		appName        = appSetting.Section("app").Key("name").String()
	)

	commandSetting = WebServiceCommandSetting{
		WebAddr:           appSetting.Section("web-service").Key("addr").MustString(":8091"),
		TcpServerEnable:   appSetting.Section("tcp-server-service").Key("enable").MustBool(false),
		TcpServerAddr:     appSetting.Section("tcp-server-service").Key("addr").MustString("0.0.0.0:8092"),
		TcpClientEnable:   appSetting.Section("tcp-client-service").Key("enable").MustBool(false),
		TcpClientAddr:     appSetting.Section("tcp-client-service").Key("addr").MustString("127.0.0.1:8080"),
		RabbitMqEnable:    appSetting.Section("rabbit-mq-service").Key("enable").MustBool(false),
		RabbitMqAddr:      appSetting.Section("rabbit-mq-service").Key("addr").MustString("127.0.0.1:5672"),
		RabbitMqUsername:  appSetting.Section("rabbit-mq-service").Key("username").MustString(""),
		RabbitMqPassword:  appSetting.Section("rabbit-mq-service").Key("password").MustString(""),
		RabbitMqVhost:     appSetting.Section("rabbit-mq-service").Key("vhost").MustString(""),
		RabbitMqQueueName: appSetting.Section("rabbit-mq-service").Key("queue-name").MustString(""),
		KafkaServerEnable: appSetting.Section("kafka-server").Key("enable").MustBool(false),
		KafkaServerAddr:   appSetting.Section("kafka-server").Key("addr").MustString(":9092"),
		KafkaServerTopic:  appSetting.Section("kafka-server").Key("topic").MustString("revolution"),
		KafkaClientEnable: appSetting.Section("kafka-client").Key("enable").MustBool(false),
		KafakClientAddr:   appSetting.Section("kafka-client").Key("addr").MustString(":9092"),
		KafkaClientTopic:  appSetting.Section("kafka-client").Key("topic").MustString("revolution"),
		RocketMqEnable:    appSetting.Section("rocket-mq").Key("enable").MustBool(false),
		RocketMqAddr:      appSetting.Section("rocket-mq").Key("addr").MustString(""),
	}

	if daemon {
		launchDaemon(fmt.Sprintf("%s %s", appName, commandName))
	}

	switch commandName {
	default:
		// println(tools.StdoutWrong(fmt.Sprintf("『执行失败』%s", strings.Join(tmp, " ")), "").GetContentAndNext("没有找到命令"))
		// os.Exit(-1)
		// 启动web-service服务

		// 模型
		database.NewGormLauncher().BootAutoMigrate(autoMigrateList)

		// 创建TCP服务端
		if commandSetting.TcpServerEnable {
			go providers.TcpServerHandler(commandSetting.TcpServerAddr)
		}

		// 创建TCP客户端
		if commandSetting.TcpClientEnable {
			go providers.TcpClientHandler(commandSetting.TcpClientAddr)
		}

		// 创建RabbitMQ监听
		if commandSetting.RabbitMqEnable {
			go providers.RabbitMqHandler(
				commandSetting.RabbitMqUsername,
				commandSetting.RabbitMqPassword,
				commandSetting.RabbitMqAddr,
				commandSetting.RabbitMqVhost,
				commandSetting.RabbitMqQueueName,
			)
		}

		// 启动Kafka服务端
		if commandSetting.KafkaServerEnable {
			go providers.KafkaServerHandler(commandSetting.KafkaServerAddr, commandSetting.KafkaServerTopic)
		}

		// 启动Kafka客户端
		if commandSetting.KafkaClientEnable {
			go providers.KafkaClientHandler(commandSetting.KafakClientAddr, commandSetting.KafkaClientTopic)
		}

		// 启动rocket-mq
		if commandSetting.RocketMqEnable {
			go providers.RocketMqHandler(commandSetting.RocketMqAddr)
		}

		engine := gin.Default()                                     // 启动服务引擎
		engine.Use(middlewares.Cors())                              // 允许跨域
		engine.Use(middlewares.TimeoutMiddleware(time.Second * 60)) // 超时处理
		engine.Use(wrongs.RecoverHandler)                           // 异常处理
		webRout.RoutHandle{}.Register(engine)                       // 加载web路由
		apiRout.RoutHandle{}.Register(engine)                       // 加载api路由

		// 绑定web-socket路由
		engine.GET("/ws", func(ctx *gin.Context) {
			providers.WebsocketHandler(ctx)
		})

		// 启动web-service服务
		if err := engine.Run(commandSetting.WebAddr); err != nil {
			log.Fatalf("[web-service-error] [启动web服务错误] %v", err)
		}

		log.Printf("[app] [程序停止] [%s]", time.Now().Format(string(types.TIME_FORMAT_DATETIME)))
	case "temp":
		commandResults = commands.NewTempCmd().Handle(commandParams)
	case "upgrade":
		commandResults = commands.NewUpgradeCmd().Do(commandParams)
	}

	fmt.Println(tools.StdoutDebug(fmt.Sprintf("『执行完成 %s』 「%s」 ", tools.NewTime().SetTimeNowAdd8Hour().ToDateTimeString(), strings.Join(tmp, " ")), "").GetContentAndNext(strings.Join(commandResults, " ")))
	os.Exit(0)
}

// main 程序入口
func main() {
	// params explain:
	// t=command 执行命令行
	// t='web-service' 启动web服务 可选项-port=8080 端口号、-daemon=false 守护进程
	var (
		title, commandName string
		commandParams, tmp []string
		daemon             bool
	)
	flag.StringVar(&title, "script", "", "命令终端参数")
	flag.BoolVar(&daemon, "daemon", false, "是否启动守护进程")
	flag.Parse()

	commandName = ""
	commandParams = make([]string, 0)

	if title != "" {
		tmp = strings.Split(title, " ")
		commandName = tmp[0]
		commandParams = tmp[1:]
	}

	launchCommand(commandName, commandParams, tmp, daemon)
}
