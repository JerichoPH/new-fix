package middlewares

import (
	"new-fix/models"
	"new-fix/tools"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			account               models.AccountMdl
			rbacPermissionUuids   []string
			currentUri            string
			currentRbacPermission *models.RbacPermissionMdl
			ret                   *gorm.DB
			yes                   = false
		)
		// 获取当前用户
		account = tools.GetAuth(ctx).(models.AccountMdl)
		if !account.BeAdmin {
			rbacPermissionUuids = account.GetPermissionUuids()

			// 获取当前路由
			currentUri = ctx.Request.URL.Path

			// 查询当前路由是否存在权限
			ret = models.NewRbacPermissionMdl().GetDb("").Where("uri", currentUri).First(&currentRbacPermission)
			wrongs.ThrowWhenEmpty(ret, "当前路由对应权限")

			// 检查当前路由是否合法
			for _, uuid := range rbacPermissionUuids {
				if uuid == currentRbacPermission.Uuid {
					yes = true
					break
				}
			}

			if !yes {
				wrongs.ThrowUnAuth("权限不足")
			}
		}

		ctx.Next()
	}
}
