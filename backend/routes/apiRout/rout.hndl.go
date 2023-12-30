package apiRout

import (
	"github.com/gin-gonic/gin"
)

// RoutHandle 分组路由
type RoutHandle struct{}

// Register 组册路由
func (RoutHandle) Register(engine *gin.Engine) {
	NewTestRout().Load(engine)          // 测试
	NewAuthRout().Load(engine)          // 权鉴
	NewAccountRout().Load(engine)       // 用户
	NewRbacRout().Load(engine)          // 权限
	NewEquipmentKindRout().Load(engine) // 器材种类型
	NewOrganizationRout().Load(engine)  // 组织结构
	NewBreakdownRout().Load(engine)     // 故障
	NewFactoryRout().Load(engine)       // 厂家
}
