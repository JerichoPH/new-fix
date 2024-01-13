package controllers

import (
	"fmt"
	"new-fix/providers"
	"new-fix/settings"
	"new-fix/types"
	"new-fix/utils"
	"new-fix/wrongs"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

type (
	// TestCtrl 测试控制器
	TestCtrl struct{}

	// 发送消息表单
	sendMessageForm struct {
		ReceiverUuid   string `json:"receiver_uuid"`
		MessageTitle   string `json:"message_title"`
		MessageContent any    `json:"message_content"`
	}
)

// NewTestCtrl 构造函数
func NewTestCtrl() *TestCtrl {
	return &TestCtrl{}
}

// ShouldBind 表单绑定
func (receiver sendMessageForm) ShouldBind(ctx *gin.Context) sendMessageForm {
	err := ctx.ShouldBind(&receiver)
	if err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
		return sendMessageForm{}
	}

	return receiver
}

// AnySendToWebsocket 接收并转发消息(websocket)
func (receiver TestCtrl) AnySendToWebsocket(ctx *gin.Context) {
	form := sendMessageForm{}.ShouldBind(ctx)

	providers.WebsocketSendMessageByUuid(utils.NewCorrectWithBusiness(form.MessageTitle, "message", "").Datum(form.MessageContent).ToJsonStr(), form.ReceiverUuid)

	ctx.JSON(utils.NewCorrectWithGinContext("OK", ctx).Datum(map[string]any{"content": fmt.Sprintf("%s ==> %s %v", form.ReceiverUuid, form.MessageTitle, form.MessageContent)}).ToGinResponse())
}

// AnySendToTcpServer 接收并转发消息(tcp server)
func (receiver TestCtrl) AnySendToTcpServer(ctx *gin.Context) {
	form := sendMessageForm{}.ShouldBind(ctx)

	providers.TcpServerSendMessageByUuid(utils.NewCorrectWithBusiness(form.MessageTitle, "message", "").Datum(form.MessageContent).ToJsonStr(), form.ReceiverUuid)

	ctx.JSON(utils.NewCorrectWithGinContext("OK", ctx).Datum(map[string]any{"content": fmt.Sprintf("%s ==> %s %v", form.ReceiverUuid, form.MessageTitle, form.MessageContent)}).ToGinResponse())
}

// AnySendToTcpClient 接收病转发消息(tcp client)
func (receiver TestCtrl) AnySendToTcpClient(ctx *gin.Context) {
	form := sendMessageForm{}.ShouldBind(ctx)

	providers.TcpClientSendMessage(utils.NewCorrectWithBusiness(form.MessageTitle, "message", "").Datum(map[string]any{"content": fmt.Sprintf("%s %v", form.MessageTitle, form.MessageContent)}).ToJsonStr())

	ctx.JSON(utils.NewCorrectWithGinContext("OK", ctx).Datum(map[string]any{"content": fmt.Sprintf("%s %v", form.MessageTitle, form.MessageContent)}).ToGinResponse())
}

// AnySendToKafkaClient 发送消息到kafka
func (receiver TestCtrl) AnySendToKafkaClient(ctx *gin.Context) {
	var (
		appSetting *ini.File
		topic      string
		err        error
	)

	appSetting = settings.NewSetting().App
	topic = appSetting.Section("kafka-server").Key("topic").MustString("")

	form := sendMessageForm{}.ShouldBind(ctx)

	if err = providers.KafkaSendMessage(topic, "message", utils.NewCorrectWithBusiness(form.MessageTitle, "message", "").Datum(map[string]any{"content": fmt.Sprintf("%s %v", form.MessageTitle, form.MessageContent)}).ToJsonStr()); err != nil {
		wrongs.ThrowForbidden("发送失败：%s", err.Error())
	}
}

// 发送消息到rabbitmq
func (receiver TestCtrl) AnySendToRabbitMq(ctx *gin.Context) {
	now := time.Now().Format(string(types.TIME_FORMAT_DATETIME))
	providers.RabbitMqSendMessage("new-fix-data-conversion-layer", utils.NewCorrectWithBusiness(now, "ping", "").Datum(map[string]any{"content": "rabbitmq"}).ToJsonStr())
	ctx.JSON(utils.NewCorrectWithGinContext("OK", ctx).Datum(map[string]any{"content": "rabbitmq"}).ToGinResponse())
}
