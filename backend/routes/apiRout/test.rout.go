package apiRout

import (
	"new-fix/controllers"

	"github.com/gin-gonic/gin"
)

// TestRout 路由
type TestRout struct{}

// NewTestRout 构造函数
func NewTestRout() TestRout {
	return TestRout{}
}

// Load 加载路由
func (TestRout) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/test",
		// middlewares.CheckAuthorization(),
		// middlewares.CheckPermission(),
	)
	{
		r.Any("sendToWebsocket", controllers.NewTestCtrl().AnySendToWebsocket)
		r.Any("sendToTcpServer", controllers.NewTestCtrl().AnySendToTcpServer)
		r.Any("sendToTcpClient", controllers.NewTestCtrl().AnySendToTcpClient)
		r.Any("sendToKafkaClient", controllers.NewTestCtrl().AnySendToKafkaClient)
		r.Any("sendToRabbitMq", controllers.NewTestCtrl().AnySendToRabbitMq)
	}
}
